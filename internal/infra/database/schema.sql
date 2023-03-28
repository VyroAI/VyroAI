CREATE TABLE users
(
    id              bigint(19) UNSIGNED                    NOT NULL AUTO_INCREMENT,
    username        varchar(64)                            NOT NULL UNIQUE,
    email           varchar(100)                           NOT NULL UNIQUE,
    password        varchar(100)                           NOT NULL,
    avatar_id       varchar(64)                            NOT NULL DEFAULT "3eb0f8fa-a594-4348-9d89-9ab7c9be4842",
    permission      TINYINT                                NOT NULL DEFAULT 1,
    email_confirmed TINYINT(1)                             NOT NULL DEFAULT 0,
    status          ENUM ('ACTIVE', 'SUSPENDED', 'BANNED') NOT NULL DEFAULT 'ACTIVE',
    created_at      TIMESTAMP                              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);


CREATE TABLE newsletter_subscribed
(
    id    bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    email varchar(100)        NOT NULL UNIQUE,
    PRIMARY KEY (id)

);

CREATE TABLE subscription_plans
(
    id          bigint(19) UNSIGNED              NOT NULL AUTO_INCREMENT,
    name        ENUM ("FREE","PRO","ENTERPRISE") NOT NULL DEFAULT "FREE",
    description TEXT,
    price       DECIMAL(10, 2)                   NOT NULL,
    duration    INT                              NOT NULL,
    PRIMARY KEY (id)

);


CREATE TABLE user_subscriptions
(
    id         bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id    bigint(19) UNSIGNED NOT NULL,
    plan_id    INT UNSIGNED        NOT NULL,
    start_date DATE                NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);


CREATE TABLE payments
(
    id                bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id           bigint(19) UNSIGNED NOT NULL,
    plan_id           INT UNSIGNED        NOT NULL,
    amount            DECIMAL(10, 2)      NOT NULL,
    stripe_payment_id VARCHAR(100)        NOT NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);