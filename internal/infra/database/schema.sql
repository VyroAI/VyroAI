CREATE TABLE users
(
    id              bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    username        varchar(64)         NOT NULL UNIQUE,
    email           varchar(100)        NOT NULL UNIQUE,
    password        varchar(100),
    avatar_id       varchar(64)         NOT NULL DEFAULT "3eb0f8fa-a594-4348-9d89-9ab7c9be4842",
    permission      TINYINT             NOT NULL DEFAULT 0,
    subscription_id bigint(19) UNSIGNED NOT NULL,
    fingerprint_id  varchar(64),
    oauth_id        bigint(19) UNSIGNED,
    email_confirmed TINYINT(1)          NOT NULL DEFAULT 0,
    is_banned       TINYINT(1)          NOT NULL DEFAULT 0,
    created_at      TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE oauth_register
(
    id             bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    oauth_provider ENUM ("APPLE","GOOGLE","INSTAGRAM","DISCORD"),

    PRIMARY KEY (id)
);



CREATE TABLE user_subscriptions
(
    id             bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    plan_id        TINYINT UNSIGNED    NOT NULL DEFAULT 1,
    openai_api_key varchar(160),
    api_key        varchar(160)        NOT NULL UNIQUE,
    created_at     TIMESTAMP                    DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP                    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id, api_key)
);

CREATE TABLE subscription_plans
(
    id                   TINYINT UNSIGNED    NOT NULL AUTO_INCREMENT,
    name                 varchar(20)         NOT NULL,
    description          TEXT                NOT NULL,
    price                DECIMAL(10, 2)      NOT NULL,
    message_limit        INT UNSIGNED        NOT NULL,
    character_limit      bigint(19) UNSIGNED NOT NULL,
    chatbot              INT UNSIGNED        NOT NULL,
    multiple_file_upload TINYINT(1)          NOT NULL DEFAULT 0,
    show_branding        TINYINT(1)          NOT NULL DEFAULT 1,
    api_access           TINYINT(1)          NOT NULL DEFAULT 0,

    PRIMARY KEY (id)
);



CREATE TABLE chat_bot
(
    id              bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    chatbot_id      bigint(19) UNSIGNED NOT NULL,
    user_id         bigint(19) UNSIGNED NOT NULL,
    title           varchar(100)        NOT NULL,
    character_count INT UNSIGNED        NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id, chatbot_id)
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


CREATE TABLE newsletter_subscribed
(
    id    bigint(19) UNSIGNED NOT NULL AUTO_INCREMENT,
    email varchar(100)        NOT NULL UNIQUE,

    PRIMARY KEY (id)
);