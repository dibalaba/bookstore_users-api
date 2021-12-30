package services

import (
	"github.com/dibalaba/bookstore_users-api/domain/users"
	"github.com/dibalaba/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//return nil, nil
	return &user, nil
	/* return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	} */
}
