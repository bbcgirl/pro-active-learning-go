-- +migrate Up
CREATE TABLE IF NOT EXISTS feature (
  "example_id" SERIAL NOT NULL,
  "feature" TEXT NOT NULL,
  CONSTRAINT featu