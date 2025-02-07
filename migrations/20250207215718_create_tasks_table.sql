-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tasks (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     title TEXT NOT NULL,
                                     description TEXT,
                                     deadline TIMESTAMP WITH TIME ZONE,
                                     completed BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE IF EXISTS tasks;
