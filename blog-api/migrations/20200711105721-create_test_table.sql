-- +migrate Up
CREATE TABLE IF NOT EXISTS tests (
    id      INTEGER      AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255)                NOT NULL,
    PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE test;