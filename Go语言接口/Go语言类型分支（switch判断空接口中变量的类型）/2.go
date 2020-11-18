package main

// 电子支付方式
type Alipay struct {
}

// 为Alipay添加CanUseFaceID()方法, 表示电子支付方式支持刷脸
func (a *Alipay) CanUseFaceID() {

}

// 现金支付方式
type Cash struct {
}

// 为Cash添加Stolen()方法, 表示现金支付方式会出现偷窃情况
func (a *Cash) Stolen() {

}

// 具备刷脸特性的接口
type CantainCanUseFaceID interface {
	CanUseFaceID()
}

// 具备被偷特性的接口
type ContainStolen interface {
	Stolen()
}

// 打印支付方式具备的特点
func print(payMethod interface{}) {
    switch payMethod.(type) {
    case CantainCanUseFaceID:  // 可以刷脸
        fmt.Printf("%T can use faceid\n", payMethod)
    case ContainStolen:  // 可能被偷
        fmt.Printf("%T may be stolen\n", payMethod)
    }
}

func main (){
	 //使用电子支付判断
	  print(new(Alipay))

	// 使用现金判断
	  print(new(Cash))
}
