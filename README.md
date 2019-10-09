# XGONET

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

The application contain a diverse amount of functionalities which can be performed by passing the correct parameters to it. Down here we describe the commands allowed by the current version:

1. `--help`. This parementer display all the commands that can be performed.
2. `--data=<test_data.json>`. This is the path and the name of the file that will be loaded to be used as mock data.
3. `--func=<sample>`. These are that functions that can be performed.
4. `--param=<params>`. This is the data that will be sent to filer.
5. `--depth=<number>`. The depeth of the search and will depend on the functionalty requested. The default value is 10 and contain the max depth of search made by to code to teremine if two users are related
6. `--debug=<boolean>`. Default value is **false**, you must set this equal to **true** so the program can display time and logs in the execution.

### Project Requirements

1. The program will use mock data at startup.
2. The program will use a flag "username1" and "username2" to determine if the 2 usernames are connected via their friend list.
3. The program will output how much time it took to determine if the 2 usernames are connected and how many users it had to traverse to find the answer.
4. Include any kind of automated tests.