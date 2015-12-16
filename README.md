# Go Binary App Template for Cloudfoundry

This is a template for deploying golang binaries apps in CF.
The template includes:

1. Local docker develpment
2. Connection tutorial to a mysql database 
3. Bindata/Gomigrate to bundle DB migrations directly to the binary

## Prerequisites

The following assumes you have a working, recent version of Go installed, and
you have a properly set-up Go workspace. Working on osx will require crosscompile. In a nutshell:

 * Install go (tested with 1.5.x)
 * Install docker for local development system
 * Registered to public or selfhosted CF service


## Running Locally Using Docker Containers
To start, given a working GOPATH , execute 

```
go get github.com/karampok/gocf
go get -u github.com/jteeuwen/go-bindata/...
go get github.com/tools/godep
make localpush
```

The `make localpush` simply 


* Binds the static data needed for db migration using [go-bindata](https://github.com/jteeuwen/go-bindata) 
* Creates the binary `./bin/gocf` (**note** crosscompile on osx)
* Starts a MariaDB docker container 

 ```
 docker run -d --name mariadb --env-file ./mariadb.env  -p 3306:3306/tcp mariadb
 ```
 
* Starts our app in specific cf container

 ```
 docker run -v ./gocf/bin:/opt/bin  --env-file ./env.list -p 4000:4000  --link mariadb:mariadb  -it cloudfoundry/cflinuxfs2 /opt/bin/gocf
 ``` 

Review the logs directly from the container.

Visit localhost:4000 in the browser.

Connect to the database by executing `make db-client`


## Running on CF  [[Swisscom Web Service](https://developer.swisscom.com/)]

Log in

```
cf login -a https://api.lyra-836.appcloud.swisscom.com
```

Create up a mysql Service.

```
cf create-service mariadb small mysql
```

Push the app. The manifest assumes you called your MariaDB instance 'mysql'.

```
make push
```

The `make push` simply 

1. Binds the static data needed for migration using [go-bindata](https://github.com/jteeuwen/go-bindata) 
2. Creates the binary `./bin/gocf` (**note** crosscompile on osx)
3. Pushes the app by executing `cf push  -c './gocf'   -b https://github.com/cloudfoundry/binary-buildpack.git`

Review the logs by executing `cf logs --recent kka-gocf` 

Visit kka-gocf.scapp.io in the browser.

Connect to the database for debugging using connector plugin

```
cf service-connector 3306 <serviceIP>:3306  &
mysql -h 127.0.0.1 -u <user> -p  <pass>
```

Documentation how to install the plugin  at [link](https://docs.developer.swisscom.com/services/services/managing-services.html)


## More Infos
[https://docs.cloudfoundry.org/buildpacks/binary/index.html]()

