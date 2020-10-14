package services

const createUsersDDL  = `create table if not exists users (
    id bigserial primary key not null,
	name varchar(30) not null,
	surname varchar(30) not null,
	lastname varchar(30),
	login varchar(30) not null unique,
	password text not null,
	phone varchar(30) not null,
	role varchar(30) not null default 'user',
	status boolean default false,
	position varchar(30) not null,
	status_line boolean default false
);`

const createStatesDDL = `create table if not exists states (
	id bigserial primary key not null,
	user_id bigint not null references users (id),
	work_time integer not null,
	status boolean not null,
	unix_date bigint not null,
	time_date timestamp
);`

const createFixLogTimeDDL = `create table if not exists login_times (
	id bigserial primary key not null,
	user_id bigint not null references users (id),
	day_date bigint not null,
	login_date text[] not null,
	logout_date text[] not null
);`

const createActivitiesDDL = `create table if not exists activities (
	id bigserial primary key not null,
	user_id bigint not null references users (id),
	token string not null,
	unix_time bigint not null,
	status bool not null, 
	work_time bigint not null,
	exited bool not null
);`