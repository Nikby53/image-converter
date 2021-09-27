CREATE SCHEMA if not exists image.convert

CREATE TABLE if not exists users
(
    id serial not null unique,
    email varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE if not exists images
(
    id serial not null unique,
    name varchar(255) not null,
    format varchar(255) not null
);

CREATE TABLE if not exists request
(
    id serial not null,
    user_id int references users(id) not null,
    images_id int references images(id) not null,
    status varchar(255),
    created timestamp without time zone default current_timestamp not null,
    updated timestamp without time zone default current_timestamp not null,
    sourceFormat varchar(255) not null,
    targetFormat varchar(255) not null,
    ratio int not null,
    filename varchar(255) not null
);