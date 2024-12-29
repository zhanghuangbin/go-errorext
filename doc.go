// Package errorext  是一个用于练手的工具，未经过严格的测试，请勿用于生产环境。
//
// errorext提供了若干工具函数，用于获取更多的debug信息，方便定位问题
//
// quick start:
//
// 假设在main.go文件中有如下代码：
//
//	err1 := errors.New(fmt.Sprintf("%v: init error", CurrentFrame()))
//	err2 := fmt.Errorf("%v: wrapper err1 \n%w", CurrentFrame(), err1)
//	err3 := fmt.Errorf("%v: wrapper err2 \n%w", CurrentFrame(), err2)
//	fmt.Printf("%v", err3)
//
// Output:
//
// main.go:11: wrapper err2
// main.go:10: wrapper err1
// main.go:9: init error
package errorext
