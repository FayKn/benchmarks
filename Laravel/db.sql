# Delete the database if it exists to reset the database
DROP DATABASE IF EXISTS `SteamUsers`;

CREATE DATABASE `SteamUsers` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `SteamUsers`;

CREATE TABLE if not exists `SteamUsers`
(
    `VanityName` varchar(255) NOT NULL,
    `SteamID`    varchar(255) NULL,
    PRIMARY KEY (`VanityName`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE if not exists `APIsCalled`
(
    `Api_url` varchar(255) NOT NULL,
    `count`   int(11)      NOT NULL,
    PRIMARY KEY (`Api_url`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO  `SteamUsers` (`VanityName`)
VALUES ('https://steamcommunity.com/id/knol_gamer/'),
       ('https://steamcommunity.com/id/SilverHazel'),
       ('https://steamcommunity.com/id/gabelogannewell'),
       ('https://steamcommunity.com/id/olbmaphlee/'),
       ('https://steamcommunity.com/id/SandsVR/'),
       ('https://steamcommunity.com/id/valvesoftware');