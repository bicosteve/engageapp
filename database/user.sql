CREATE TABLE `user` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT, 
    `email` VARCHAR(255) NOT NULL, 
    `password_hash` VARCHAR(255) NOT NULL, 
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;