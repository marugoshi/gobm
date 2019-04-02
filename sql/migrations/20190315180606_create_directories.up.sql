CREATE TABLE directories (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  parent_id BIGINT UNSIGNED,
  name varchar(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY(id),
  INDEX(parent_id),
  INDEX(name),
  INDEX(created_at),
  INDEX(updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
