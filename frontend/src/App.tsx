import React, { useEffect, useState } from 'react';
import { Table } from 'antd';
import axios from 'axios';

interface ContainerStatus {
  ip: string;
  alive: boolean;
  checked: string;
  lastSuccess: string;
}

const formatDateTime = (dateString: string) => {
  const date = new Date(dateString);
  return new Intl.DateTimeFormat('ru-RU', {
    year: 'numeric',
    month: 'numeric',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
    second: 'numeric',
  }).format(date);
};

const columns = [
  {
    title: 'IP',
    dataIndex: 'ip',
    key: 'ip',
  },
  {
    title: 'Статус',
    dataIndex: 'alive',
    key: 'alive',
    render: (alive: boolean) => (alive ? '🟢 Alive' : '🔴 Dead'),
  },
  {
    title: 'Последняя проверка',
    dataIndex: 'checked',
    key: 'checked',
    render: (checked: string) => formatDateTime(checked),
  },
  {
    title: 'Успешная проверка',
    dataIndex: 'lastSuccess',
    key: 'lastSuccess',
    render: (lastSuccess: string) => formatDateTime(lastSuccess),
  },
];

const App: () => JSX.Element = () => {
  const [statuses, setStatuses] = useState<ContainerStatus[]>([]);

  useEffect(() => {
    const fetchStatuses = async () => {
      const response = await axios.get('http://localhost:8200/status');
      setStatuses(response.data);
    };

    fetchStatuses();
    const interval = setInterval(fetchStatuses, 5000); // Обновление каждые 5 секунд
    return () => clearInterval(interval);
  }, []);

  return (
      <div style={{ padding: '20px' }}>
        <h1>Состояние контейнеров</h1>
        <Table dataSource={statuses} columns={columns} rowKey="ip" />
      </div>
  );
};

export default App;