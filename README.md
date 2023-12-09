# go_simple_api
This is simple api that get some data (study for clean architecture)

# how to start
```shell
$ make run-db
$ make db-migrate
$ make run-server
```

# how to check them
```shell
$ curl http://localhost:3000/healthcheck
-> Health Check is OK% 
$ curl http://localhost:3000/drama
-> You can get some data! 
```
