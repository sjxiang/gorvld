package utils

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Fatal(v interface{}) {
	fmt.Printf("gorvld:\033[0;1;31m fatal:\033[0m %v\n", v)
	debug.PrintStack()  // 打印堆栈信息
	os.Exit(1)
}

func MustNo(err error) {
	if err != nil {
		Fatal(err)
	}
}


func Assert(condition bool) {
	if !condition {
		Fatal("assert failed")
	}
}

