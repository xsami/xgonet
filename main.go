package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	flags "github.com/jessevdk/go-flags"
	model "github.com/xsami/xgonet/models"
)

// UFStruct is the structure to parse the mock file with friend and user model
type UFStruct struct {
	Users   []model.User   `json:"users"`
	Friends []model.Friend `json:"friends"`
}

var (
	opts struct {
		Debug    bool              `short:"d" long:"debug" description:"Show debug information"`
		Data     string            `short:"D" long:"data" description:"Data path (absolute path)"`
		Func     string            `short:"f" long:"function" description:"Call the specified function"`
		Param    map[string]string `short:"p" long:"param" description:"parameters which is a map from string to string"`
		Treshold uint              `short:"t" long:"treshold" description:"treshold of max amount of iterations"`
	}
	data UFStruct
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Parse the args passed on the program execution
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	// Parse the mock data
	if err := LoadModel(opts.Data, &data); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	if opts.Debug {
		// fmt.Printf("%+v\n\n%+v\n", data, opts)
	}

	log.Fatal(execFunction()) // Execute the called function
}
