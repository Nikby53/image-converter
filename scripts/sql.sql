CREATE SCHEMA if not exists image.convert

CREATE TABLE if not exists users
(
    id serial not null unique,
    email varchar(255) not null,
    password varchar(255) not null
);