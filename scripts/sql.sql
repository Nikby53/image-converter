CREATE database if not exists image.converter

CREATE TABLE if not exists users
(
    id serial unique not null ,
    email varchar(255) unique not null,
    password varchar(255) not null
);

CREATE TABLE if not exists images
(
    id serial unique not null,
    name varchar(255) not null,
    format varchar(255) not null
);

CREATE TABLE if not exists request
(
    id serial unique not null,
    user_id int references users(id) not null,
    images_id int references images(id) not null,
    target_id serial unique not null,
    filename varchar(255) not null,
    status varchar(255),
    created timestamp without time zone default current_timestamp not null,
    updated timestamp without time zone default current_timestamp not null,
    sourceFormat varchar(255) not null,
    targetFormat varchar(255) not null,
    ratio int
);