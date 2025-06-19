package models

type BalanceOperation struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

var (
	Withdraw = "WITHDRAW"
	TopUp    = "TOPUP"
)
