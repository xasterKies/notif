# notif
===

> A simple and pluggable system notification package for Go.

![most language](https://img.shields.io/github/languages/top/xasterKies/notif?color=blue&style=for-the-badge)
![stars](https://img.shields.io/github/stars/xasterKies/notif?color=blue&style=for-the-badge)
![Contributors](https://img.shields.io/github/contributors/xasterKies/notif?color=blue&style=for-the-badge)

![Alt text](./notif-test-linux.png "notif")

### Why?
===

I started building this because of a personal project I am currently working on. I just needed a simple way to display system notifications in the operating system and all the other open source projects I stumbled on were a little too cumbersome for my needs. For a bit I started including this as one of the modules of the project but later on decided to make it a package so I could easily reuse.


## Features
===

- Easy display of system notifications compatible with Linux, Mac & Windows, all available and reusable withen the flow of your program.
- Easy to understand, making the package API easily customizable and extensible to fit your own need.
- End-to-end tested package and can be used in a production enviroment.


## API
===
A sample notification will need to fufill a few attributes: a `title`, a `message` and `severity` where severity provides how severe or urgent your notification is to be interpreted by the operating system. An example of a notification schema looks like this:

    ```go
        // Notification schema
        type Notify struct {
            title string
            message string
            severity Severity
        }
    ```

## TODO
===
 - [ ] Dockerize the package
 - [ ] Distribute the package on the go package manager `go get`


### Dependencies and Installs
===

- Golang v1.20
  - Install go for [linux/mac/windows](https://go.dev/doc/install)

## Setting up the Project
===

1. Fork repository

2. Clone the forked repository in prefered directory

   ```bash
   git clone <project-url>
   ```

3. Enter project directory
  
   ```bash
   cd notif
   ```

4. Initialize project

   ```bash
   git init
   ```

### Running the project
===

1. Install library dependencies

    ```bash
    go mod tidy
    ```

2. Build project

   ```bash
   go build
   ```

3. Run

    ```bash
   go run main.go
   ```


### Contributing
===
if you will like to make a contribution or suggest anything to the project,will be happy to hear from you in issues or PRs section :)    

