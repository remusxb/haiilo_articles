package postgres

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/lib/pq"
)

const (
	UniqueKeyCode = "23505"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
)

type ErrUniqueFieldViolation struct {
	Field string `json:"field"`
}

func (err *ErrUniqueFieldViolation) Error() string {
	return "unique key violation for field " + err.Field
}

// ParseError will look after few constraints and return not found error or a generic error
func ParseError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrResourceNotFound
	}

	dbErr, ok := err.(*pq.Error)
	if !ok {
		return err
	}

	switch dbErr.Code {
	case UniqueKeyCode:
		fieldName := strings.SplitN(strings.SplitN(dbErr.Detail, ")=(", 2)[0], "(", 2)[1]
		return &ErrUniqueFieldViolation{Field: fieldName}
	default:
		return err
	}
}
