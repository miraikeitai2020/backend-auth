---- create database ----
DROP DATABASE IF EXISTS user;
CREATE DATABASE user;

---- create user ----
CREATE USER IF NOT EXISTS 'miraiketai2020'@'%' IDENTIFIED BY 'miraiketai2020';
GRANT ALL ON user.* TO miraiketai2020;

---- create tables ----
USE michishirube;
DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users`
(
 `id`               VARCHAR(256),
 `pass`             VARCHAR(256) NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;