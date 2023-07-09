CREATE TABLE `roots`
(
    `trap_id` VARCHAR(32) PRIMARY KEY
);

CREATE TABLE `admins`
(
    `trap_id` VARCHAR(32) PRIMARY KEY
);

# This table must have only one record.
CREATE TABLE `contest_info`
(
    `title`                VARCHAR(30) PRIMARY KEY,
    `description`          TEXT,
    `scheduled_start_time` TIMESTAMP NULL DEFAULT NULL,
    `start_time`           TIMESTAMP NULL DEFAULT NULL,
    `end_time`             TIMESTAMP NULL DEFAULT NULL,
    `voting_freeze_time`   TIMESTAMP NULL DEFAULT NULL
);

CREATE TABLE `contestants`
(
    `trap_id` VARCHAR(32) PRIMARY KEY
);

CREATE TABLE `votes`
(
    `id`         INT AUTO_INCREMENT PRIMARY KEY,
    `voter`      VARCHAR(32) NOT NULL,
    `target`     VARCHAR(32) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY `vote_contestant` (`target`) REFERENCES `contestants` (`trap_id`) ON DELETE CASCADE
);

CREATE TABLE `problems`
(
    `id`          INT PRIMARY KEY, # Also used to specify the display order.
    `title`       VARCHAR(20) NOT NULL,
    `description` TEXT
);

CREATE TABLE `solved`
(
    `problem_id`    INT,
    `contestant_id` VARCHAR(32),
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`problem_id`, `contestant_id`),
    FOREIGN KEY `solved_problem` (`problem_id`) REFERENCES `problems` (`id`) ON DELETE CASCADE,
    FOREIGN KEY `solved_contestant` (`contestant_id`) REFERENCES `contestants` (`trap_id`) ON DELETE CASCADE
);

CREATE TABLE `scores`
(
    `id`            INT AUTO_INCREMENT PRIMARY KEY,
    `contestant_id` VARCHAR(32) NOT NULL,
    `score`         DOUBLE      NOT NULL,
    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY `score_contestant` (`contestant_id`) REFERENCES `contestants` (`trap_id`) ON DELETE CASCADE
);