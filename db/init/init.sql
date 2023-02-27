DROP SCHEMA IF EXISTS test;
CREATE SCHEMA test;
USE test;

DROP TABLE IF EXISTS data;

CREATE TABLE data
(
  id           INT(10),
  txt     VARCHAR(40)
);

INSERT INTO data (id, txt) VALUES (1, "A thousand Miles");
INSERT INTO data (id, txt) VALUES (2, "it the hard way");
