
-- +migrate Up
CREATE TABLE IF NOT EXISTS `manuals` (
    `id`             BIGINT                UNSIGNED AUTO_INCREMENT,
    `title`          VARCHAR(255)          NOT NULL,
    `description`    TEXT                  NOT NULL,
    `tag_id`         BIGINT                ,
    `created_at`     DATETIME              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP             NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- +migrate Down
DROP TABLE IF EXISTS `manuals`;
