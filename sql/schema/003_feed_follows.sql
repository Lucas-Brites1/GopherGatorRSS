-- +goose Up
CREATE TABLE Feed_Follows (
  ID UUID PRIMARY KEY,
  user_id UUID REFERENCES Users(ID) ON DELETE CASCADE,
  feed_id UUID REFERENCES Feeds(ID) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,  
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT unique_user_feed UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE IF EXISTS Feed_Follows;