package postgres

import (
	"Backend/internal/model"
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecord = errors.New("подходящей записи не найдено")

func (s *PostgresStorage) AddStatus(ip string, alive bool, checked, lastSuccess time.Time) error {
	stmt := `
	INSERT INTO container_status (ip, alive, checked, lastSuccess)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (ip) 
	DO UPDATE SET alive = EXCLUDED.alive, checked = EXCLUDED.checked, lastSuccess = EXCLUDED.lastSuccess`
	_, err := s.DB.Exec(stmt, ip, alive, checked, lastSuccess)
	return err
}

func (s *PostgresStorage) GetStatus(ip string) (*model.ContainerStatus, error) {
	stmt := `SELECT ip, alive, checked, lastSuccess FROM container_status WHERE ip = $1`
	row := s.DB.QueryRow(stmt, ip)

	status := &model.ContainerStatus{}
	err := row.Scan(&status.IP, &status.Alive, &status.Checked, &status.LastSuccess)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return status, nil
}

func (s *PostgresStorage) GetAllStatuses() ([]*model.ContainerStatus, error) {
	stmt := `SELECT ip, alive, checked, lastSuccess FROM container_status ORDER BY checked DESC`
	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []*model.ContainerStatus
	for rows.Next() {
		status := &model.ContainerStatus{}
		err = rows.Scan(&status.IP, &status.Alive, &status.Checked, &status.LastSuccess)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return statuses, nil
}
