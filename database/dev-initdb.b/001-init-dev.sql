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
    id_user INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(25) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_user)
);
