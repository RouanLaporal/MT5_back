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
INSERT INTO users (firstName, lastName, phone, email, password, role) VALUES
('Paul', 'Smith', '0737485936', 'paul.smith@gmail.com', 'password', 'admin'),
('Marie', 'Dubois', '0737482875', 'marie.dubois@gmail.com', 'password', 'admin'),
('John', 'Doe', '0634094877', 'john.doe@gmail.com', 'password', 'customer'),
('Carl', 'Johnson', '0648627756', 'carl.johnson@gmail.com', 'password', 'customer'),
('Sophie', 'Laverre', '0648107569', 'sophie.laverre@gmail.com', 'password', 'customer');

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
INSERT INTO shops (name, city, latitude, longitude, country, zip_code, phone, email, id_user, description) VALUES
('Defosso Coiffure', 'Saint Denis', '48.936181', '2.357443', 'France', '93200', '0158569754', 'defosso-coiffure@gmail.com', 1),
('Prestige Barber', 'Paris', '48.50N', '2.20E', 'France', '75017', '0158583754', 'prestige-barber@outlook.fr', 1),
('Beautiful nails', 'Rosny sous Bois', '48.875661', '2.485932', 'France', '93064', '0158583128', 'beautiful-nails@gmail.com', 1),
('House of beauty', 'Montpellier', '43.610769', '3.876716', 'France', '34000', '0158583128', 'house-of-beauty@gmail.com', 1),
('Ink Place', 'Bordeaux', '44.837789', '-0.579180', 'France', '330003', '0158583128', 'ink-place@gmail.com', 1);

--Table: Lien Etablissement et Type
CREATE TABLE IF NOT EXISTS shop_kind(
    id INT NOT NULL AUTO_INCREMENT,
    id_shop INT NOT NULL,
    id_kind INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id_shop) REFERENCES shops(id_shop) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_kind) REFERENCES kinds(id_kind) ON DELETE CASCADE ON UPDATE CASCADE
);
INSERT INTO shop_kind (id_shop, id_kind) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 2),
(5, 1);

-- Table: Prestation
CREATE TABLE IF NOT EXISTS benefits (
    id_benefit INTEGER NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,
    price VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id_benefit),
    FOREIGN KEY (id_shop) REFERENCES shops(id_shop) ON DELETE CASCADE ON UPDATE CASCADE
);
INSERT INTO benefits (id_shop, name, description, duration, price) VALUES
(1, 'Coupe de cheveux homme', 'Coupe de cheveux complète par exemple dégradé avec contours', '20 minutes', '10€'),
(1, 'Coupe de cheveux femme', "Coupe de cheveux femme à la demande de la cliente", '30 minutes', '40€'),
(1, 'Shampoing et coiffage homme', "Coupe de cheveux et shampoing pour homme", '35 minutes', '40€'),
(2, 'Taille barbe rasoir', 'Taille de la barbe avec rituel serviette chaude, huiles essentielles et crème de soin', '30 minutes', '20€'),
(2, 'Taille barbe tondeuse', 'Taille de la barbe avec une tondeuse de précision pour barbe', '10 minutes', '15€'),
(3, 'Manucure Spa', 'Soin des mains avec mise en beauté', '20 minutes', '20€'),
(3, 'Manucure Spa + pose des ongles', 'Soin des mains avec mise en beauté, pose, polissage et limage des ongles', '30 minutes', '30€'),
(3, 'Pédicure', 'Beauté des pieds Spa, pédicure simple', '30 minutes', '40€'),
(3, 'Pédicure + pose de vernis', 'Beauté des pieds Spa, pédicure et pose de vernis French', '40 minutes', '45€'),
(4, 'Extension ongles gel acrylique', 'Pose complète gel acrylique + pose de vernis normal', '50 minutes', '40 €'),
(4, 'Extension ongles', 'Pose complète résine sans vernis', '30 minutes', '25€'),
(5, "Rendez vous tatouage 1 heure', 'Le montant indiqué correspond à l'acompte le paiement intégral se fera sur place", '1 heure', '80€'),
(5, "Rendez vous tatouage 2 heures', 'Le montant indiqué correspond à l'acompte le paiement intégral se fera sur place", '2 heures', '80€'),
(1, 'Contours', "Contours d'une coupe de cheveux seulement", '5 minutes', '5€');

-- Table: Collaborateur de l'établissement (les employés qui seront associés a une prestation ou choisie par le client)
CREATE TABLE IF NOT EXISTS collaborator (
    id_collaborator INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(id_shop) REFERENCES shops(id_shop)
);

-- Table: Avis
CREATE TABLE IF NOT EXISTS reviews (
    id_review INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    id_user INTEGER NOT NULL,
    rating INTEGER NOT NULL,
    comment VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_shop) REFERENCES shops(id_shop) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE ON UPDATE CASCADE
);
INSERT INTO reviews (id_shop, id_user, rating, comment) VALUES
(1, 3, 5, "Très bon salon de coiffure je recommande"),
(1, 4, 4, "Bon salon et bonne prestation dans l'ensemble"),
(2, 4, 5, 'Très content du rasage avec le rasoir et employés chaleureux'),
(3, 5, 4, 'Très contente de ma manucure'),
(4, 5, 2, "Pose des ongles ratés l'employée était très désagréable"),
(5, 4, 4, "Très content de mon tatouage je recommande ce tatoueur");

-- Table: Réservation
CREATE TABLE IF NOT EXISTS reservations (
    id_reservation INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT,
    id_shop INTEGER NOT NULL,
    id_user INTEGER NOT NULL,
    id_benefit INTEGER NOT NULL,
    date VARCHAR(255) NOT NULL,
    time VARCHAR(255) NOT NULL,
    comment VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_shop) REFERENCES shops(id_shop) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (id_benefit) REFERENCES benefits(id_benefit) ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO reservations (id_shop, id_user, id_benefit, date, time) VALUES
(1, 3, 1, '2023-03-04', '13h'),
(2, 4, 2, '2023-03-04', '15h'),
(3, 5, 7, '2023-03-05', '11h'),
(4, 5, 10, '2023-03-05', '16h'),
(5, 4, 13, '2023-03-15', '12h30');
