package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type StateReq struct {
	OP        int
	Page      int
	PageSize  int
	StartTime time.Time
	EndTime   time.Time
	Data      interface{}
}

type Task struct {
	OperationID string
	Operation   string
	SubType     string
	State       string
	Target      *[]string `gorm:"-"`
	MemberID    int
	Total       int
	Succeed     int
	Failed      int
	Ongoing     int
}

func main() {
		req := StateReq{
		OP: 1,
		Data: Task{
			OperationID: "1234",
			Operation:   "CacheClean",
			SubType:     "domain",

			// 指针也能正确的 Marshal
			Target:      &[]string{"1234", "5678"},
			MemberID:    58182367,
		},
	}

	reqData, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(reqData))

	//stru := LogMsgT{}
	//json.Unmarshal(nsqLog, &stru)
	//fmt.Println(stru)
}
