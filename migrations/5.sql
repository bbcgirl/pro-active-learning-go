-- +migrate Up
CREATE TABLE IF NOT EXISTS hatena_bookmark (
  "id" SERIAL NOT NULL PRIMARY KEY,
  "example_id" SERIAL NOT NULL,
  "title" TEXT NOT NULL,
  "screenshot" TEXT NOT NULL,
  "entry_url" TEXT NOT NULL,
  "count" INT NOT NULL,
  "url" TEXT NOT NULL,
  "eid" TEXT NOT NULL,
  CONSTRAINT hatena_bookmark_example_id_fkey FOREIGN KEY ("example_id") REFERENCES example("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "example_id_idx_hatena_bookmark" ON hatena_bookmark ("example_id");
CREATE UNIQUE INDEX IF NOT EXISTS "url_idx_hatena_bookmark" ON hatena_bookmark ("url");

CREATE TABLE IF NOT E