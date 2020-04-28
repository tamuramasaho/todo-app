
-- +migrate Up
CREATE TABLE IF NOT EXISTS `todos` (
    `id`             BIGINT                UNSIGNED AUTO_INCREMENT,
    `title`          VARCHAR(255)          NOT NULL,
    `due_date`       DATETIME              NOT NULL,
    `remind_at`      DATETIME              NOT NULL, 
    `should_remind`  BOOLEAN               NOT NULL,
    `created_at`     DATETIME              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP             NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

-- +migrate Down
DROP TABLE IF EXISTS `todos`;
