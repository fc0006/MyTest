package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var ttt bool
	ttt = true
	fmt.Printf("%v\n", ttt)
	//var utcLoc *time.Location   .In(utcLoc)
	//根据时间戳转换时间
	var res time.Time //时间格式
	res = time.Unix(1665373149, 0)

	//把时间转换成时间戳
	fmt.Println(res, "时间戳转时间")
	//
	fmt.Println(time.Now().Unix()) //秒
	fmt.Println(time.Now().UnixMicro())
	fmt.Println(time.Now().UnixMilli()) //毫秒

}
