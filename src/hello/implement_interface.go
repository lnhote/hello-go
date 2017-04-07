package main

import "fmt"

type PayAction interface {
	payBill()
}

type CreditCard struct {
	name string
	cardNo string
}
// This method means CreditCard implements the interface PayAction,
// but we don't need to explicitly declare that it does so.
func (payer *CreditCard) payBill() {
	fmt.Printf("%s pay the bill with credit card [%s].\n", payer.name, payer.cardNo)
}

type AlipayAccount struct {
    name string
    userId int
}
func (payer *AlipayAccount) payBill() {
    fmt.Printf("%s pay the bill with alipay account [%d].\n", payer.name, payer.userId)
}

func main() {
	var creditCardPay PayAction = &CreditCard{"Alice", "12345"}
	var alipay PayAction = &AlipayAccount{"Bob", 10001}
	creditCardPay.payBill()
	alipay.payBill()
}
