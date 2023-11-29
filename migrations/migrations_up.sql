create table if not exists users (
    userID serial primary key,
    fname varchar(64) not null,
    lname varchar (64) not null,
    age int not null,
    email varchar(64) not null,
    passwordHash varchar (64) not null
);