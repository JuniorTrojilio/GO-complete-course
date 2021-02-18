CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(255) not null unique,
    password varchar(50) not null,
    createat timestamp default current_timestamp()
) ENGINE=INNODB;