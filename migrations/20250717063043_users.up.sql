CREATE EXTENSION IF NOT EXISTS "pgcrypto";
-- Создаёт расширение pgcrypto, которое нужно для генерации UUID через gen_random_uuid()
-- Добавлено вручную для поддержки автоматической генерации UUID в PostgreSQL
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

