-- +goose Up
CREATE TABLE Posts (
  ID UUID PRIMARY KEY UNIQUE,
  title VARCHAR(255),
  description VARCHAR(500),
  url VARCHAR(255) UNIQUE,
  published_at TIMESTAMP NOT NULL,
  feed_id UUID REFERENCES Feeds(ID) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS Posts;