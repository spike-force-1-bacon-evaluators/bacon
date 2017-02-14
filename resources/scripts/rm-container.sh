#!/bin/bash
set -e
set -u

CONTAINER=$1

docker stop ${CONTAINER} 2> /dev/null || true && \
  docker rm ${CONTAINER} 2> /dev/null || true
