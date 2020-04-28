
-- +migrate Up
CREATE TABLE IF NOT EXISTS `groups` (
    `id`                 BIGINT                UNSIGNED AUTO_INCREMENT,
    `name`               VARCHAR(255)          NOT NULL,
    `description`        VARCHAR(255)          NOT NULL,
    `parent_group_id`    BIGINT                ,
    `created_at`         DATETIME              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`         TIMESTAMP             NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- +migrate Down
DROP TABLE IF EXISTS `groups`;
