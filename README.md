# Mysql PTO exporter
Just an example on how to export results of mysql queries as prometheus metrics

## build
```
go build -o main
```

## run
```
./main
```

## envs
```
export MYSQL_USERNAME=user
export MYSQL_ADDR=127.0.0.1
export MYSQL_DB=payments
export MYSQL_PORT=3306 
```

## run local mysql
```
docker-compose up
```
Then use the file migrations/000-base.sql to populate local db

## run using docker-compose
- 1 
```
docker-compose up
```
- 2
Using a mysql client connect to local mysql ( workbench for instance )
host: 127.0.0.1
username: user
password: pass
database: payments
- 3
Run script available on migrations/000-base.sql using your connected mysql client

## to check metrics
curl -XGET http://localhost:8090/metrics

