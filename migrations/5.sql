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

CREATE UNIQU