# XGONET

[![Go Report Card](https://goreportcard.com/badge/xsami/xgonet)](https://goreportcard.com/report/xsami/xgonet)
[![GoDoc](https://godoc.org/xsami/xgonet?status.svg)](https://godoc.org/xsami/xgonet)
[![CircleCI](https://circleci.com/gh/xsami/xgonet.svg?style=svg)](https://circleci.com/gh/xsami/xgonet)
[![Build Status](https://travis-ci.com/xsami/xgonet.svg?branch=master)](https://travis-ci.com/xsami/xgonet)


A Social Network built with `go`.

### Configuration

Current version used for the development of the application is go1.11.4 but it would probably be compatible with any version >= go1.11.X 

### Usage

To run the project you can clone the repository using the following command:
```
$ git clone https://github.com/xsami/xgonet
```
Or by the dependency manager with the comand:
```
$ go get github.com/xsami/xgonet
```
which will download the project inside your directory `$GOPATH/src/github.com/xsami/xgonet`


To run the project just enter to the project main directory and run the command: `go run main.go` and for building run the command: `go build ${project_path}`.

#### Commands

The application contain a diverse amount of functionalities which can be performed by passing the correct parameters to it. Down here we describe the commands allowed by the current version

1. `--help`. This parameter display all the commands that can be performed.
2. `--data=<test_data.json>`. This is the path and the name of the file that will be loaded to be used as mock data.
3. `--func=<sample>`. These are functions that can be performed by the application. More details about this in the **Functions** section
4. `--param=<params>`. These are the parameters passed to the function that will be called. More details about this in the **Functions** section 
5. `--treshold=<number>`. The deepest amount of the search. The default value is 0 (which mean infinite) and contain the max depth of search made by to code to determinate if two users are related.
6. `--debug=<boolean>`. Default value is **false**, you must set this equal to **true** so the program can display time and logs in the execution.

##### Example on Mac OS X
```
$ go build .
$ ./xgonet -h
$ ./xgonet -d true -D mock_data/test_data1.json -f findtwouserrelationship -p username1:mhallihan0 -p username2:mbrereton1
```

##### Example on Windows
Note that for windows build the argument prefix is `/` instead of `-`
```
$ go build .
$ xgonet.exe /h
$ xgonet.exe /d true /D mock_data\test_data1.json /f findtwouserrelationship /p username1:mhallihan0 /p username2:mbrereton1
```

### Functions

This section contain the functions that can be called in this application, which is the main purpose of this application. Find in the table below which functionality you want to use: 


 Function   |      Parameter      |  Type |  Description 
------------|---------------------|-------|---------------
| FindTwoUserRelationShip |  username1, username2 | string, string | Find if 2 users can have related friends and display this relationship. **Note**: the treshold on this function means the following: if `0` there is no limit on the recursion search. `1` means that only look if they are direct friends. `2` means that look if at least have 1 friend related. `3` means that they have a friend that have a friend in common, and so on |


### Project Requirements

1. The program will use mock data at startup.
2. The program will use a flag "username1" and "username2" to determine if the 2 usernames are connected via their friend list.
3. The program will output how much time it took to determine if the 2 usernames are connected and how many users it had to traverse to find the answer.
4. Include any kind of automated tests.