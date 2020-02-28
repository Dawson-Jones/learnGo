package main

import (
	"fmt"
	p "my_package/person"
	"my_package/pk1"
	"my_package/utils"
	"my_package/utils/time_utils"
) // 绝对路径, 从src下面开始导入

func main() {
	utils.Count()
	time_utils.PrintTime()
	pk1.MyTest()
	utils.Mytest2()
	fmt.Println("----------------")
	person := p.Person{
		Name: "dawson",
		Age:  23,
		Sex:  1,
	}
	fmt.Println(person)
	fmt.Println(person.Name, person.Age, person.Sex)

}
