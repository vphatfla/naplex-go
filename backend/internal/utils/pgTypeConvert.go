package utils

import "github.com/jackc/pgx/v5/pgtype"

func PgTextToString(t pgtype.Text) string {
	if t.Valid {
		return t.String
	}
	return ""
}
