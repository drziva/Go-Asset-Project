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

		switch pgErr.Code {

		case "23505": // unique_violation
			// optionally inspect constraint name
			if pgErr.ConstraintName == "idx_users_email" {
				return domainErrors.ErrEmailAlreadyExists
			}
			return domainErrors.ErrConflict

		case "23503": // foreign_key_violation
			return domainErrors.ErrInvalidReference

		case "23502": // not_null_violation
			return domainErrors.ErrMissingRequiredField

		case "23514": // check_violation
			return domainErrors.ErrInvalidInput

		case "22P02": // invalid_text_representation
			return domainErrors.ErrInvalidFormat

		case "40P01": // deadlock_detected
			return domainErrors.ErrConflict // or retry logic if you want to be fancy
		}
	}

	return err
}
