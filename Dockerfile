# Используем официальный образ Golang
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum, устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все исходные файлы
COPY . .

# Переходим в директорию с main.go и собираем приложение
WORKDIR /app/cmd/app
RUN go build -o duno_main .

# Используем минимальный образ для запуска
FROM ubuntu:22.04

# Устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Создаем директорию для хранения загруженных файлов
RUN mkdir -p /app/uploads/videos

# Устанавливаем рабочую директорию
WORKDIR .

# Копируем бинарник из предыдущего шага
COPY --from=builder /app/cmd/app/duno_main /app

# Копируем конфигурационный файл
COPY config/config.yml /config/

# Указываем порт, на котором приложение будет работать
EXPOSE 8001

# Устанавливаем разрешения для выполнения файла
RUN chmod +x ./app/duno_main

# Копируем папку для загрузки видео
COPY ./uploads /app/uploads

# Указываем команду для запуска
CMD ["./app/duno_main"]
