create database if not exists red-social;
use red-social;

drop table if exists usuarios;

create table usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha text not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;