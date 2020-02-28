package about_time

import (
	"fmt"
	"math/rand"
	"time"
)

func TimeTemplate() {
	// 模板时间要是 (2006 1 2 15 4 5)
	s := "1996-01-15 13:00:00"
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	// time类型的数据
	year, month, day := t.Date()
	fmt.Println(year, month, day)
	hour, minute, sec := t.Clock()
	fmt.Println(hour, minute, sec)

	// 睡眠 3 秒
	//time.Sleep(time.Duration(3))
	time.Sleep(time.Second * 3)
	fmt.Println("-----------")

	// 随机睡眠 1-10 秒
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println("-----------", randNum)
}
