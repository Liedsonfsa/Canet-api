CREATE DATABASE IF NOT EXISTS canet;

USE canet;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;


CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key (usuario_id, seguidor_id)
) ENGINE=INNODB;

insert into seguidores (usuario_id, seguidor_id)
values
(2, 3),
(2, 4),
(4, 3);

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang'
GRANT ALL PRIVILEGES ON canet.* TO 'golang'@'localhost'

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;

insert into publicacoes (titulo, conteudo, autor_id) values
("Olá mundo", "Essa é a minha primeira publicação aqui", 2),
("Algém aqui", "Alguém paea fazer amizade?", 3),
("Reyes de Europa", "Espanha campeã da euro.", 4);