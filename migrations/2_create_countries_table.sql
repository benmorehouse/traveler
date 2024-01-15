-- +migrate Up
CREATE TABLE countries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    region VARCHAR(255),
    language VARCHAR(255),
);

-- +migrate Down
DROP TABLE countries;
