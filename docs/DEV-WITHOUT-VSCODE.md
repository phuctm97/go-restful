## Development without Visual Studio Code

Without VSCode, developers will not be able to achieve a full-time local-quality VSCode-powered containerized development environment. However, if there's any reason that you can not or do not want to work with Visual Studio Code, you can still start a containerized development environment and start working on that or even build your own development solution on top of that.

To start the development container, issue following command:

``` shell
./scripts/start-devcontainer.sh
```

For the first run, the script will build an image based on `Dockerfile`, the process may take a while. Latter runs will reuse the prebuilt image.

Once the image was built successfully, it will launch a container in interactive mode and mount the repository's root directory to `/go/src/github.com/the-evengers/go-restful` within the container, so that developers can make changes to the repository locally and have those changes automatically reflected into the container.

In interactive mode within the container, developers can issue every `go` CLI commands, e.g:

``` shell
# To start the application.
go run main.go
```
