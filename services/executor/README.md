# Goxpert

Goxpert which is graphql based Apis
The Contribution Guidelines contains information (such as local development setup) for contributing to this project.

## Development Setup

You will need the following installed in your development environment.

# Golang

The project is built using golang. You can find instructions on how to install it here.

# mux

This is an gorila mux router to acessing defferant resources. The API uses the mux router package.

# gomock

The API uses gomock to generates mocks from interfaces for easier testing. This project uses golang in-built testing package and gomock library

# Docker

To execute local integration tests you will need to install Docker, you can find the instructions here. yoy can do mongo installation by using docker

# Postman

This project contains a set of postman tests that can be both used for automated testing with tools like newman as well as for development and debugging. You can find installation instructions here.

# mysqldb

This project utilizes mysqldb for excuting seed scripts for setting up a local dev environment.

# Makefile

This project strives to use the Makefile as a way for the team to utilize individual commands related to the project, e.g. gather dependencies is a seperate one from build. These individual commands can then be aggregated to achieve more complicated operations.

## Runing File User Service

Creating a Build

    	$ go build -o user/bin/user-service user/cmd/main.go

Run user service

    	$ ./user/bin/user-service mysqldb-host localhost -mysqldb-user rahul -mysqldb-pass password -mysqldb-addr=3306
