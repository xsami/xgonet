package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	flags "github.com/jessevdk/go-flags"
	loader "github.com/xsami/xgonet/loader"
	logger "github.com/xsami/xgonet/logger"
	model "github.com/xsami/xgonet/models"
)

var (
	opts struct {
		Debug    bool              `short:"d" long:"debug" description:"Show debug information"`
		Data     string            `short:"D" long:"data" description:"Data path (absolute path)" default:"mock_data\\test_data1.json"`
		Func     string            `short:"f" long:"function" description:"Call the specified function" default:"findtwouserrelationship"`
		Param    map[string]string `short:"p" long:"param" description:"parameters which is a map from string to string"`
		Treshold uint              `short:"t" long:"treshold" description:"treshold of max amount of iterations"`
	}
	data loader.UFStruct
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Parse the args passed on the program execution
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		log.Fatal(err.Error() + "\n")
	}

	// Parse the mock data
	if err := loader.LoadModel(opts.Data, &data); err != nil {
		log.Fatal(err.Error() + "\n")
	}

	// Inject FriendList, UserList and global variables for model & logger package
	model.FriendList = data.Friends
	model.UserList = data.Users
	model.FriendMap = model.BuildFriendMap(data.Friends)
	logger.Debug = opts.Debug

	logger.Log("main", data, opts)

	log.Fatal(execFunction()) // Execute the called function
}
