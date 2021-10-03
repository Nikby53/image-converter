CREATE database if not exists image.converter

CREATE TABLE if not exists users
(
    id serial primary key,
    email varchar(255) unique not null,
    password varchar(255) not null
);

CREATE TABLE if not exists images
(
    id serial primary key,
    name varchar(255) not null,
    format varchar(255) not null
);

CREATE TABLE if not exists request
(
    id serial primary key,
    user_id int not null,
    image_id int not null,
    target_id int,
    filename varchar(255) not null,
    status varchar(255) not null,
    created timestamp without time zone default current_timestamp not null,
    updated timestamp without time zone default current_timestamp not null,
    sourceFormat varchar(255) not null,
    targetFormat varchar(255) not null,
    ratio int,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (image_id) REFERENCES images(id)
);