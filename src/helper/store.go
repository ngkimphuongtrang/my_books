package helper

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

const mysqlDuplicateKeyErrorCode = 1062

func IsDuplicateKeyError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == mysqlDuplicateKeyErrorCode
	}
	return false
}
