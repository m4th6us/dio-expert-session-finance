package http

import (
	"net/http"

	"github.com/m4th6us/dio-expert-session-finance/adapter/http/actuator"
	"github.com/m4th6us/dio-expert-session-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
