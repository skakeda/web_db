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

DROP SCHEMA IF EXISTS myamagata;
CREATE SCHEMA myamagata_db;
USE myamagata_db;

DROP TABLE IF EXISTS data;

CREATE TABLE data
(
  id           INT(10),
  txt     VARCHAR(40)
);

INSERT INTO data (id, txt) VALUES (1, "I like tabble tennis");
INSERT INTO data (id, txt) VALUES (2, "I like game");

DROP SCHEMA IF EXISTS syatani;
CREATE SCHEMA syatani;
USE syatani;


DROP TABLE IF EXISTS data;

CREATE TABLE data
(
  id           INT(10),
  txt     VARCHAR(40)
);


INSERT INTO data (id, txt) VALUES (1, "YATANI");
INSERT INTO data (id, txt) VALUES (2, "XXXXXXXXXXXXXXXXX");


CREATE SCHEMA tozono_db;
USE tozono_db;

INSERT INTO data (id, txt) VALUES (1, "tozono!");
INSERT INTO data (id, txt) VALUES (2, "tozono!!");


CREATE SCHEMA yichikawa;
USE yichikawa;


DROP TABLE IF EXISTS data;

CREATE TABLE data
(
  id           INT(10),
  txt     VARCHAR(40)
);


INSERT INTO data (id, txt) VALUES (1, "yichikawa");
INSERT INTO data (id, txt) VALUES (2, "it the hard way");