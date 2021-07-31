CREATE DATABASE IF NOT EXISTS dev;
USE dev;
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    nick VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(150) NOT NULL,
    createdat timestamp default current_timestamp()
) ENGINE=INNODB;