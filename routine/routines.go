package routine

var panicHandler func(interface{})

//SetPanicHandler 注册panic 处理函数
func SetPanicHandler(handler func(interface{})) {
	panicHandler = handler
}

//SaftyRun Goroutines 安全执行，增加defer处理panic
func SaftyRun(method func() (interface{}, error)) (resp interface{}, err error) {
	defer func() {
		if err := recover(); err != nil {
			if panicHandler != nil {
				panicHandler(err)
			}
		}
	}()
	return method()
}
