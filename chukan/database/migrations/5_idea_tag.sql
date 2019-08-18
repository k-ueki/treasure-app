
-- +goose Up
CREATE TABLE idea_tag (
  id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  idea_id INT UNSIGNED NOT NULL,
  tag_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY(idea_id) REFERENCES idea(id),
  FOREIGN KEY(tag_id) REFERENCES tag(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DROP TABLE idea_tag;
