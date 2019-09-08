#!/bin/bash

# Exit immediately on failure.
set -e

# Find absolute path of local workspace based on script's path.
local_workdir=$(cd $(dirname $(dirname "${BASH_SOURCE[0]}")) >/dev/null 2>&1 && pwd)

main() {
  local image_name=go-restful-dev
  local container_name=go-restful-dev
  local container_workdir=/go/src/github.com/the-evengers/go-restful

  # Build image (if necessary).
  docker build --rm -t $image_name $local_workdir

  # Start container.
  docker run --rm -it \
    --name $container_name \
    --volume $local_workdir:$container_workdir \
    --workdir $container_workdir \
    $image_name
}

main

