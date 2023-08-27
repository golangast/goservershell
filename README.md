# goservershell

<p align="left"> <img src="https://komarev.com/ghpvc/?username=golangast&label=Profile%20views&color=0e75b6&style=flat" alt="golangast" /> </p>


![GitHub repo file count](https://img.shields.io/github/directory-file-count/golangast/goservershell) 
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/golangast/goservershell)
![GitHub repo size](https://img.shields.io/github/repo-size/golangast/goservershell)
![GitHub](https://img.shields.io/github/license/golangast/goservershell)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/golangast/goservershell)
![Go 100%](https://img.shields.io/badge/Go-100%25-blue)
![status beta](https://img.shields.io/badge/Status-Beta-red)

<h3 align="left">Languages and Tools:</h3>
<p align="left"> <a href="https://getbootstrap.com" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/bootstrap/bootstrap-plain-wordmark.svg" alt="bootstrap" width="40" height="40"/> </a> <a href="https://www.w3schools.com/css/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/css3/css3-original-wordmark.svg" alt="css3" width="40" height="40"/> </a> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.w3.org/html/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/html5/html5-original-wordmark.svg" alt="html5" width="40" height="40"/> </a> <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/javascript/javascript-original.svg" alt="javascript" width="40" height="40"/> </a> <a href="https://www.mysql.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mysql/mysql-original-wordmark.svg" alt="mysql" width="40" height="40"/> </a> </p>

## goservershell
- [goservershell](#goservershell)
  - [goservershell](#goservershell-1)
  - [General info](#general-info)
  - [Why build this?](#why-build-this)
  - [What does it do?](#what-does-it-do)
  - [Technologies](#technologies)
  - [Requirements](#requirements)
  - [Setup](#setup)
  - [Commands](#commands)
  - [Repository overview](#repository-overview)
  - [Special thanks](#special-thanks)



## General info
This project is a template for gonew and is used for setting up a webserver using echo framework.


## Why build this?
* Go never changes
* It is a nice way to start out a webserver without doing much


## What does it do?
* It is in pure Go so faster build times and since Go never changes it will always compile.
* No need for npm with assets because it concatenates and optimizes them (with min command)
* Provides code for bare basic security
* Allows for sending emails
* Database setup (sqlite)
* Has middleware support
* Is scaffolding for apis, crud, security


## Technologies
Project is created with:
* [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) - database
* [go-ps](https://github.com/mitchellh/go-ps) - getting pids in all OS's
* [viper](https://github.com/spf13/cobra) - build cli commands
* [echo](https://github.com/labstack/echo/v4) - web framework to shorten code
* [sprig](https://github.com/Masterminds/sprig) - template functions
* [imagecompression](https://github.com/nurlantulemisov/imagecompression) - image compression
* [minify](https://github.com/tdewolff/minify) - assets optimization
* [gomail](https://gopkg.in/gomail.v2) - email accessibility
* [jwt](https://github.com/golang-jwt/jwt) - JWT authentication
* [validator](https://github.com/go-playground/validator) - Validation


## Requirements
* go 1.21 for gonew

## Setup
Just use the new [gonew](https://go.dev/blog/gonew)

```
go install golang.org/x/tools/cmd/gonew@latest

gonew github.com/golangast/goservershell example.com/myserver

go mod tidy

go mod vendor


```
REMEMBER TO RUN 'go mod tidy' and 'go mod vendor' after to pull down dependencies

## Commands
```
go run . st //to run the program

go run . min //to optimize assets. It optimizes whats in assets/build and then adds them to assets/optimized
```

## Repository overview
```bash
├── cmd
├── internal (services)[just a simple database example]
│   ├── dbsql
│   │   ├── dbconn (holds the database.db)
│   │   └── user
│   ├── email (to send codes)
│   └── security (left to the user to configure)
│       ├── cookies
│       ├── crypt
│       ├── jwt
│       ├── tokens
│       └── validate
├── optimize (to optimize assets and !remember to set assets paths here)[funcs for cantenating and minifying assets and images]
├── src (app)[meat and bones of the application]
│   ├── funcmaps (functation for templates)
│   ├── handler
│       ├── get
│       │    ├── home (form to start)
│       │    └── loginemail (logging in with email)
│       └── post
│            └── createuser (post request to create a user)
│   ├── routes
│   └── server

```

<h3 align="left">Support:</h3>
<p><a href="https://ko-fi.com/zacharyendrulat98451"> <img align="left" src="https://cdn.ko-fi.com/cdn/kofi3.png?v=3" height="50" width="210" alt="zacharyendrulat98451" /></a></p><br><br>




## Special thanks
* [Go Team because they are gods](https://github.com/golang/go/graphs/contributors)
* [Creators of go echo](https://github.com/labstack/echo/graphs/contributors)
* [Creators of go Viper](https://github.com/spf13/viper/graphs/contributors)
* [Creators of sqlite and the go sqlite](https://gitlab.com/cznic/sqlite/-/project_members)
* [Creator of go-ps ](https://github.com/mitchellh/go-ps/graphs/contributors)