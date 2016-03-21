#!/usr/bin/env bash
set -e
set -x


export GOPATH=$(pwd)/go
export OUTPUTDIR=$(pwd)/build-artifact
cd go/src/github.com/karampok/gocf #defined in build-artifacts.yml

clean() {
    cd -
}


export GOBIN="$GOPATH/bin"
export PATH=$PATH:$GOBIN
check_go_version() {
    version=$(go version)
    regex="(go1.5.[0-9])"
    if ! [[ $version =~ $regex ]]; then 
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


[ -d $OUTPUTDIR ] || mkdir $OUTPUTDIR
make build 
[ "$?" -ne 0 ] && { echo "cannot build gocf binaries";clean; exit 1; }

ls $OUTPUTDIR

clean
echo "All green"
exit 0
