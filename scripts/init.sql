CREATE DATABASE IF NOT EXISTS `evc`;

DROP TABLE IF EXISTS `evc`.`rates_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`rates_tab`
(
    `id`              BIGINT(20) UNSIGNED   NOT NULL auto_increment,
    `normal_rate`     DOUBLE(3, 3) UNSIGNED NOT NULL,
    `penalty_rate`    DOUBLE(3, 3) UNSIGNED NOT NULL,
    `no_show_penalty` DOUBLE(3, 3) UNSIGNED NOT NULL,
    `create_time`     INT(10) UNSIGNED      NOT NULL,
    `update_time`     INT(10) UNSIGNED      NOT NULL,
    `status`          INT(10) UNSIGNED      NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `evc`.`provider_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`provider_tab`
(
    `id`           BIGINT(20) UNSIGNED NOT NULL auto_increment,
    `company_name` VARCHAR(256)        NOT NULL,
    `description`  VARCHAR(256)        NOT NULL,
    `create_time`  INT(10) UNSIGNED    NOT NULL,
    `update_time`  INT(10) UNSIGNED    NOT NULL,
    `status`       INT(10) UNSIGNED    NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `evc`.`charger_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`charger_tab`
(
    `id`          BIGINT(20) UNSIGNED NOT NULL auto_increment,
    `provider_id` BIGINT(20) UNSIGNED NOT NULL,
    FOREIGN KEY (provider_id) REFERENCES provider_tab (id),
    `rates_id`    BIGINT(20) UNSIGNED NOT NULL,
    FOREIGN KEY (rates_id) REFERENCES rates_tab (id),
    `create_time` INT(10) UNSIGNED    NOT NULL,
    `update_time` INT(10) UNSIGNED    NOT NULL,
    `status`      INT(10) UNSIGNED    NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `evc`.`booking_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`booking_tab`
(
    `id`          BIGINT(20) UNSIGNED NOT NULL auto_increment,
    `charger_id`  BIGINT(20) UNSIGNED NOT NULL,
    FOREIGN KEY (charger_id) REFERENCES charger_tab (id),
    `email`       VARCHAR(256)        NOT NULL UNIQUE,
    `start_time`  INT(10) UNSIGNED    NOT NULL,
    `end_time`    INT(10) UNSIGNED    NOT NULL,
    `create_time` INT(10) UNSIGNED    NOT NULL,
    `update_time` INT(10) UNSIGNED    NOT NULL,
    `status`      INT(10) UNSIGNED    NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);