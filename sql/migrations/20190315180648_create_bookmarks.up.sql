CREATE TABLE bookmarks (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  directory_id BIGINT UNSIGNED,
  url TEXT NOT NULL,
  title TEXT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY(id),
  INDEX(directory_id),
  UNIQUE(url(255)),
  INDEX(title(255)),
  INDEX(created_at),
  INDEX(updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
