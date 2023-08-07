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
* [General info](#general-info)
* [Why build this?](#why-build-this)
* [Technologies](#technologies)
* [Setup](#setup)
* [Repository overview](#repository-overview)
* [Special thanks](#special-thanks)



## General info
This project is a template for gonew and is used for setting up a webserver using echo framework.


## Why build this?
* Go never changes
* It is a nice way to start out a webserver without doing much


## Technologies
Project is created with:
* [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) - database
* [go-ps](https://github.com/mitchellh/go-ps) - getting pids in all OS's
* [viper](github.com/spf13/cobra) - build cli commands
* [echo](github.com/labstack/echo/v4) - web framework to shorten code
* [sprig](https://github.com/Masterminds/sprig) - template functions

## Setup
Just use the new [gonew](https://go.dev/blog/gonew)

```
$ go install golang.org/x/tools/cmd/gonew@latest
$ gonew github.com/golangast/goservershell example.com/myserver
```

## Repository overview
```bash
├── cmd
├── internal (services)[just a simple database example]
│   ├── dbsql
│   │   ├── dbconn
│   │   └── gettable
│   └── security (left to the user to configure)
│       ├── cookies
│       ├── crypt
│       ├── jwt
│       ├── tokens
│       └── validate
├── src (app)[meat and bones of the application]
│   ├── funcmaps
│   ├── handler
│   ├── routes
│   └── server
├──database.db (sqlite database)
```

<h3 align="left">Support:</h3>
<p><a href="https://ko-fi.com/zacharyendrulat98451"> <img align="left" src="https://cdn.ko-fi.com/cdn/kofi3.png?v=3" height="50" width="210" alt="zacharyendrulat98451" /></a></p><br><br>




## Special thanks
* [Go Team because they are gods](https://github.com/golang/go/graphs/contributors)
* [Creators of go echo](https://github.com/labstack/echo/graphs/contributors)
* [Creators of go Viper](https://github.com/spf13/viper/graphs/contributors)
* [Creators of sqlite and the go sqlite](https://gitlab.com/cznic/sqlite/-/project_members)
* [Creator of go-ps ](https://github.com/mitchellh/go-ps/graphs/contributors)