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







