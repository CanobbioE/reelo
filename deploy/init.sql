-- BACKUP DATABASE reelo
--  TO DISK = ''
--  WITH DIFFERENTIAL

GRANT ALL PRIVILEGES ON reelo.* TO 'reeloUser'@'localhost' IDENTIFIED BY 'password';


CREATE TABLE Giocatore (
	id int AUTO_INCREMENT,
	nome varchar(255) NOT NULL,
	cognome varchar(255) NOT NULL,
	reelo int,
	PRIMARY KEY (id)
);

CREATE TABLE Giochi (
	id int AUTO_INCREMENT,
	anno int,
	categoria varchar(4),
	PRIMARY KEY (id, anno, categoria)
);

CREATE TABLE Risultato (
	id int AUTO_INCREMENT,
	tempo int,
	esercizi int,
	punteggio int,
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

INSERT INTO Costanti (
	anno_inizio,
	k_esercizi,
	finale,
	fattore_moltiplicativo,
	exploit,
	no_partecipazione
) VALUES (2002, 20.0, 1.5, 10000.0, 0.9, 0.9);

INSERT INTO Utenti (nomeutente, parolachiave)
VALUES ('admin@reelo.it', 'b133a0c0e9bee3be20163d2ad31d6248db292aa6dcb1ee087a2aa50e0fc75ae2');

