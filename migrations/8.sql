-- +migrate Up
ALTER TABLE "tweet" DROP COLUMN "retweeted";

-- +migrate Down
ALTER TABLE "tweet" ADD