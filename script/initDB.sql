CREATE DATABASE example;

\c example;

CREATE TABLE IF NOT EXISTS users (
    id INT GENERATED ALWAYS AS IDENTIFIED PRIMARY KEY,
    name VARCHAR(256),
    age INT,
    password VARCHAR(256)
);

INSERT INTO users (name , age , password) VALUES ('tiramis',19,'tiramis2334'),('neko',21,'nekopara');
 