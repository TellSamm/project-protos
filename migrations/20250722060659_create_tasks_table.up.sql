CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE tasks (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
   title VARCHAR(255) NOT NULL,
   is_done BOOLEAN NOT NULL DEFAULT false,
   user_id UUID NOT NULL,
   deleted_at TIMESTAMP DEFAULT NULL
);