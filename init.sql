-- 建表
CREATE TABLE IF NOT EXISTS `ip_history`(
   `id`  INT AUTO_INCREMENT,
   `created_at` DATE NOT NULL,
   `updated_at` DATE NOT NULL,
    `ip` VARCHAR(255) NOT NULL,
    `message` VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
) CHARSET=utf8;