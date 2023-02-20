SET NAMES 'utf8mb4';

USE testDB;

CREATE TABLE `solitaire_1m`
(
    `init_cards`      varchar(512) NOT NULL COMMENT 'init cards key',
    `solution_count`  int          NOT NULL COMMENT 'solution count',
    `average_step`    int          NOT NULL COMMENT 'average step',
    `diff_step`       int          NOT NULL COMMENT 'average step difference',
    `min_step`        int          NOT NULL COMMENT 'min step',
    `max_step`        int          NOT NULL COMMENT 'max step',
    `max_stack_size`  int          NOT NULL COMMENT 'max stack size',
    `max_search_size` int          NOT NULL COMMENT 'max search size',
    `search_count`    int          NOT NULL COMMENT 'search count',
    `end_proc_reason` int          NOT NULL COMMENT 'end proc reason (0-normal -1 -- max stack size -2 -- max search size)',
    PRIMARY KEY (`init_cards`),
    INDEX `solution_count` (`solution_count`),
    INDEX `average_step` (`average_step`),
    INDEX `min_step` (`min_step`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;
