SET NAMES 'utf8mb4';

USE testDB;

CREATE TABLE `klondike`
(
    `id`         bigint  NOT NULL COMMENT 'id key',
    `min_search` int     NOT NULL COMMENT 'min search step level',
    `step`       int     NOT NULL COMMENT 'min step',
    `road`       int     NOT NULL COMMENT 'solution road count',
    `cards`      blob    NOT NULL COMMENT 'cards',

    PRIMARY KEY (`id`),
    INDEX `mix` (`min_search`, `step`, `road`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE `klondikegate`
(
    `id`         bigint  NOT NULL COMMENT 'id key',
    `difficulty` int     NOT NULL COMMENT 'difficulty',
    `cards`      blob    NOT NULL COMMENT 'cards',
    `road`       int     NOT NULL COMMENT 'solution road count',
    `step`       int     NOT NULL COMMENT 'min step',
    `min_search` int     NOT NULL COMMENT 'min search step level',

    PRIMARY KEY (`id`),
    INDEX `difficulty` (`difficulty`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;
