#!/bin/bash
set -u

IMAGE=$1

docker rmi -f ${IMAGE} 2> /dev/null

exit 0
