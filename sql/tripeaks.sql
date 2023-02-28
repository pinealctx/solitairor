SET NAMES 'utf8mb4';

USE testDB;

CREATE TABLE `tripeaks`
(
    `id`    bigint  NOT NULL COMMENT 'id key',
    `road`  int  NOT NULL COMMENT 'solution road count',
    `step`  int  NOT NULL COMMENT 'min step',
    `cards` blob NOT NULL COMMENT 'cards',

    PRIMARY KEY (`id`),
    INDEX `road_step` (`road`, `step`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE `tripeakgate`
(
    `id`         bigint  NOT NULL COMMENT 'id key',
    `difficulty` int  NOT NULL COMMENT 'difficulty',
    `cards`      blob NOT NULL COMMENT 'cards',
    `road`       int  NOT NULL COMMENT 'solution road count',
    `step`       int  NOT NULL COMMENT 'min step',

    PRIMARY KEY (`id`),
    INDEX `difficulty` (`difficulty`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;
