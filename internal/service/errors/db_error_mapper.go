package errors

import (
	"errors"

	domainErrors "go-project/internal/errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// had to implement both gorm and pg error checking here since gorm errors seemed to propagate inconsistently
func MapDBError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domainErrors.ErrNotFound
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return domainErrors.ErrEmailAlreadyExists
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// code 23505 is unique_violation in PostgreSQL -> email already exists
		if pgErr.Code == "23505" {
			return domainErrors.ErrEmailAlreadyExists
		}
	}

	return err
}
