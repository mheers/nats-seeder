#!/usr/bin/env bash

set -eo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" ; pwd -P)"

VERSION=$(date +%Y%m%d.%H%M%S)
pushd "${SCRIPT_DIR}/.." > /dev/null
echo "export VERSION=${VERSION}" > .VERSION
popd > /dev/null

TRG_PKG='main'
BUILD_TIME=$(date +"%Y%m%d.%H%M%S")
CommitHash=N/A
GoVersion=N/A
GitTag=N/A
GitBranch=N/A

if [[ $(go version) =~ [0-9]+\.[0-9]+\.[0-9]+ ]];
then
    GoVersion=${BASH_REMATCH[0]}
fi

GV=$(git tag || echo 'N/A')
if [[ $GV =~ [^[:space:]]+ ]];
then
    GitTag=${BASH_REMATCH[0]}
fi

GV=$(git branch || echo 'N/A')
if [[ $GV =~ [^[:space:]]+ ]];
then
    GitBranch=${BASH_REMATCH[0]}
fi

GH=$(git log -1 --pretty=format:%h || echo 'N/A')
if [[ GH =~ 'fatal' ]];
then
    CommitHash=N/A
else
    CommitHash=$GH
fi

export TRG_PKG
export BUILD_TIME
export CommitHash
export GoVersion
export GitTag
export GitBranch
export VERSION
