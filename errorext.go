package errorext

import (
	"errors"
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// Frame 栈的一帧信息
type Frame struct {
	// 当前程序计数器(program counter)
	pc uintptr
	// 调用函数的名称
	Name string
	// 源代码的文件
	File string
	// 源代码的行数
	Line int
}

// CurrentFrame 获取当前调用栈的帧
//
// 该方法返回的帧信息，该方法的caller调用currentFrame的位置信息。
func CurrentFrame() Frame {
	var pcs [1]uintptr
	n := runtime.Callers(2, pcs[:])
	if n < 1 {
		panic(errors.New("can not parse frame through CurrentFrame fn"))
	}

	pc := pcs[0]
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	file, line := fn.FileLine(pc - 1)

	return Frame{
		pc:   pc,
		Name: name,
		File: file,
		Line: line,
	}
}

// Format 根据fmt.Formatter接口格式化
//
//	%s    源代码文件
//	%d    源代码的函数
//	%n    调用函数名
//	%v    %s:%d#%n 的简写
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//	%+s   即 函数名\n完整文件名
//
//	%+v   %+s:%d 的简写
func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case s.Flag('+'):
			io.WriteString(s, f.Name)
			io.WriteString(s, "\n")
			io.WriteString(s, f.File)
		default:
			io.WriteString(s, path.Base(f.File))
		}
	case 'd':
		io.WriteString(s, strconv.Itoa(f.Line))
	case 'n':
		io.WriteString(s, simplifyFuncName(f.Name))
	case 'v':
		f.Format(s, 's')
		io.WriteString(s, ":")
		f.Format(s, 'd')
	}
}

// 简化文件路径,只保留文件名
func simplifyFuncName(path string) string {
	i := strings.LastIndex(path, "/")
	path = path[i+1:]
	i = strings.Index(path, ".")
	return path[i+1:]
}
