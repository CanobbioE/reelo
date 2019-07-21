-- BACKUP DATABASE reelo
--  TO DISK = ''
--  WITH DIFFERENTIAL

GRANT ALL PRIVILEGES ON reelo.* TO 'reeloUser'@'localhost' IDENTIFIED BY 'password';


CREATE TABLE Giocatore (
	id int AUTO_INCREMENT,
	nome varchar(255) NOT NULL,
	cognome varchar(255) NOT NULL,
	reelo float,
	PRIMARY KEY (id)
);

CREATE TABLE Giochi (
	id int AUTO_INCREMENT,
	anno int,
	categoria varchar(4),
	inizio int,
	fine int,
	PRIMARY KEY (id),
	CONSTRAINT duplicate_giochi UNIQUE (anno, categoria)
);

CREATE TABLE Risultato (
	id int AUTO_INCREMENT,
	tempo int,
	esercizi int,
	punteggio int,
	posizione int,
	pseudo_reelo float,
	PRIMARY KEY (id)
);

CREATE TABLE Partecipazione (
	giocatore int,
	giochi int,
	risultato int,
	sede varchar(255),
	PRIMARY KEY (giocatore, giochi),
	FOREIGN KEY (giocatore) REFERENCES Giocatore(id) ON UPDATE CASCADE,
	FOREIGN KEY (giochi) REFERENCES Giochi(id) ON UPDATE CASCADE,
	FOREIGN KEY (risultato) REFERENCES Risultato(id) ON DELETE SET NULL
);

CREATE TABLE Utenti (
	nomeutente varchar(255),
	parolachiave varchar(255),
	PRIMARY KEY (nomeutente)
);

CREATE TABLE Costanti (
	id int AUTO_INCREMENT,
	anno_inizio int,
	k_esercizi float,
	finale float,
	fattore_moltiplicativo float,
	exploit float,
	no_partecipazione float,
	PRIMARY KEY (id)
);

-- INIT VALUES
SET collation_connection = 'utf8_general_ci';

ALTER DATABASE reelo CHARACTER SET utf8 COLLATE utf8_general_ci;

ALTER TABLE  Giochi CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE Giocatore CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE Risultato CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE Utenti CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE Costanti CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;
ALTER TABLE Partecipazione CONVERT TO CHARACTER SET utf8 COLLATE utf8_general_ci;


INSERT INTO Costanti (
	anno_inizio,
	k_esercizi,
	finale,
	fattore_moltiplicativo,
	exploit,
	no_partecipazione
) VALUES (2002, 20.0, 1.5, 10000.0, 0.9, 0.9);

INSERT INTO Utenti (nomeutente, parolachiave)
VALUES ('admin@reelo.it', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8');

