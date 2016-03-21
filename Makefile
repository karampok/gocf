all: build

GIT_COMMIT = $(shell git describe --always --dirty)
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ" | tr -d '\n')
GGOPATH = $(shell pwd)/Godeps/_workspace:$(GOPATH)
OUTPUTDIR ?= .

setup:
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/axw/gocov/gocov
	go get -u github.com/AlekSi/gocov-xml

build: bindata
	GOPATH=$(GGOPATH) GOARCH=amd64 GOOS=linux go build -ldflags "-X main.buildstamp=$(BUILD_TIME) -X main.githash=$(GIT_COMMIT)" -o $(OUTPUTDIR)/gocf-5

bindata:
	go-bindata -o  migrations/bindata.go -pkg migration migrations_data/

test:
	GOPATH=$(GGOPATH) go test  ./...

cfpush: build
	cf push  -c './gocf'   -b https://github.com/cloudfoundry/binary-buildpack.git

localpush: build db-start
	docker run -v ${PWD}/bin:/opt/bin  --env-file ./cf.env -p 4000:4000  --link mariadb:mariadb  -it cloudfoundry/cflinuxfs2 /opt/bin/gocf

db-start: 
	@echo  "$(OK_COLOR)==> Starting the mariadb $(NO_COLOR)"
	docker run -d --name mariadb --env-file ./mariadb.env  -p 3306:3306/tcp mariadb  2>/dev/null || echo "MariaDB is already running (make db-stop to start from scratch)"

db-stop: 
	@echo  "$(OK_COLOR)==> Stoping the mariadb  $(NO_COLOR)"
	docker rm -f  mariadb || exit 0

db-client:
	@echo  "$(OK_COLOR)==> Start a client $(NO_COLOR)"
	docker run -it --env-file ./mariadb.env  --link mariadb:mysql --rm mariadb sh -c 'exec mysql -h"$$MYSQL_PORT_3306_TCP_ADDR" -P"$$MYSQL_PORT_3306_TCP_PORT" -u"$$MYSQL_USER" -p"$$MYSQL_PASSWORD"'

db-client-root:
	@echo  "$(OK_COLOR)==> Start a client $(NO_COLOR)"
	docker run -it --env-file ./mariadb.env  --link mariadb:mysql --rm mariadb sh -c 'exec mysql -h"$$MYSQL_PORT_3306_TCP_ADDR" -P"$$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$$MYSQL_ROOT_PASSWORD"'
