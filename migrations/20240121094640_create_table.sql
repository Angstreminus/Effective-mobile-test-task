-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY NOT NULL UNIQUE,
  name varchar(255),
  surname varchar(255),
  patronymic varchar(255),
  age int,
  gender varchar(255),
  nationality varchar(255),
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
