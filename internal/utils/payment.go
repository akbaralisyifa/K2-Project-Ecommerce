package utils

import (
	"ecommerce/config"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func Payment(orderID uint, totalAmount uint64) (string, error) {
	serverKey := config.ImportSetting().MidTransKey
	s := snap.Client{}

	s.New(serverKey, midtrans.Sandbox)

	req := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "payment-" + strconv.Itoa(int(orderID)),
			GrossAmt: int64(totalAmount),
		},
	}
	res, err := s.CreateTransaction(&req)
	if err != nil {
		return "", err
	}
	return res.RedirectURL, nil
}
