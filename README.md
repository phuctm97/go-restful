# Go RESTful

A real world production-grade RESTful proof-of-concept project.

## Objectives

- An optimized Go implementation follows clean architecture, provides mechanism to declare entities, use cases - - and external services (e.g. data access).
- An optimized Go implementation provides mechanism to exposes entities and use cases as RESTful web services.
- An optimized Go implementation of token-based authentication and authorization.
- An optimized Go implementation provides abstract mechanism to access relational databases.
- An optimized Go development environment with Git, Docker, Golang/dep, Golang debuggers (GDB/Delve) and editors (VSCode/GoLand).
- A scalable and highly available production deployment solution over AWS cloud using Terraform.
- An optimized staging environment replicating production environment for testing purposes.
- An optimized CI/CD solution with CircleCI and AWS.
- An optimized issues tracking mechanism with Github projects, issues and pull requests.
- Continual improvements.

## Issues Tracking

![Git project](https://img.shields.io/static/v1?label=issue%20tracking&message=Github%27s%20project&color=lightgrey) ![GitHub issues](https://img.shields.io/github/issues/the-evengers/go-restful) ![GitHub closed issues](https://img.shields.io/github/issues-closed/the-evengers/go-restful) ![GitHub PRs](https://img.shields.io/github/issues-pr/the-evengers/go-restful) ![GitHub closed PRs](https://img.shields.io/github/issues-pr-closed/the-evengers/go-restful)

This project use Github project, issues and pull requests to manage and track issues. Refer to [Go RESTful](https://github.com/the-evengers/go-restful/projects/1) for details.

## Development

### Requirements

This project uses *Docker* to simplify the process of setting up local development environment. Therefore, [Docker](https://www.docker.com) is the only requirement to run this project in development mode.

**Note**: Developers are still able to set up their own local development environment directly on their local machine or use other environment management technology. However, those methods are beyond this project's scope and therefore some later documented features might not work properly.

### Usage

To start development environment, run following command:

``` shell
./scripts/start-dev.sh
```

Above command will start a Docker container in interactive mode and mount the repo's base directory to `/go/src/github.com/the-evengers/go-restful` inside the container, so that developers can make changes to the repo from their local machine and have those changes automatically reflected into the container. 

In interactive mode within the container, developers can execute every `go` CLI commands, e.g:

``` shell
# To start the application
go run main.go
```

### License

![GitHub](https://img.shields.io/github/license/the-evengers/go-restful) ![The-Evengers](https://img.shields.io/static/v1?label=copyright&message=The%20Evengers&color=lightgrey)

This project is publicly developed and maintained by *The Evengers*, an open source community organization, and therfore is always available for free.
