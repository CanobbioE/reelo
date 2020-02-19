-- BACKUP DATABASE reelo
--  TO DISK = ''
--  WITH DIFFERENTIAL

GRANT ALL PRIVILEGES ON reelo.* TO 'reeloUser'@'localhost' IDENTIFIED BY 'password';

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `Costanti`;
CREATE TABLE `Costanti` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `anno_inizio` int(11) DEFAULT NULL,
  `k_esercizi` float DEFAULT NULL,
  `finale` float DEFAULT NULL,
  `fattore_moltiplicativo` float DEFAULT NULL,
  `exploit` float DEFAULT NULL,
  `no_partecipazione` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `Giocatore`;
CREATE TABLE `Giocatore` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar(255) NOT NULL,
  `cognome` varchar(255) NOT NULL,
  `reelo` float DEFAULT NULL,
  `accent` varchar(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `Giochi`;
CREATE TABLE `Giochi` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `anno` int(11) DEFAULT NULL,
  `categoria` varchar(4) DEFAULT NULL,
  `inizio` int(11) DEFAULT NULL,
  `fine` int(11) DEFAULT NULL,
  `internazionale` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `duplicate_giochi` (`anno`,`categoria`,`internazionale`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `Partecipazione`;
CREATE TABLE `Partecipazione` (
  `giocatore` int(11) NOT NULL,
  `giochi` int(11) NOT NULL,
  `risultato` int(11) DEFAULT NULL,
  `sede` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`giocatore`,`giochi`),
  KEY `FK_partecipazione_giochi` (`giochi`),
  KEY `FK_partecipazione_risultato` (`risultato`),
  CONSTRAINT `FK_partecipazione_giochi` FOREIGN KEY (`giochi`) REFERENCES `Giochi` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FK_partecipazione_risultato` FOREIGN KEY (`risultato`) REFERENCES `Risultato` (`id`) ON DELETE CASCADE,
  CONSTRAINT `Partecipazione_ibfk_5` FOREIGN KEY (`giocatore`) REFERENCES `Giocatore` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `Risultato`;
CREATE TABLE `Risultato` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tempo` int(11) DEFAULT NULL,
  `esercizi` int(11) DEFAULT NULL,
  `punteggio` int(11) DEFAULT NULL,
  `posizione` int(11) DEFAULT NULL,
  `pseudo_reelo` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `Utenti`;
CREATE TABLE `Utenti` (
  `nomeutente` varchar(255) NOT NULL,
  `parolachiave` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`nomeutente`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
) VALUES (2002, 20.0, 0, 10000.0, 0.9, 0.9);

INSERT INTO Utenti (nomeutente, parolachiave)
VALUES ('admin@reelo.it', '5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8');

