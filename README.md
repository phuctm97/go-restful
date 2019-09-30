# Go RESTful

🚀 A real world production-grade RESTful Web Services proof-of-concept project.

## Objectives

- An optimized Go implementation follows [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), provides mechanism to declare Entities, Use cases, and External Services (e.g. data access).

- An optimized Go implementation provides mechanism to expose Entities and Use Cases as [RESTful Web Services](https://en.wikipedia.org/wiki/Representational_state_transfer).

- An optimized Go implementation of Token-based [Authentication](https://en.wikipedia.org/wiki/Authentication) and [Authorization](https://en.wikipedia.org/wiki/Authorization).

- An optimized Go implementation provides abstract mechanism to access [Relational Databases](https://en.wikipedia.org/wiki/Relational_database).

- An optimized Go Development Environment with [Git](https://git-scm.com/), [Docker](https://www.docker.com/), [Go Modules](https://github.com/golang/go/wiki/Modules), Go Debuggers ([GDB](https://golang.org/doc/gdb)/[Delve](https://github.com/go-delve/delve)) and popular code editors ([VSCode](https://www.google.com/search?client=safari&rls=en&q=vscode&ie=UTF-8&oe=UTF-8)/[GoLand](https://www.jetbrains.com/go/)/[Vim](https://github.com/fatih/vim-go)).

- An optimized CI/CD Solution with [Github Actions](https://github.com/features/actions) and [AWS](https://aws.amazon.com/).

- An optimized Distribution Solution with [Github Releases](https://help.github.com/en/enterprise/2.16/user/articles/about-releases) and [Github Package Registry](https://help.github.com/en/articles/about-github-package-registry).

- A scalable and highly-available Production Deployment Solution over [AWS](https://aws.amazon.com/) using [Terraform](https://www.terraform.io/).

- An optimized Staging Environment replicating Production Environment for testing purposes.

- An optimized Issues Tracking mechanism with [Github Project](https://github.com/features/project-management/), [Issues](https://help.github.com/en/articles/about-issues) and [Pull Requests](https://help.github.com/en/articles/about-pull-requests).

- Continual improvements.

## Issues Tracking

[![Issue Tracking](https://img.shields.io/static/v1?label=issue%20tracking&message=Github%20project&color=lightgrey)](https://github.com/the-evengers/go-restful/projects/1)
[![Open issues](https://img.shields.io/github/issues/the-evengers/go-restful)](https://github.com/the-evengers/go-restful/issues) [![Closed issues](https://img.shields.io/github/issues-closed/the-evengers/go-restful)](https://github.com/the-evengers/go-restful/issues?q=is%3Aissue+is%3Aclosed) [![Open PRs](https://img.shields.io/github/issues-pr/the-evengers/go-restful)](https://github.com/the-evengers/go-restful/pulls) [![Closed PRs](https://img.shields.io/github/issues-pr-closed/the-evengers/go-restful)](https://github.com/the-evengers/go-restful/pulls?q=is%3Apr+is%3Aclosed)

This project use Github project, issues and pull requests to manage and track issues. Refer to [this Github project](https://github.com/the-evengers/go-restful/projects/1) for further details.

## Development

### Requirements

Local development machines need to have following tools installed and working properly:

- [Docker](https:://www.docker.com) for running a full-time containerized development environment.

- [Visual Studio Code](https://code.visualstudio.com) with [Remote - Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) for writing code with Intellisense, running and debugging code within containers.

Windows users need to additionally have an Unix-shell emulator to be able to run utility scripts ([Git Bash](https://gitforwindows.org) is recommended).

### Usage

#### Develop with Visual Studio Code

With all of above requirements fullfiled, developers can experience a full-time local-quality VSCode-powered containerized development environment by just opening the repository in VSCode container mode.

Go [here](docs/DEV-WITH-VSCODE.md) for further details.

#### Develop without Visual Studio Code

Without VSCode, developers will not be able to achieve a full-time local-quality VSCode-powered containerized development environment. However, if there's any reason that you can not or do not want to work with Visual Studio Code, you can still start a containerized development environment and start working on that or even build your own development solution on top of that.

Go [here](docs/DEV-WITHOUT-VSCODE.md) for further details.

#### Project structure

Quick overview of project structure, components and their roles.

```
├── 📁.vscode/              # VSCode configurations.
├── 📦common/               # Common, utility Go components.
├── 📁docs/                 # Documentation & assets.
├── 📁scripts/              # Utility scripts.
├── 📦users/                # Users-related Go components.
├── 📄.devcontainer.json    # VSCode Remote-Containers configuration.
├── 📄.gitignore
├── 📄Dockerfile            # Instructions to build development Docker image.
├── 📄go.mod                # Go module configuration.
├── 📄go.sum                # Go module checksum/lock.
├── 📖LICENSE
├── 📖README.md
├── 🚀main.go               # Application's main entry point.

```

### License

[![Apache License 2.0](https://img.shields.io/github/license/the-evengers/go-restful)](https://github.com/the-evengers/go-restful/blob/master/LICENSE) ![Copyright © The Evengers](https://img.shields.io/static/v1?label=copyright&message=The%20Evengers&color=lightgrey)

Copyright © The Evengers. All rights reserved.

This project is licensed under the [Apache License 2.0](https://github.com/the-evengers/go-restful/blob/master/LICENSE) and is available for free.
