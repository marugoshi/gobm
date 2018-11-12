DROP DATABASE IF EXISTS gobm_d;
CREATE DATABASE gobm_d CHARACTER SET utf8mb4;
USE gobm_d;

CREATE TABLE directories (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  parent_id BIGINT UNSIGNED,
  name varchar(255) NOT NULL,
  PRIMARY KEY(id),
  INDEX(parent_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE bookmarks (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  directory_id BIGINT UNSIGNED,
  url TEXT NOT NULL,
  title TEXT NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(url(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO bookmarks (url, title) VALUES("https://www.makototokuyama.com", "MAKOTO TOKUYAMA");
INSERT INTO bookmarks (url, title) VALUES("https://t.makototokuyama.com", "T MAKOTO TOKUYAMA");
