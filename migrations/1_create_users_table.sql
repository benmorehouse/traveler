-- 1_create_users_table.sql
-- +migrate Up
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE users;
