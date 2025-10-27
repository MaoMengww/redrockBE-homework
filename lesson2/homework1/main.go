package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int //库存
}

func (p Product) TotalValue() float64 {
	sum := float64(p.Stock) * p.Price
	return sum

}

func (p Product) IsInStock() bool {
	if p.Stock > 0 {
		return true
	}
	return false
}

func (p Product) Info() string {
	info := fmt.Sprintf("商品： %v, 单价: ￥%v, 库存： %v件", p.Name, p.Price, p.Stock)
	return info
}

func Restock(p *Product, amount int) {
	p.Stock = p.Stock + amount
}

func Sell(p *Product, amount int) (success bool, message string) {
	if p.Stock-amount >= 0 {
		p.Stock = p.Stock - amount
		return true, "售卖成功"
	}
	return false, "库存不足"
}

func main() {
	p := Product{
		Name:  "Go编程书",
		Price: 89.5,
		Stock: 10,
	}
	result, message := Sell(&p, 5)
	if result {
		fmt.Println(message, "剩余库存：", p.Stock)
	} else {
		fmt.Println("失败," + message)
	}
	Restock(&p, 20)
	fmt.Printf("进货%v本, 当前库存：%v\n", 20, p.Stock)
	result, message = Sell(&p, 30)
	if result {
		fmt.Println(message, "剩余库存：", p.Stock)
	} else {
		fmt.Println("失败," + message)
	}
	res := p.Info()
	fmt.Printf("%v, 总价：%v", res, p.TotalValue())
}
