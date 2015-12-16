all: binary

binary: bindata
	rm -f ./bin/*
	GOARCH=amd64 GOOS=linux godep go build -ldflags "-X main.buildstamp `date '+%Y-%m-%d_%H:%M:%S'` -X main.githash `git rev-parse HEAD`" -o ./bin/gocf

bindata:
	go-bindata -o  migrations/bindata.go -pkg migration migrations_data/

test:
	GOARCH=amd64 GOOS=linux godep go test -v ./...

push: binary
	cf push  -c './gocf'   -b https://github.com/cloudfoundry/binary-buildpack.git


localpush: binary db-start
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
