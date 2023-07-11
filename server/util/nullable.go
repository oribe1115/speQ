package util

import (
	"database/sql"
	"time"
)

func NullString2Pointer(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}

	return nil
}

func Pointer2NullString(ps *string) sql.NullString {
	if ps != nil {
		return sql.NullString{
			String: *ps,
			Valid:  true,
		}
	}

	return sql.NullString{}
}

func NullTime2Pointer(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}

	return nil
}

func Pointer2NullTime(pt *time.Time) sql.NullTime {
	if pt != nil {
		return sql.NullTime{
			Time:  *pt,
			Valid: true,
		}
	}

	return sql.NullTime{}
}
