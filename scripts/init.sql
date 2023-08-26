CREATE DATABASE IF NOT EXISTS `evc`;

DROP TABLE IF EXISTS `evc`.`rates_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`rates_tab`
(
    `id`              BIGINT(20) UNSIGNED   NOT NULL auto_increment,
    `normal_rate`     DOUBLE(3, 3) UNSIGNED NOT NULL,
    `penalty_rate`    DOUBLE(3, 3) UNSIGNED NOT NULL,
    `no_show_penalty` DOUBLE(3, 3) UNSIGNED NOT NULL,
    `created_at`      timestamp             NOT NULL,
    `updated_at`      timestamp             NOT NULL,
    `deleted_at`      timestamp,
    `status`          INT(10) UNSIGNED      NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `evc`.`provider_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`provider_tab`
(
    `id`           BIGINT(20) UNSIGNED NOT NULL auto_increment,
    `company_name` VARCHAR(256)        NOT NULL,
    `description`  VARCHAR(256)        NOT NULL,
    `created_at`   timestamp           NOT NULL,
    `updated_at`   timestamp           NOT NULL,
    `deleted_at`   timestamp,
    `status`       INT(10) UNSIGNED    NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `evc`.`charger_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`charger_tab`
(
    `id`          BIGINT(20) UNSIGNED NOT NULL auto_increment,
    `provider_id` BIGINT(20) UNSIGNED NOT NULL,
    `location`    VARCHAR(256)        NOT NULL,
    FOREIGN KEY (provider_id) REFERENCES provider_tab (id),
    `rates_id`    BIGINT(20) UNSIGNED NOT NULL,
    FOREIGN KEY (rates_id) REFERENCES rates_tab (id),
    `created_at`  timestamp           NOT NULL,
    `updated_at`  timestamp           NOT NULL,
    `deleted_at`  timestamp,
    `status`      INT(10) UNSIGNED    NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `evc`.`booking_tab`;
CREATE TABLE IF NOT EXISTS `evc`.`booking_tab`
(
    `id`         BIGINT(20) UNSIGNED NOT NULL auto_increment,
    `charger_id` BIGINT(20) UNSIGNED NOT NULL,
    FOREIGN KEY (charger_id) REFERENCES charger_tab (id),
    `email`      VARCHAR(256)        NOT NULL,
    `start_time` timestamp           NOT NULL,
    `end_time`   timestamp           NOT NULL,
    `created_at` timestamp           NOT NULL,
    `updated_at` timestamp           NOT NULL,
    `deleted_at` timestamp,
    `status`     VARCHAR(256)        NOT NULL,
    PRIMARY KEY (`id`)
);


INSERT INTO `evc`.`provider_tab` (`id`, `company_name`, `description`, `created_at`, `updated_at`, `deleted_at`, `status`) VALUES
(1,	'companyA',	'',	'2023-08-26 08:39:23',	'2023-08-26 08:39:23',	NULL,	0);

INSERT INTO `evc`.`rates_tab` (`id`, `normal_rate`, `penalty_rate`, `no_show_penalty`, `created_at`, `updated_at`, `deleted_at`, `status`) VALUES
(1,	0.300,	0.900,	0.900,	'2023-08-26 08:43:45',	'2023-08-26 08:43:45',	NULL,	0);

INSERT INTO `evc`.`charger_tab` (`id`, `provider_id`, `location`,`rates_id`, `created_at`, `updated_at`, `deleted_at`, `status`) VALUES
(1,	1, 'yishun',	1,	'2023-08-26 08:46:42',	'2023-08-26 08:46:42',	NULL,	0);