package utils

import (
	"log"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransPayment struct {
	snapClient snap.Client
}

func NewMidtransPayment(serverKey string) midtransPayment {
	return midtransPayment{
		snapClient: snap.Client{ServerKey: serverKey, Env: midtrans.Sandbox},
	}
};


func (mp *midtransPayment) RequestPayment(orderId string, amount int64)(string, error){
	req := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: orderId,
			GrossAmt: amount,
		},
	};

	res, err := mp.snapClient.CreateTransaction(&req)
	if err != nil {
		log.Println("midtrans error :", err.Error())
		return "", err
	};

	return res.RedirectURL, nil
}