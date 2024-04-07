-- +migrate Up
CREATE TABLE IF NOT EXISTS recommendation (
  "list_type" INT NOT NULL,
  "example_id" SERIAL NOT NULL,
   CONSTRAINT recommendation_example_id_fkey FOREIGN KEY ("example_id") REFERENCES example("id") ON UPDATE N