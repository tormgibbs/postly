package db

import (
	"database/sql"
	"time"
)

func ToNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}
