package main

import "fmt"

type Payment interface {
	Pay(amount float64) (string, error)
	Refund(transactionID string, amount float64) (string, error)
	Query(transactionID string) (string, error)
}

type Alipay struct {
	AppID      string
	AppSecret  string
	MerchantID string
}

type WechatPay struct {
	AppID      string
	AppSecret  string
	MerchantID string
}

func (a *Alipay) Pay(amount float64) (string, error) {
	// Implementation for Alipay payment
	return "Alipay payment successful", nil
}

func (a *Alipay) Refund(transactionID string, amount float64) (string, error) {
	// Implementation for Alipay refund
	return "Alipay refund successful", nil
}

func (a *Alipay) Query(transactionID string) (string, error) {
	// Implementation for Alipay query
	return "Alipay transaction details", nil
}

func (a *WechatPay) Pay(amount float64) (string, error) {
	// Implementation for Alipay payment
	return "WechatPay payment successful", nil
}

func (w *WechatPay) Refund(transactionID string, amount float64) (string, error) {
	// Implementation for Alipay refund
	return "WechatPay refund successful", nil
}

func (w *WechatPay) Query(transactionID string) (string, error) {
	// Implementation for Alipay query
	return "WechatPay transaction details", nil
}

func ProcessPayment(p Payment, amount float64) (string, error) {
	fmt.Println(p.Pay(amount))
	fmt.Println(amount)
	return "", nil
}

func main() {
	aplipay := Alipay{
		AppID:      "123456",
		AppSecret:  "asdasdasd",
		MerchantID: "123",
	}

	wechatPay := WechatPay{
		AppID:      "654321",
		AppSecret:  "dsadsadsad",
		MerchantID: "321",
	}

	transactions := []struct {
		Payment Payment
		Amount  float64
		Name    string
	}{
		{Payment: &aplipay, Amount: 100.0, Name: "Alipay"},
		{Payment: &wechatPay, Amount: 200.0, Name: "WechatPay"},
	}

	for _, pp := range transactions {
		ProcessPayment(pp.Payment, pp.Amount)
	}
}
