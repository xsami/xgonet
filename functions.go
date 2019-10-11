package main

import (
	"fmt"
	"strings"
	"time"

	model "github.com/xsami/xgonet/models"
)

const (
	// FindTwoUserRelationShip constant function name
	FindTwoUserRelationShip = "findtwouserrelationship"
)

func fmtDuration(d time.Duration) {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	fmt.Printf("%02d:%02d", h, m)
}

// execFunction try to execute the functions that
// can be executed by the users
func execFunction() error {

	modTime := time.Now().Round(0).Add(-(3600 + 60 + 45) * time.Second)
	since := time.Since(modTime)
	fmt.Println("Start: ", since)
	defer fmtDuration(since)

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
		return fmt.Errorf("Function: %v wasn't found", functionName)
	}

	return nil
}
