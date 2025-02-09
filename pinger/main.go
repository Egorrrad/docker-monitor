package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainerStatus struct {
	IP          string    `json:"ip"`
	Alive       bool      `json:"alive"`
	Checked     time.Time `json:"checked"`
	LastSuccess time.Time `json:"lastSuccess"`
}

var lastSuccessMap = sync.Map{} // Храним время последнего успешного пинга

func main() {
	for {
		ipMap, err := getContainerIPs()
		if err != nil {
			log.Println("Ошибка получения IP-адресов:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		var wg sync.WaitGroup

		for _, ip := range ipMap {
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				checkAndSend(ip)
			}(ip)
		}

		wg.Wait() // Ждём завершения всех горутин
		time.Sleep(10 * time.Second)
	}
}

func getContainerIPs() (map[string]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}

	ipMap := make(map[string]string)
	for _, container := range containers {
		for _, network := range container.NetworkSettings.Networks {
			if network.IPAddress != "" {
				ipMap[container.ID[:12]] = network.IPAddress
				break
			}
		}
	}
	return ipMap, nil
}

func pingContainer(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	return err == nil
}

func checkAndSend(ip string) {
	alive := pingContainer(ip)

	lastSuccess, _ := lastSuccessMap.LoadOrStore(ip, time.Time{})
	if alive {
		lastSuccessMap.Store(ip, time.Now())
	} else {
		lastSuccessMap.Store(ip, lastSuccess) // Оставляем прошлое значение
	}
	lastSuccess, _ = lastSuccessMap.Load(ip)
	status := ContainerStatus{
		IP:          ip,
		Alive:       alive,
		Checked:     time.Now(),
		LastSuccess: lastSuccess.(time.Time),
	}

	for retries := 0; retries < 3; retries++ { // 3 попытки отправки
		err := sendStatus(status)
		if err == nil {
			return
		}
		log.Println("Ошибка отправки статуса, попытка", retries+1, ":", err)
		time.Sleep(2 * time.Second)
	}
}

func sendStatus(status ContainerStatus) error {
	jsonData, err := json.Marshal(status)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://backend:8200/status", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	clientt := &http.Client{Timeout: 5 * time.Second}
	resp, err := clientt.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API вернул %d", resp.StatusCode)
	}

	return nil
}
