package services

const createUsersDDL  = `create table if not exists users (
    id bigserial primary key not null,
	name varchar(30) not null,
	surname varchar(30) not null,
	lastname varchar(30),
	login varchar(30) not null unique,
	password text not null,
	phone varchar(30) not null,
	role varchar(30) not null
);`
