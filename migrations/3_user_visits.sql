-- +migrate Up
CREATE TABLE user_visits (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    country_id INT NOT NULL
);

-- +migrate Down
DROP TABLE user_visits;
