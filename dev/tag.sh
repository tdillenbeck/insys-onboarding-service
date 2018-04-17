#!/bin/bash

set -e

BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo "Current branch is ${BRANCH}"

GITREV=$(git rev-list HEAD --count)

TAG=v0.0.${GITREV}

if [ $BRANCH != "master" ]; then
  TAG="${TAG}-${BRANCH}"
fi

echo "Tagging as ${TAG}"

git tag ${TAG}

git push origin ${TAG}
