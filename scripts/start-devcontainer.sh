#!/bin/bash

# Exit immediately on failure.
set -e

# Derive local workspace's absolute path from script's path.
local_workdir=$(cd $(dirname $(dirname "${BASH_SOURCE[0]}")) >/dev/null 2>&1 && pwd)

main() {
  local image_name=vsc-go-restful
  local container_name=$image_name
  # Mounted location within containers of local workspace
  # should be updated respectively to module importing schema using in Go source files,
  # and remote repository location if necessary (recommended).
  local container_workdir=/go/src/github.com/the-evengers/go-restful

  # Build image (if necessary).
  docker build --rm -t $image_name $local_workdir

  # Start container.
  docker run --rm -it \
    --name $container_name \
    --volume $local_workdir:$container_workdir \
    --workdir $container_workdir \
    --publish 8000:8000 \
    $image_name
}

main
