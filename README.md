# Demo demonstrating multiple API GWs for Services

## List of Services
API GW Threats - api_svc_threats
API GW Janus - api_svc_cfw
API GW DNS - api_svc_dns

## Create Persistence
```sh
$ createdb demo
```

```sh
$ bee api api_svc_threats -driver=postgres -conn=postgres://seizadi:seizadi@127.0.0.1:5432/demo?sslmode=disable
$ cd api_svc_threats
$ bee generate scaffold host -fields="fqdn:string" -driver=postgres -conn=postgres://seizadi:seizadi@127.0.0.1:5432/demo?sslmode=disable
$ bee generate scaffold client_name -fields="name:string, ip:string" -driver=postgres -conn=postgres://seizadi:seizadi@127.0.0.1:5432/demo?sslmode=disable
$ bee generate scaffold zone -fields="fqdn:string" -driver=postgres -conn=postgres://seizadi:seizadi@127.0.0.1:5432/demo?sslmode=disable

$ bee api api_svc_cfw -driver=postgres -conn=postgres://seizadi:seizadi@127.0.0.1:5432/demo?sslmode=disable
$ bee api api_svc_dns -driver=postgres -conn=postgres://seizadi:seizadi@127.0.0.1:5432/demo?sslmode=disable
```
## Configure instances
Go to the respective projects and modify the port in <project>/conf/app.conf from default 8080 to a unique one, we will use 8001, 8002, and 8003. Here is the app.conf for the api_svc_cfw project:
```ini
appname = api_svc_cfw
httpport = 8001
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true

```

## Run the projects

For now I will run them manually but we will put them in a cluster and run them at once with NGNIX and Database.

Run the project with the Swagger Documentation (http://127.0.0.1:8001/swagger/.)

```sh
$ cd api_svc_cfw
$ bee run -downdoc=true -gendoc=true
```

## Load some sample data
```sh
$ curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{ "fqdn": "www.google.com"}' 'http://localhost:8001/v1/host/'
$ curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{ "name": "soheil", "ip":"1.2.3.4"}' 'http://localhost:8001/v1/client_name/'
$ curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json' -d '{ "fqdn": "webserver.acme.com"}' 'http://localhost:8001/v1/zone/'
```

## Read the sample data
```sh
$ curl -X GET --header 'Accept: application/json' 'http://localhost:8001/v1/client_name/'
$ curl -X GET --header 'Accept: application/json' 'http://localhost:8002/v1/host/'
$ curl -X GET --header 'Accept: application/json' 'http://localhost:8003/v1/zone/'
```





