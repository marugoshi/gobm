DROP DATABASE IF EXISTS gobm;
CREATE DATABASE gobm_d CHARACTER SET utf8mb4;
USE gobm_d
CREATE TABLE bookmarks (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  url TEXT NOT NULL,
  title TEXT NOT NULL,
  memo TEXT,
  PRIMARY KEY(id),
  UNIQUE(url(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO bookmarks (url, title, memo) VALUES("https://www.makototokuyama.com", "MAKOTO TOKUYAMA", "test");
