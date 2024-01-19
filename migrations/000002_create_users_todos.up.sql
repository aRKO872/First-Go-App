CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" text NOT NULL,
  "email" text UNIQUE NOT NULL,
  "password" text NOT NULL
);

CREATE TABLE IF NOT EXISTS "todos" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" text NOT NULL,
  "description" text NOT NULL,
  "done" BOOLEAN DEFAULT FALSE,
  "user_id" UUID NOT NULL REFERENCES users("id") ON DELETE CASCADE
);