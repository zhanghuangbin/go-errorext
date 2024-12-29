# go-errorext

go-errorext提供了若干工具方法，用于获取更多的debug信息，方便定位问题。

## Import

```go

import (
	errorext "github.com/zhanghuangbin/go-errorext"
)

```

## Quick Start

结合go 1.20引入的链式错误api，逐层记录调用的位置。

```go
import (
    errorext "github.com/zhanghuangbin/go-errorext"
)

func error1() error {
    return errors.New(fmt.Sprintf("%v: init error", errorext.CurrentFrame()))
}


func error2() error {
    err1 := error1()
    if err1 != nil {
        return fmt.Errorf("%v: wrapper err1 \n%w", errorext.CurrentFrame(), err1)
    }
    return nil
}



func main()  {
    err2 := error2()
    if err2 != nil {
        err3 := fmt.Errorf("%v: wrapper err2 \n%w", errorext.CurrentFrame(), err2)
        if err3 != nil {
            fmt.Printf("%v", err3)
        }
    }
}

// Output:
//	main.go:31: wrapper err2
//	main.go:23: wrapper err1
//	main.go:17: init error
```
其他的输出格式，请详见文档。

## more

这个仓库，只是学习go的玩具，未经过严格测试，请勿用于生产环境