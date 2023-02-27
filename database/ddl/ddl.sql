-- Table: Utilisateur
CREATE TABLE IF NOT EXISTS users (
    id_user INTEGER NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(25) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(25) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_user)
);

-- Table: Token
CREATE TABLE IF NOT EXISTS Token {
    id_token INTEGER NOT NULL AUTO_INCREMENT,
    id_user INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_token),
    FOREIGN KEY (id_user) REFERENCES User(id_user)
};

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
    id_kind INT NOT NULL,
    id_user INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_shop),
    FOREIGN KEY (id_kind) REFERENCES kind(id_kind),
    FOREIGN KEY (id_user) REFERENCES users(id_user)
);

-- Table: kind d'étalissement (ex: coiffeur, barbier, tatoueur, etc.)
CREATE TABLE IF NOT EXISTS kinds (
    id_kind INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: Horaires d'ouverture
CREATE TABLE IF NOT EXISTS openingHours (
    id INT PRIMARY KEY NOT NULL,
    day INT NOT NULL,
    id_shop INTEGER NOT NULL,
    open VARCHAR(255) NOT NULL,
    close VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_shop) REFERENCES shops(id_shop)
);

-- Table: Prestation
CREATE TABLE IF NOT EXISTS Benefit (
    id_benefit SERIAL PRIMARY KEY NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,
    price VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_shop) REFERENCES Shop(id_shop)
);

-- Table: Collaborateur de l'établissement (les employés qui seront associés a une prestation ou choisie par le client)
CREATE TABLE IF NOT EXISTS collaborators (
    id_collaborator INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(id_shop) REFERENCES shops(id_shop)
);

-- Table: Avis
CREATE TABLE IF NOT EXISTS Review (
    id_review SERIAL PRIMARY KEY NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    id_user INTEGER NOT NULL,
    rating INTEGER NOT NULL,
    comment VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_shop) REFERENCES Shop(id_shop),
    FOREIGN KEY (id_user) REFERENCES User(id)
);

-- Table: Réservation
CREATE TABLE IF NOT EXISTS Reservation (
    id_reservation SERIAL PRIMARY KEY NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    id_user INTEGER NOT NULL,
    id_benefit INTEGER NOT NULL,
    date VARCHAR(255) NOT NULL,
    time VARCHAR(255) NOT NULL,
    comment VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_shop) REFERENCES Shop(id_shop),
    FOREIGN KEY (id_user) REFERENCES User(id),
    FOREIGN KEY (id_benefit) REFERENCES Benefit(id_benefit)
);