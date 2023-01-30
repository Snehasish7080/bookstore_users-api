package mysql_utils

import (
	"github.com/Snehasish7080/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return errors.NewInternalServerError("error in parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")

}
