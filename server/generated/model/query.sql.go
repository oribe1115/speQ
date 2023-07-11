// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: query.sql

package model

import (
	"context"
	"database/sql"
)

const countContestInfoRow = `-- name: CountContestInfoRow :one
SELECT COUNT(*)
FROM ` + "`" + `contest_info` + "`" + `
`

func (q *Queries) CountContestInfoRow(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countContestInfoRow)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countRowAsAdmin = `-- name: CountRowAsAdmin :one
SELECT COUNT(*)
FROM ` + "`" + `admins` + "`" + `
WHERE ` + "`" + `trap_id` + "`" + ` = ?
`

func (q *Queries) CountRowAsAdmin(ctx context.Context, trapID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRowAsAdmin, trapID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countRowAsRoot = `-- name: CountRowAsRoot :one
SELECT COUNT(*)
FROM ` + "`" + `roots` + "`" + `
WHERE ` + "`" + `trap_id` + "`" + ` = ?
`

func (q *Queries) CountRowAsRoot(ctx context.Context, trapID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRowAsRoot, trapID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteAllAdminUsers = `-- name: DeleteAllAdminUsers :exec
TRUNCATE ` + "`" + `admins` + "`" + `
`

func (q *Queries) DeleteAllAdminUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllAdminUsers)
	return err
}

const deleteAllContestInfo = `-- name: DeleteAllContestInfo :exec
TRUNCATE ` + "`" + `contest_info` + "`" + `
`

func (q *Queries) DeleteAllContestInfo(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllContestInfo)
	return err
}

const deleteAllContestants = `-- name: DeleteAllContestants :exec
DELETE
FROM ` + "`" + `contestants` + "`" + `
WHERE TRUE
`

func (q *Queries) DeleteAllContestants(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllContestants)
	return err
}

const deleteAllRootUsers = `-- name: DeleteAllRootUsers :exec
TRUNCATE ` + "`" + `roots` + "`" + `
`

func (q *Queries) DeleteAllRootUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllRootUsers)
	return err
}

const getAdminUsers = `-- name: GetAdminUsers :many
SELECT trap_id
FROM ` + "`" + `admins` + "`" + `
`

