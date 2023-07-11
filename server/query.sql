-- name: InsertRootUser :exec
INSERT INTO `roots` (`trap_id`)
VALUES (?);

-- name: DeleteAllRootUsers :exec
TRUNCATE `roots`;

-- name: CountRowAsRoot :one
SELECT COUNT(*)
FROM `roots`
WHERE `trap_id` = ?;

-- name: GetRootUsers :many
SELECT *
FROM `roots`;

-- name: InsertAdminUser :exec
INSERT INTO `admins` (`trap_id`)
VALUES (?);

-- name: DeleteAllAdminUsers :exec
TRUNCATE `admins`;

-- name: CountRowAsAdmin :one
SELECT COUNT(*)
FROM `admins`
WHERE `trap_id` = ?;

-- name: GetAdminUsers :many
SELECT *
FROM `admins`;

-- name: CountContestInfoRow :one
SELECT COUNT(*)
FROM `contest_info`;

-- name: GetContestInfo :one
SELECT *
FROM `contest_info`
LIMIT 1;

-- name: InsertContestInfo :exec
INSERT INTO `contest_info`
(`title`, `description`, `scheduled_start_time`, `start_time`, `end_time`, `voting_freeze_time`)
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteAllContestInfo :exec
TRUNCATE `contest_info`;

-- name: GetContestants :many
SELECT *
FROM `contestants`;

-- name: DeleteAllContestants :exec
DELETE
FROM `contestants`
WHERE TRUE;
# TRUNCATE is forbidden since it's referenced with foreign keys

-- name: InsertContestant :exec
INSERT INTO `contestants`(`trap_id`)
VALUES (?);

-- name: GetLastVote :one
SELECT *
FROM `votes`
WHERE `voter` = ?
ORDER BY `created_at` DESC
LIMIT 1;

-- name: InsertVote :exec
INSERT INTO `votes` (`voter`, `target`)
VALUES (?, ?);

-- name: GetLatestVotes :many
SELECT *
FROM `votes` AS `main`
WHERE `main`.`id` = (SELECT `sub`.`id`
                     FROM `votes` AS `sub`
                     WHERE `main`.`voter` = `sub`.`voter`
                     ORDER BY `sub`.`created_at` DESC
                     LIMIT 1);

-- name: InsertScore :exec
INSERT INTO `scores` (`contestant_id`, `score`)
VALUES (?, ?);

-- name: GetLatestScores :many
SELECT *
FROM `scores` AS `main`
WHERE `main`.`id` = (SELECT `sub`.`id`
                     FROM `scores` AS `sub`
                     WHERE `main`.`id` = `sub`.`id`
                     ORDER BY `sub`.`created_at` DESC
                     LIMIT 1);