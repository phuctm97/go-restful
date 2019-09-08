#!/bin/bash

# Exit immediately on failure.
set -e

# Derive local workspace's absolute path from script's path.
local_workdir=$(cd $(dirname $(dirname "${BASH_SOURCE[0]}")) >/dev/null 2>&1 && pwd)

main() {
  local image_name=go-restful-dev
  local container_name=go-restful-dev
  local container_workdir=/go/src/github.com/the-evengers/go-restful

  # Build image (if necessary).
  docker build --rm -t $image_name $local_workdir

  # Start container.
  docker run --rm -dt \
    --name $container_name \
    --volume $local_workdir:$container_workdir \
    --workdir $container_workdir \
    $image_name
}

main
