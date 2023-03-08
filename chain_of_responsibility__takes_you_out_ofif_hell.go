package main

import "fmt"

type SellInfo struct {
	Price      float64
	OrderCount int
	TotalCount int
	MemberShip int
}

func ugly() {
	var a = SellInfo{
		Price:      1.0,
		OrderCount: 1,
		TotalCount: 20,
		MemberShip: 1,
	}
	if a.Price > 0 {
		print("true")
	}
	if a.TotalCount > a.OrderCount {
		println("true")
	} else {
		println("false")
	}
	if a.MemberShip == 1 {
		println("true")
	}
	if a.Price < 100 && a.MemberShip == 2 {
		println("true")
	}
}

type Rule interface {
	Check(sellInfo *SellInfo) bool
}

func Chain(info *SellInfo, rules ...Rule) bool {
	for _, r := range rules {
		if !r.Check(info) {
			return false
		}
	}
	return true
}
func elegant() {
	sellInfo := &SellInfo{
		Price:      1.0,
		OrderCount: 1,
		TotalCount: 20,
		MemberShip: 1,
	}

	rules := []Rule{
		&PriceRule{},
		&OrderCountRule{},
		&MemberShipRule{},
		&DiscountRule{},
		//...

	}
	r := Chain(sellInfo, rules...)
	fmt.Println(r)

}

type PriceRule struct{}

func (pr *PriceRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.Price > 0
}

type OrderCountRule struct{}

func (ocr *OrderCountRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.TotalCount > sellInfo.OrderCount
}

type MemberShipRule struct{}

func (msr *MemberShipRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.MemberShip == 1
}

type DiscountRule struct{}

func (dr *DiscountRule) Check(sellInfo *SellInfo) bool {
	return sellInfo.Price < 100 && sellInfo.MemberShip == 2
}

func main() {
	ugly()
	elegant()
}
