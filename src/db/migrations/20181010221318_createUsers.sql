
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users`(
    `id` INT (11) NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(256) NOT NULL,
    `last_name` VARCHAR(256) NOT NULL,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;