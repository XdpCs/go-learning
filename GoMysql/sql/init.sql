CREATE DATABASE IF NOT EXISTS goTestMysql CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE goTestMysql;
CREATE TABLE `person`
(
    `user_id`  int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(260) DEFAULT NULL,
    `sex`      varchar(260) DEFAULT NULL,
    `email`    varchar(260) DEFAULT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8;

CREATE TABLE place
(
    country varchar(200),
    city    varchar(200),
    telcode int
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8;