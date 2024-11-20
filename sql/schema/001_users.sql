-- +goose Up
CREATE TABLE Users (
  ID UUID PRIMARY KEY,           
  name VARCHAR(255) UNIQUE NOT NULL, 
  created_at TIMESTAMP NOT NULL,  
  updated_at TIMESTAMP NOT NULL   
);

-- +goose Down
DROP TABLE IF EXISTS Users;