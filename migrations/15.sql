-- +migrate Up
CREATE TABLE IF NOT EXISTS related_example (
  "example_id" SERIAL NOT NULL,
  "related_example_id" SERIAL NOT NULL,
  CONSTRAINT related_example_example_id_fkey FOREIGN KEY ("example_id") REFERENCES example("id") ON UPDATE NO ACTION ON DELE