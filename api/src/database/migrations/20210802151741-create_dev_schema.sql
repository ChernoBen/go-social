
-- +migrate Up
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    nick VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(150) NOT NULL,
    createdat timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers (
    user_id INT NOT NULL,
    FOREIGN kEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    follower_id INT,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    PRIMARY KEY (user_id,follower_id)
) ENGINE=INNODB;

-- +migrate Down
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;