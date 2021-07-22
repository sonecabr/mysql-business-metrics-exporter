FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN apk add git && \
go get github.com/sirupsen/logrus && \
go get github.com/prometheus/client_golang/prometheus && \
go get github.com/prometheus/client_golang/prometheus/promhttp && \
go get github.com/go-sql-driver/mysql && \
go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]