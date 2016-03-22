#!/usr/bin/env bash
set -e
#set -x


export GOPATH=$(pwd)/go
export VERSION=$(cat version/version)
export OUTPUTDIR=$(pwd)/build-artifact
[ -d $OUTPUTDIR ] || mkdir $OUTPUTDIR
export GOBIN="$GOPATH/bin"
export PATH=$PATH:$GOBIN

cd go/src/github.com/karampok/gocf #defined in build-artifacts.yml

clean() {
    :
}

check_go_version() {
    goversion=$(go version)
    regex="(go1.5.[0-9])"
    if ! [[ $goversion =~ $regex ]]; then 
        echo "go is not installed or wrong version ";
        clean
        exit 1
    fi
}


check_go_version

make setup

NEEDED_COMMANDS="" 
for cmd in ${NEEDED_COMMANDS} ; do
    if ! command -v ${cmd} &> /dev/null ; then
    echo Please install ${cmd}!
    exit 1
    fi
done


make build 
[ "$?" -ne 0 ] && { echo "cannot build gocf binaries";clean; exit 1; }

clean
echo "All green"
exit 0
