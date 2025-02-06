CREATE DATABASE IF NOT EXISTS twitch-genie;

CREATE TABLE IF NOT EXISTS users (
  user_id varchar(255) PRIMARY KEY,
  login varchar(255) NOT NULL,
  access_token text,
  refresh_token text
)
