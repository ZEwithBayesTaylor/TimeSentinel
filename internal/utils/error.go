package utils

import (
	"errors"
	mysql2 "github.com/go-sql-driver/mysql"
)

const DuplicateEntryErrCode = 1062

func IsDuplicateEntryErr(err error) bool {
	var mysqlErr *mysql2.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == DuplicateEntryErrCode
}
