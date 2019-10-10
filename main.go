package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	flags "github.com/jessevdk/go-flags"
	model "github.com/xsami/xgonet/models"
)

type UFStruct struct {
	Users   []model.User   `json:"users"`
	Friends []model.Friend `json:"friends"`
}

var (
	opts struct {
		Debug    bool              `short:"d" long:"debug" description:"Show debug information"`
		Data     string            `short:"D" long:"data" description:"Data path (absolute path)"`
		Func     string            `short:"f" long:"function" description:"Call the specified function"`
		Param    map[string]string `short:"p" long:"param" description:"paramters which is a map from string to string"`
		Treshold uint              `short:"t" long:"treshold" description:"treshold of max amount of iterations"`
	}
	data UFStruct
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	if err := LoadModel(opts.Data, &data); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	fmt.Printf("%+v\n\n%+v", opts, data)
}

// GetOptions return the option with the paremeter passed
func GetOptions() interface{} {
	return opts
}

// GetData return the data value
func GetData() UFStruct {
	return data
}
