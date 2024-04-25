-- +migrate Up
CREATE TABLE IF NOT EXISTS hatena_bookmark (
  "id" SERIAL NOT NULL PRIMARY KEY,
  "example_id" SERIAL NOT NULL,
  "title" TEXT NOT NULL,
  "screenshot" TEXT NOT NULL,
  "entry_url" TEXT NOT NULL,
  "count" INT NOT NULL,
  