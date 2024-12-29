package errorext_test

import (
	"errors"
	"fmt"
	"github.com/zhanghuangbin/go-errorext"
)

func ExampleCurrentFrame() {
	err1 := errors.New(fmt.Sprintf("%v: init error", errorext.CurrentFrame()))
	err2 := fmt.Errorf("%v: wrapper err1 \n%w", errorext.CurrentFrame(), err1)
	err3 := fmt.Errorf("%v: wrapper err2 \n%w", errorext.CurrentFrame(), err2)

	fmt.Printf("%v\n", err3)
	// FIXME 输出看起来是一样的，但是样例测试不通过，注释的格式有误？空格或者回车问题？
	// Output:
	//
	// example_errorext_test.go:12: wrapper err2
	//
	// example_errorext_test.go:11: wrapper err1
	//
	// example_errorext_test.go:10: init error
	//
}
