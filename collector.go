package main

import (
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"strconv"
)

type PtoMetric struct {
	amount int64
	qtt int64
}

type PtoMetricsCollector struct {
	ptometrics   *prometheus.Desc
}

func ptoCheck() *PtoMetric {
	log.Info("collecting metrics...")
	username := string(os.Getenv("MYSQL_USERNAME"))
	password := string(os.Getenv("MYSQL_PASSWORD"))
	addr := string(os.Getenv("MYSQL_ADDR"))
	port := string(os.Getenv("MYSQL_PORT"))
	database := string(os.Getenv("MYSQL_DB"))

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, addr, port, database))

	if err != nil {
		panic(err.Error())
	}

	var line PtoMetric

	err = db.QueryRow("select sum(amount * qt) as amount from paymentfirst").Scan(&line.amount)
	err = db.QueryRow("select sum(qt) as qtt from paymentfirst").Scan(&line.qtt)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return &PtoMetric {
		amount: line.amount,
		qtt: line.qtt,
	}

}


func newPTOCollector() *PtoMetricsCollector {
	return &PtoMetricsCollector {
		ptometrics: prometheus.NewDesc("ptometrics", "show business counters ", []string{"amount", "qtt"}, nil),
	}
}

func (collector *PtoMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.ptometrics
}

func (collector *PtoMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	metric := ptoCheck()
	ch <- prometheus.MustNewConstMetric(collector.ptometrics, prometheus.GaugeValue, 1, strconv.FormatInt(metric.amount, 10), strconv.FormatInt(metric.qtt, 10))

}

func main() {
	ptoCollector := newPTOCollector()
	prometheus.MustRegister(ptoCollector)
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Starting prometheus exporter on port :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}