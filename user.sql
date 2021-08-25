CREATE TABLE users (
    id int not null primary key auto_increment,
    username varchar (64),
    first_name varchar(64),
    last_name varchar(64)
);

insert into users (username, first_name, last_name) values ('nombre', 'apellido', 'apellido')