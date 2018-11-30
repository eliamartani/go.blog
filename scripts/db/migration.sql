CREATE DATABASE `godb`

CREATE TABLE `godb`.`category` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(50),
  `Active` TINYINT(1) NULL,
  PRIMARY KEY (`ID`));

CREATE TABLE `godb`.`post` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `CategoryID` INT NOT NULL,
  `Title` VARCHAR(150) NULL,
  `Url` VARCHAR(150) NULL,
  `Description` VARCHAR(500) NULL,
  `ImageUrl` VARCHAR(200) NULL,
  `Content` TEXT NULL,
  `Author` VARCHAR(50) NULL,
  `DateCreation` DATETIME NOT NULL,
  `DatePublished` DATETIME NULL,
  `Active` TINYINT(1) NOT NULL,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (`CategoryID`) REFERENCES `godb`.`category`(`ID`) );

CREATE TABLE `godb`.`tag` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(50),
  PRIMARY KEY (`ID`));

CREATE TABLE `godb`.`post_tag` (
  `PostID` INT NOT NULL,
  `TagID` INT NOT NULL,
  PRIMARY KEY (`PostID`, `TagID`),
  FOREIGN KEY (`PostID`) REFERENCES `godb`.`post`(`ID`),
  FOREIGN KEY (`TagID`) REFERENCES `godb`.`tag`(`ID`)
);
