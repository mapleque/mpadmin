DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user` (
    `id` INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `app_id` INT UNSIGNED NOT NULL,
    `user_id` INT UNSIGNED  NOT NULL,
    `create_at` DATETIME NOT NULL DEFAULT NOW()
) DEFAULT CHARSET=utf8;
ALTER TABLE `app_user` ADD UNIQUE(`app_id`), ADD INDEX(`user_id`);

