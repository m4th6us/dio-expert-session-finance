package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/m4th6us/dio-expert-session-finance/model/transaction"
	"github.com/m4th6us/dio-expert-session-finance/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2020-04-05T11:11:00"),
		},
	}

	json.NewEncoder(w).Encode(transactions)

}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = io.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Println(res)

}
