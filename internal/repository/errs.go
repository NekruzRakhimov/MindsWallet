package repository

import (
	"MindsWallet/internal/errs"
	"database/sql"
	"errors"
)

func translateError(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		return errs.ErrNotFound
	} else {
		return err
	}
}
