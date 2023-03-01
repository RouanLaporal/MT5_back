/*
Script de création de la base de données de test.
A noter, on utilise une stratégie avec DROP et IF NOT EXISTS afin de rendre 
notre script réutilisable dans le future, même si la base existe déjà
*/
create database IF NOT EXISTS MT5_back_project;

/* Créer l'utilisateur API */
create user IF NOT EXISTS 'api-dev'@'%.%.%.%' identified by 'api-dev-password';
grant select, update, insert, delete on challenge_project.* to 'api-dev'@'%.%.%.%';
grant trigger on challenge_project.* to 'api-dev'@'%.%.%.%';
flush privileges;

-- Table: Utilisateur
CREATE TABLE IF NOT EXISTS users (
    id_user INTEGER NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(50) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    phone VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(25) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_user)
);

-- Table: kind d'étalissement (ex: coiffeur, barbier, tatoueur, etc.)
CREATE TABLE IF NOT EXISTS kinds (
    id_kind INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO kinds (name) VALUES 
('Coiffeurs'),
('Barbiers'),
('Manucure'),
('Instituts de beauté'),
('Tatoueurs');

-- Table: Etablissement
CREATE TABLE IF NOT EXISTS shops (
    id_shop INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    latitude VARCHAR(255) NOT NULL,
    longitude VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    zip_code VARCHAR(255) NOT NULL,
    phone VARCHAR(25) NOT NULL,
    email VARCHAR(255) NOT NULL,
    id_user INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_shop),
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE ON UPDATE CASCADE
);

--Table: Lien Etablissement et Type
CREATE TABLE IF NOT EXISTS shop_kind(
    id INT NOT NULL AUTO_INCREMENT,
    id_shop INT NOT NULL,
    id_kind INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id_shop) REFERENCES shops(id_shop) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_kind) REFERENCES kinds(id_kind) ON DELETE CASCADE ON UPDATE CASCADE
);
