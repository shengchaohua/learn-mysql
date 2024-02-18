# Create
CREATE TABLE IF NOT EXISTS `join_index_1_tab`
(
    `id`       BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `category` VARCHAR(100) NOT NULL,
    `uid`      BIGINT UNSIGNED NOT NULL,
    UNIQUE INDEX `idx_category_uid` (`category`, `uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS `join_index_1_tab`
(
    `id`       BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `category` VARCHAR(100) NOT NULL,
    `uid`      BIGINT UNSIGNED NOT NULL,
    UNIQUE INDEX `idx_uid_category_uid` (`uid`, `category`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


# Drop
DROP TABLE join_index_1_tab;
DROP TABLE join_index_2_tab;

# Copy data
INSERT INTO join_index_2_tab SELECT * FROM join_index_1_tab order by id;

# Query
SELECT * FROM join_index_1_tab order by id limit 10;
