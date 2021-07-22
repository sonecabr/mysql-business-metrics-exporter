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
