package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// StringArray is a helper for JSON array stored in JSONB column.
type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *StringArray) Scan(src any) error {
	var b []byte
	switch v := src.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	case nil:
		*s = StringArray{}
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
	return json.Unmarshal(b, s)
}

// QuizQuestion represents a quiz question in the pool.
type QuizQuestion struct {
	ID             pgtype.UUID        `db:"id"`
	Body           string             `db:"body"`
	CorrectAnswers StringArray        `db:"correct_answers"`
	Published      bool               `db:"published"`
	CreatedAt      pgtype.Timestamptz `db:"created_at"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at"`
}
