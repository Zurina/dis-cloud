create database test;
use test;

CREATE TABLE uuid
(
id INTEGER AUTO_INCREMENT,
uuid TEXT,
PRIMARY KEY (id)
);

INSERT INTO uuid (uuid) VALUES ("123e4567-e89b-12d3-a456-426614174000")