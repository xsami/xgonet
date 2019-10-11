package main

import (
	"fmt"
	"strings"

	model "github.com/xsami/xgonet/models"
)

const (
	// FindTwoUserRelationShip constant function name
	FindTwoUserRelationShip = "findtwouserrelationship"
)

// execFunction try to execute the functions that
// can be executed by the users
func execFunction() error {
	var (
		functionName string
		parameters   map[string]string
	)

	functionName = strings.ToLower(opts.Func)
	parameters = opts.Param

	switch functionName {
	case FindTwoUserRelationShip:

		userA, err := model.FindUserByUsername(data.Users, parameters["username1"])
		if err != nil {
			return err
		}

		userB, err := model.FindUserByUsername(data.Users, parameters["username2"])
		if err != nil {
			return err
		}

		resultUser, _ := model.FindTwoUserRelationShip(data.Friends, userA, userB, 0, opts.Treshold)

		fmt.Println(resultUser)
	default:
		return fmt.Errorf("Function: %v wasn't found", opts.Func)
	}
	return nil
}
