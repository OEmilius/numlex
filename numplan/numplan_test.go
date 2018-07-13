package numplan

import (
	"fmt"
	//	"sync"
	//"testing"
	"time"
)

//func ExampleLoadFile() {
//	//fname := "loadfile.go"
//	fname := "Numbering_plan_201806270000_1681.csv"
//	sm := sync.Map{}
//	t1 := time.Now()
//	err := LoadFile(fname, &sm)
//	fmt.Println(err)

//	//hm.String()
//	//fmt.Println("total keys in hm=", hm.Len())
//	fmt.Println(time.Now().Sub(t1))
//	//GetInfo(&sm)
//	fmt.Println("ddddddd")
//	//Output: 1
//}

//func ExampleLoadFileAndCheck() {
//	fname := "Numbering_plan_201806270000_1681.csv"
//	sm := sync.Map{}
//	err := LoadFile(fname, &sm)
//	fmt.Println(err)
//	t1 := time.Now()
//	for i := 0; i < 500000; i++ {
//		pfxInfo, _ := sm.Load("90000")
//		pfxInfo, _ = sm.Load("99998")
//		_ = pfxInfo
//	}
//	fmt.Println("seconds=", time.Now().Sub(t1).Seconds())
//	fmt.Println(sm.Load("99998"))
//	//Output: 2
//}

func ExampleGetPhoneInfo() {
	fname := "Numbering_plan_201806270000_1681.csv"
	//sm := sync.Map{}
	nplan := New()
	//err := LoadFile(fname, &sm)
	err := nplan.LoadNumberingPlan(fname)
	fmt.Println(err)
	t1 := time.Now()
	fmt.Println(nplan.GetNumInfo("9999812345"))
	for i := 0; i < 500000; i++ {
		nplan.GetNumInfo("9999812345")
		nplan.GetNumInfo("9022199999")
		//		GetPhoneInfo(&sm, "9999812345")
		//		GetPhoneInfo(&sm, "9022199999")
	}

	fmt.Println(time.Now().Sub(t1))
	//Output: 2
}
