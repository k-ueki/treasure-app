
-- +goose Up
CREATE TABLE iine (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  idea_id INT UNSIGNED NOT NULL,
  user_id INT UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY(idea_id) REFERENCES idea(id),
  FOREIGN KEY(user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE iine;
