package payments

import "fmt"

type PayMethod interface {
	Pay(usd int) int
	Cansel(ID int)
}

type PaymetModul struct {
	paymethod PayMethod
}

type NewPaymetModul(paymethod PayMethod) PaymetModul struct{
	
}

func (p PaymetModul) Pay() {

}
func (p PaymetModul) Cansel() {}

func (p PaymetModul) Info() {
	fmt.Println("hj")
}
func (p PaymetModul) AllInfo() {}
