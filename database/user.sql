CREATE TABLE `user` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT, 
    `email` VARCHAR(255) NOT NULL, 
    `password` VARCHAR(255) NOT NULL, 
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY_KEY (`user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10 DEFAULT CHARSET = utf8;