CREATE TABLE IF NOT EXISTS `join_index_1_tab`
(
    `id`       BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `category` VARCHAR(100)    NOT NULL,
    `uid`      INT UNSIGNED NOT NULL,
    UNIQUE INDEX `idx_category_uid` (`category`, `uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS `join_index_2_tab`
(
    `id`       BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `category` VARCHAR(100)    NOT NULL,
    `uid`      INT UNSIGNED NOT NULL,
    UNIQUE INDEX `idx_uid_category` (`uid`, `category`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