func (q *Queries) GetAdminUsers(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getAdminUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var trap_id string
		if err := rows.Scan(&trap_id); err != nil {
			return nil, err
		}
		items = append(items, trap_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getContestInfo = `-- name: GetContestInfo :one
SELECT title, description, scheduled_start_time, start_time, end_time, voting_freeze_time
FROM ` + "`" + `contest_info` + "`" + `
LIMIT 1
`

func (q *Queries) GetContestInfo(ctx context.Context) (ContestInfo, error) {
	row := q.db.QueryRowContext(ctx, getContestInfo)
	var i ContestInfo
	err := row.Scan(
		&i.Title,
		&i.Description,
		&i.ScheduledStartTime,
		&i.StartTime,
		&i.EndTime,
		&i.VotingFreezeTime,
	)
	return i, err
}

const getContestants = `-- name: GetContestants :many
SELECT trap_id
FROM ` + "`" + `contestants` + "`" + `
`

func (q *Queries) GetContestants(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getContestants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var trap_id string
		if err := rows.Scan(&trap_id); err != nil {
			return nil, err
		}
		items = append(items, trap_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLastVote = `-- name: GetLastVote :one
SELECT id, voter, target, created_at
FROM ` + "`" + `votes` + "`" + `
WHERE ` + "`" + `voter` + "`" + ` = ?
ORDER BY ` + "`" + `created_at` + "`" + ` DESC
LIMIT 1
`

func (q *Queries) GetLastVote(ctx context.Context, voter string) (Vote, error) {
	row := q.db.QueryRowContext(ctx, getLastVote, voter)
	var i Vote
	err := row.Scan(
		&i.ID,
		&i.Voter,
		&i.Target,
		&i.CreatedAt,
	)
	return i, err
}

const getLatestVotes = `-- name: GetLatestVotes :many
SELECT id, voter, target, created_at
FROM ` + "`" + `votes` + "`" + ` AS ` + "`" + `main` + "`" + `
WHERE ` + "`" + `main` + "`" + `.` + "`" + `id` + "`" + ` = (SELECT ` + "`" + `sub` + "`" + `.` + "`" + `id` + "`" + `
                     FROM ` + "`" + `votes` + "`" + ` AS ` + "`" + `sub` + "`" + `
                     WHERE ` + "`" + `main` + "`" + `.` + "`" + `voter` + "`" + ` = ` + "`" + `sub` + "`" + `.` + "`" + `voter` + "`" + `
                     ORDER BY ` + "`" + `sub` + "`" + `.` + "`" + `created_at` + "`" + ` DESC
                     LIMIT 1)
`

func (q *Queries) GetLatestVotes(ctx context.Context) ([]Vote, error) {
	rows, err := q.db.QueryContext(ctx, getLatestVotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Vote
	for rows.Next() {
		var i Vote
		if err := rows.Scan(
			&i.ID,
			&i.Voter,
			&i.Target,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRootUsers = `-- name: GetRootUsers :many
SELECT trap_id
FROM ` + "`" + `roots` + "`" + `
`

func (q *Queries) GetRootUsers(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getRootUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var trap_id string
		if err := rows.Scan(&trap_id); err != nil {
			return nil, err
		}
		items = append(items, trap_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAdminUser = `-- name: InsertAdminUser :exec
INSERT INTO ` + "`" + `admins` + "`" + ` (` + "`" + `trap_id` + "`" + `)
VALUES (?)
`

func (q *Queries) InsertAdminUser(ctx context.Context, trapID string) error {
	_, err := q.db.ExecContext(ctx, insertAdminUser, trapID)
	return err
}

const insertContestInfo = `-- name: InsertContestInfo :exec
INSERT INTO ` + "`" + `contest_info` + "`" + `
(` + "`" + `title` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `scheduled_start_time` + "`" + `, ` + "`" + `start_time` + "`" + `, ` + "`" + `end_time` + "`" + `, ` + "`" + `voting_freeze_time` + "`" + `)
VALUES (?, ?, ?, ?, ?, ?)
`

type InsertContestInfoParams struct {
	Title              string
	Description        sql.NullString
	ScheduledStartTime sql.NullTime
	StartTime          sql.NullTime
	EndTime            sql.NullTime
	VotingFreezeTime   sql.NullTime
}

func (q *Queries) InsertContestInfo(ctx context.Context, arg InsertContestInfoParams) error {
	_, err := q.db.ExecContext(ctx, insertContestInfo,
		arg.Title,
		arg.Description,
		arg.ScheduledStartTime,
		arg.StartTime,
		arg.EndTime,
		arg.VotingFreezeTime,
	)
	return err
}

const insertContestant = `-- name: InsertContestant :exec
# TRUNCATE is forbidden since it's referenced with foreign keys

INSERT INTO ` + "`" + `contestants` + "`" + `(` + "`" + `trap_id` + "`" + `)
VALUES (?)
`

func (q *Queries) InsertContestant(ctx context.Context, trapID string) error {
	_, err := q.db.ExecContext(ctx, insertContestant, trapID)
	return err
}

const insertRootUser = `-- name: InsertRootUser :exec
INSERT INTO ` + "`" + `roots` + "`" + ` (` + "`" + `trap_id` + "`" + `)
VALUES (?)
`

func (q *Queries) InsertRootUser(ctx context.Context, trapID string) error {
	_, err := q.db.ExecContext(ctx, insertRootUser, trapID)
	return err
}

const insertScore = `-- name: InsertScore :exec
INSERT INTO ` + "`" + `scores` + "`" + ` (` + "`" + `contestant_id` + "`" + `, ` + "`" + `score` + "`" + `)
VALUES (?, ?)
`

type InsertScoreParams struct {
	ContestantID string
	Score        float64
}

func (q *Queries) InsertScore(ctx context.Context, arg InsertScoreParams) error {
	_, err := q.db.ExecContext(ctx, insertScore, arg.ContestantID, arg.Score)
	return err
}

const insertVote = `-- name: InsertVote :exec
INSERT INTO ` + "`" + `votes` + "`" + ` (` + "`" + `voter` + "`" + `, ` + "`" + `target` + "`" + `)
VALUES (?, ?)
`

type InsertVoteParams struct {
	Voter  string
	Target string
}

func (q *Queries) InsertVote(ctx context.Context, arg InsertVoteParams) error {
	_, err := q.db.ExecContext(ctx, insertVote, arg.Voter, arg.Target)
	return err
}
