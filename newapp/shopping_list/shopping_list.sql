-- Create the shopping_list database
CREATE DATABASE IF NOT EXISTS shopping_list;
USE shopping_list;

-- Shopping list items
CREATE TABLE IF NOT EXISTS shopping_list_items (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    ingredient_id BIGINT NOT NULL,
    ingredient VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- Index for faster lookups
CREATE INDEX idx_user_shopping_list ON shopping_list_items (user_id);
