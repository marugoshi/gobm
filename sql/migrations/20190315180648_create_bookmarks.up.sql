CREATE TABLE bookmarks (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  directory_id BIGINT UNSIGNED,
  url TEXT NOT NULL,
  title TEXT NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(url(255))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;