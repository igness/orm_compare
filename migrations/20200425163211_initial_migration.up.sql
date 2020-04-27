CREATE TABLE `users` (
    `id`            BIGINT UNSIGNED     NOT NULL    AUTO_INCREMENT,
    `created_at`    DATETIME            NOT NULL,
    `updated_at`    DATETIME            NOT NULL,
    `name`          VARCHAR(255)        NOT NULL,
    `gender`        VARCHAR(255)            NULL,
    `email`         VARCHAR(255)            NULL,
    PRIMARY KEY(`id`)
);

CREATE TABLE `foods` (
    `id`            BIGINT UNSIGNED     NOT NULL    AUTO_INCREMENT,
    `created_at`    DATETIME            NOT NULL,
    `updated_at`    DATETIME            NOT NULL,
    `name`          VARCHAR(255)        NOT NULL,
    `price`         NUMERIC(27,9)       NOT NULL,
    PRIMARY KEY(`id`)
);

CREATE TABLE `orders` (
    `id`                BIGINT UNSIGNED     NOT NULL    AUTO_INCREMENT,
    `order_id`          CHAR(36)            NOT NULL,
    `created_at`        DATETIME            NOT NULL,
    `updated_at`        DATETIME            NOT NULL,
    `user_id`           BIGINT UNSIGNED     NOT NULL,
    `food_id`           BIGINT UNSIGNED     NOT NULL,
    `purchase_count`    INT                 NOT NULL,
    `purchase_amount`   NUMERIC(27,9)       NOT NULL,
    PRIMARY KEY(`id`)
);