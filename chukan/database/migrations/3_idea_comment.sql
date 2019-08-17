
-- +goose Up
CREATE TABLE idea_comment (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  body TEXT,
  idea_id INT UNSIGNED NOT NULL,
  user_id INT UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY(user_id) REFERENCES user(id),
  FOREIGN KEY(idea_id) REFERENCES idea(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE idea_comment;