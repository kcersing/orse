package order

import (
	"fmt"
)

var (
	//handlerCreate 创建订单
	handlerCreate = Handler(func(opt *Opt) (State, error) {
		message := fmt.Sprintf("正在处理创建订单逻辑，订单ID(%d),订单名称(%s) ... 处理完毕！", opt.OrderId, opt.OrderName)
		fmt.Println(message)
		if opt.HandlerSendSMS != nil {
			_ = opt.HandlerSendSMS("13937173036", "恭喜你预定成功了！")
		}
		return StatusReserved, nil
	})
	//handlerConfirm 确定订单
	handlerConfirm = Handler(func(opt *Opt) (State, error) {
		//业务逻辑
		return StatusConfirmed, nil
	})
	//handlerModify 修改订单
	handlerModify = Handler(func(opt *Opt) (State, error) {
		//业务逻辑
		return StatusReserved,nil
	})
	//handlerPay 支付订单
	handlerPay= Handler(func(opt *Opt) (State, error) {
		//业务逻辑
		return StatusLocked,nil
	})

)
