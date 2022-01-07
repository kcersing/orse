package order

import (
	"context"
	"fmt"
)

func main() {
	//通过订单ID 或 其他的信息查询到订单状态
	orderStatus := StatusDefault
	orderMachine, err := NewFSM(orderStatus)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//创建订单，订单创建成功后再给用户发短信
	if _, err = orderMachine.Call(
		EventCreate,
		WithOrderId(1),
		WithOrderName("测试订单"),
		WithHandlerSendSMS(sendSMS),
	); err != nil {
		fmt.Println(err.Error())
	}
	//修改订单
	if _, err = orderMachine.Call(EventModify); err != nil {
		fmt.Println(err.Error())
	}
	//确认订单
	if _, err = orderMachine.Call(EventConfirm); err != nil {
		fmt.Println(err.Error())
	}
	//修改订单
	if _, err = orderMachine.Call(EventModify); err != nil {
		fmt.Println(err.Error())
	}
	//确认订单
	if _, err = orderMachine.Call(EventConfirm); err != nil {
		fmt.Println(err.Error())
	}
	//支付订单
	if _, err = orderMachine.Call(EventPay); err != nil {
		fmt.Println(err.Error())
	}

}


func CreateOrder(ctx *context.Context){
	//通过订单ID 或 其他的信息查询到订单状态
	orderStatus := StatusDefault
	orderMachine, err := NewFSM(orderStatus)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//创建订单，订单创建成功后再给用户发短信
	if _, err = orderMachine.Call(
		EventCreate,
		WithOrderId(1),
		WithOrderName("测试订单"),
		WithHandlerSendSMS(sendSMS),
	); err != nil {
		fmt.Println(err.Error())
	}

}

func UpdateOrder(ctx *context.Context)  {
	
}






// sendSMS 发送短信
func sendSMS(mobile, content string) error {
	fmt.Println(fmt.Sprintf("发送短信，给(%s)发送了(%s)", mobile, content))
	return nil
}
