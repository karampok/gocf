all: binary

binary: bindata
	rm -f ./bin/*
	GOARCH=amd64 GOOS=linux godep go build -ldflags "-X main.buildstamp=`date '+%Y-%m-%d_%H:%M:%S'` -X main.githash=`git rev-parse HEAD`" -o ./bin/gocf

bindata:
	go-bindata -o  migrations/bindata.go -pkg migration migrations_data/

test:
	GOARCH=amd64 GOOS=linux godep go test -v ./...

push: binary
	cf push  -c './gocf'   -b https://github.com/cloudfoundry/binary-buildpack.git


localpush: binary 
	docker run -v ${PWD}/bin:/opt/bin  --env-file ./cf.env -p 4000:4000  \
		--link mariadb:mariadb  --link logstash:logstash \
	   	-it cloudfoundry/cflinuxfs2 /opt/bin/gocf

services-start: 
	#docker run -d -v "$PWD/config":/usr/share/elasticsearch/config elasticsearch
	docker run -d  -p 9200:9200 -p 9300:9300 elasticsearch || echo "elasticsearch is already running" 
	docker run -d --name logstash  -p 5000:5000 -p 9292:9292 logstash  logstash -e 'input { tcp { port => 5000 } } output { stdout { } }' || echo "Logstash is already running"
	docker run -d --name mariadb --env-file ./mariadb.env  -p 3306:3306/tcp mariadb  2>/dev/null || echo "MariaDB is already running" 
	sleep 15

services-stop:
	@docker rm -fv elasticsearch >/dev/null 2>&1 || exit 0
	@docker rm -fv logstash >/dev/null 2>&1 || exit 0
	@docker rm -fv mariadb >/dev/null 2>&1 || exit 0


db-client:
	@echo  "$(OK_COLOR)==> Start a client $(NO_COLOR)"
	docker run -it --env-file ./mariadb.env  --link mariadb:mysql --rm mariadb sh -c 'exec mysql -h"$$MYSQL_PORT_3306_TCP_ADDR" -P"$$MYSQL_PORT_3306_TCP_PORT" -u"$$MYSQL_USER" -p"$$MYSQL_PASSWORD"'

db-client-root:
	@echo  "$(OK_COLOR)==> Start a client $(NO_COLOR)"
	docker run -it --env-file ./mariadb.env  --link mariadb:mysql --rm mariadb sh -c 'exec mysql -h"$$MYSQL_PORT_3306_TCP_ADDR" -P"$$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$$MYSQL_ROOT_PASSWORD"'
