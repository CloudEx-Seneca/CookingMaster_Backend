CREATE DATABASE `usercenter`;
USE `usercenter`;

CREATE TABLE users (
                       id            BIGINT PRIMARY KEY,  -- Snowflake ID
                       email         VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,  -- Bcrypt or Argon2 hash
                       nickname      VARCHAR(100) UNIQUE,
                       sex           TINYINT CHECK (sex IN (0,1,2)),  -- 0: Unknown, 1: Male, 2: Female
                       avatar_url    VARCHAR(500),
                       info          TEXT,
                       created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                       INDEX idx_email (email),
                       INDEX idx_nickname (nickname)
);

CREATE TABLE user_tokens (
                             id         BIGINT PRIMARY KEY AUTO_INCREMENT,
                             user_id    BIGINT NOT NULL,
                             token      VARCHAR(500) NOT NULL,
                             type       TINYINT NOT NULL,  -- 1: Refresh, 2: Reset, 3: Register
                             status     TINYINT NOT NULL DEFAULT 1,  -- 1: Active, 2: Used, 3: Revoked, 4: Expired
                             expires_at TIMESTAMP NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                             INDEX idx_user_type (user_id, type),
                             FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
