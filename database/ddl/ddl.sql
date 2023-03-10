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

CREATE TABLE IF NOT EXISTS kinds (
    id_kind INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: Etablissement
CREATE TABLE IF NOT EXISTS shops (
    id_shop INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    -- latitude VARCHAR(255) NOT NULL,
    -- longitude VARCHAR(255) NOT NULL,
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

-- Table: Collaborateur de l'??tablissement (les employ??s qui seront associ??s a une prestation ou choisie par le client)
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

-- Table: R??servation
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

/* 
INSERT INTO users (firstName, lastName, phone, email, password, role) VALUES
('Paul', 'Smith', '0737485936', 'paul.smith@gmail.com', 'password', 'trader'),
('Marie', 'Dubois', '0737482875', 'marie.dubois@gmail.com', 'password', 'trader'),
('John', 'Doe', '0634094877', 'john.doe@gmail.com', 'password', 'customer'),
('Carl', 'Johnson', '0648627756', 'carl.johnson@gmail.com', 'password', 'customer'),
('Sophie', 'Laverre', '0648107569', 'sophie.laverre@gmail.com', 'password', 'customer');

INSERT INTO kinds (name) VALUES 
('Coiffeurs'),
('Barbiers'),
('Manucure'),
('Instituts de beaut??'),
('Tatoueurs');

INSERT INTO shops (name, city, latitude, longitude, country, zip_code, phone, email, id_user, description) VALUES
('Defosso Coiffure', 'Saint Denis', '48.936181', '2.357443', 'France', '93200', '0158569754', 'defosso-coiffure@gmail.com', 1, 'Salon de coiffure'),
('Prestige Barber', 'Paris', '48.50N', '2.20E', 'France', '75017', '0158583754', 'prestige-barber@outlook.fr', 1, 'Barbier'),
('Beautiful nails', 'Rosny sous Bois', '48.875661', '2.485932', 'France', '93064', '0158583128', 'beautiful-nails@gmail.com', 1, 'Salon de manucure'),
('House of beauty', 'Montpellier', '43.610769', '3.876716', 'France', '34000', '0158583128', 'house-of-beauty@gmail.com', 1, 'Institut de beaut??'),
('Ink Place', 'Bordeaux', '44.837789', '-0.579180', 'France', '330003', '0158583128', 'ink-place@gmail.com', 1, 'Salon de tatouage');

INSERT INTO shop_kind (id_shop, id_kind) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 2),
(5, 1);

INSERT INTO benefits (id_shop, name, description, duration, price) VALUES
(1, 'Coupe de cheveux homme', 'Coupe de cheveux compl??te par exemple d??grad?? avec contours', '20 minutes', '10'),
(1, 'Coupe de cheveux femme', "Coupe de cheveux femme ?? la demande de la cliente", '30 minutes', '40'),
(1, 'Shampoing et coiffage homme', "Coupe de cheveux et shampoing pour homme", '35 minutes', '40'),
(2, 'Taille barbe rasoir', 'Taille de la barbe avec rituel serviette chaude, huiles essentielles et cr??me de soin', '30 minutes', '20'),
(2, 'Taille barbe tondeuse', 'Taille de la barbe avec une tondeuse de pr??cision pour barbe', '10 minutes', '15'),
(3, 'Manucure Spa', 'Soin des mains avec mise en beaut??', '20 minutes', '20'),
(3, 'Manucure Spa + pose des ongles', 'Soin des mains avec mise en beaut??, pose, polissage et limage des ongles', '30 minutes', '30'),
(3, 'P??dicure', 'Beaut?? des pieds Spa, p??dicure simple', '30 minutes', '40'),
(3, 'P??dicure + pose de vernis', 'Beaut?? des pieds Spa, p??dicure et pose de vernis French', '40 minutes', '45'),
(4, 'Extension ongles gel acrylique', 'Pose compl??te gel acrylique + pose de vernis normal', '50 minutes', '40 '),
(4, 'Extension ongles', 'Pose compl??te r??sine sans vernis', '30 minutes', '25'),
(5, 'Rendez vous tatouage 1 heure', "Le montant indiqu?? correspond ?? l'acompte le paiement int??gral se fera sur place", '1 heure', '80'),
(5, 'Rendez vous tatouage 2 heures', "Le montant indiqu?? correspond ?? l'acompte le paiement int??gral se fera sur place", '2 heures', '80'),
(1, 'Contours', "Contours d'une coupe de cheveux seulement", '5 minutes', '5');
*/


SELECT shops.id_shop, shops.name, address, zip_code, city, latitude, longitude, country, phone, email, description, ST_Distance_Sphere( point (2.3483915,48.8534951), point(longitude, latitude)) * .000621371192 AS distance_in_miles FROM shops INNER JOIN shop_kind ON shops.id_shop = shop_kind.id_shop INNER JOIN kinds  ON shop_kind.id_kind = kinds.id_kind WHERE kinds.name = "Coiffeurs" having distance_in_miles <= 15 order by distance_in_miles asc;