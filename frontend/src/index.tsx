import React from 'react';
import ReactDOM from 'react-dom/client'; // Используйте правильный импорт
import App from './App.tsx';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);