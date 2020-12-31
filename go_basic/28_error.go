package main

import (
	"fmt"
	"math"
)

type areaError struct {
	msg    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("error: %s", (*e).msg)
}

// 要求返回的是一个error类型的接口, 这个类型的接口就是要实现Error()方法且返回值为string
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{
			msg:    "radius is not positive.",
			radius: radius,
		}
	}
	return math.Pow(radius, 2) * math.Pi, nil
}

func main() {
	var radius float64 = -1
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		// 现在得到的err是一个接口, 通过断言得到具体的 接口实例结构
		if errDetail, ok := err.(*areaError); ok {
			//fmt.Printf("errDetail的类型: %T", errDetail) // *main.areaError-1
			fmt.Println((*errDetail).radius)
		}
		return
	}
	fmt.Printf("the area of circle with a radius of %.2f is %.2f\n", radius, area)
}
