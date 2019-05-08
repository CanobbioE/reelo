-- BACKUP DATABASE reelo
--  TO DISK = ''
--  WITH DIFFERENTIAL

DROP TABLE IF EXISTS Giocatore;
CREATE TABLE Giocatore (
	id int AUTO_INCREMENT,
	nome varchar(255) NOT NULL,
	cognome varchar(255) NOT NULL,
	reelo int,
	PRIMARY KEY (id)
);

DROP TABLE IF EXISTS Giochi;
CREATE TABLE Giochi (
	id int AUTO_INCREMENT,
	anno int,
	categoria varchar(4),
	PRIMARY KEY (id, anno, categoria)
);

DROP TABLE IF EXISTS Risultato;
CREATE TABLE Risultato (
	id int AUTO_INCREMENT,
	tempo int,
	esercizi int,
	punteggio int,
	PRIMARY KEY (id)
);

DROP TABLE IF EXISTS Partecipazione;
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

DROP TABLE IF EXISTS Utenti;
CREATE TABLE Utenti (
	nomeutente varchar(255),
	parolachiave varchar(255),
	PRIMARY KEY (nomeutente)
);
