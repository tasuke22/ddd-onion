CREATE TABLE users
(
    id         VARCHAR(255) PRIMARY KEY NOT NULL,
    email      VARCHAR(255)             NOT NULL UNIQUE,
    password   VARCHAR(255)             NOT NULL,
    name       VARCHAR(255)             NOT NULL,
    profile    VARCHAR(255)             NOT NULL,
    created_at DATETIME                 NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME                 NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE tags
(
    id         VARCHAR(255) PRIMARY KEY NOT NULL,
    name       VARCHAR(255)             NOT NULL UNIQUE,
    created_at DATETIME                 NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME                 NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE skills
(
    id         VARCHAR(255) PRIMARY KEY NOT NULL,
    user_id    VARCHAR(255)             NOT NULL,
    tag_id     VARCHAR(255)             NOT NULL,
    evaluation INT                      NOT NULL,
    years      INT                      NOT NULL,
    created_at DATETIME                 NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME                 NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_skills_user_id FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT fk_skills_tag_id FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;


CREATE TABLE careers
(
    id         VARCHAR(255) PRIMARY KEY,
    user_id    VARCHAR(255) NOT NULL,
    detail     VARCHAR(255) NOT NULL,
    start_year INT          NOT NULL,
    end_year   INT          NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_careers_user_id FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

