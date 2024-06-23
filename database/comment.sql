CREATE TABLE `comment`(
    `comment_id` INT(11) NOT NULL AUTO_INCREMENT,
    `comment` VARCHAR(500) NOT NULL,
    `created_at` DATETIME DEFAULT  CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  	`user_id` INT NOT NULL, 
    `post_id` INT NOT NULL, 
    PRIMARY KEY (`comment_id`),
    FOREIGN KEY(`user_id`) REFERENCES user(`user_id`),
    FOREIGN KEY(`post_id`) REFERENCES post(`post_id`)
)ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;