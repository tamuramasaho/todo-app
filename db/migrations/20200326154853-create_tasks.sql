
-- +migrate Up
CREATE TABLE IF NOT EXISTS `tasks` (
    `id`             BIGINT                UNSIGNED AUTO_INCREMENT,
    `title`          VARCHAR(255)          NOT NULL,
    `date`           VARCHAR(255)          NOT NULL,
    `remaind_at`     DATETIME              NOT NULL, 
    `remainder`      BOOLEAN               NOT NULL,
    `created_at`     DATETIME              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP             NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- +migrate Down
DROP TABLE IF EXISTS `tasks`;