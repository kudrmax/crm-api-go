#!/bin/bash

# Настройки
DB_CONTAINER="db" # Имя контейнера Docker
DB_USER="${POSTGRES_USER}" # Имя пользователя базы данных
DB_PASSWORD="${POSTGRES_PASSWORD}" # Пароль
DB_NAME="${POSTGRES_DB}" # Имя базы данных
MIGRATIONS_DIR="./postgresql/master" # Папка с миграциями

# Установить пароль
export PGPASSWORD="$DB_PASSWORD"

# Выполнить каждую миграцию
for file in $(ls $MIGRATIONS_DIR/*.sql | sort); do
    echo "Running migration: $file"
    docker exec -i $DB_CONTAINER psql -U "$DB_USER" -d "$DB_NAME" -f - < "$file"
    if [ $? -ne 0 ]; then
        echo "Error while running migration: $file"
        exit 1
    fi
done

# Очистить переменную
unset PGPASSWORD

echo "All migrations applied successfully."
