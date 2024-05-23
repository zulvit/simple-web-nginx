# Установка и настройка проекта Simple Web Nginx

Этот документ содержит инструкции по установке и настройке проекта Simple Web Nginx.

## Шаги установки

Выполните следующие команды по порядку в вашем терминале:

```bash
# Обновление пакетов
sudo apt update

# Установка Go и Nginx
sudo apt install -y golang-go nginx

# Клонирование репозитория и переход в директорию проекта
echo "Клонирование репозитория..."
mkdir -p /home/$(whoami)/app
cd /home/$(whoami)/app
git clone https://github.com/zulvit/simple-web-nginx.git
cd simple-web-nginx

# Сборка проекта Go
echo "Сборка приложения Go..."
go build -o simple-web

# Копирование собранного приложения в директорию деплоя
echo "Размещение приложения..."
sudo mkdir -p /var/www/app
sudo cp simple-web /var/www/app/
sudo chmod +x /var/www/app/simple-web

# Настройка systemd сервиса для приложения
echo "Настройка systemd сервиса..."
SERVICE_FILE=/etc/systemd/system/simple-web.service
echo "[Unit]
Description=Simple Web App

[Service]
User=$(whoami)
ExecStart=/var/www/app/simple-web

[Install]
WantedBy=multi-user.target" | sudo tee $SERVICE_FILE

# Перезагрузка systemd и запуск приложения
echo "Перезапуск и запуск приложения..."
sudo systemctl daemon-reload
sudo systemctl start simple-web.service
sudo systemctl enable simple-web.service

# Настройка Nginx
echo "Настройка Nginx..."
NGINX_CONFIG="/etc/nginx/sites-available/default"
echo "server {
    listen 80;
    server_name localhost;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host \$host;
        proxy_cache_bypass \$http_upgrade;
    }
}" | sudo tee $NGINX_CONFIG

# Перезапуск Nginx
echo "Перезапуск Nginx..."
sudo systemctl restart nginx

# Вывод окончания скрипта
echo "Скрипт завершен. Приложение должно быть доступно через Nginx на порту 80"
