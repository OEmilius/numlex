package portations

import (
	"fmt"
	"time"
)

//func ExampleLoadPortAllFull() {
//	//fname := "loadfile.go"
//	ports := New()
//	fname := "Port_All_Full_201806280000_1569.csv"
//	t1 := time.Now()
//	err := ports.LoadPortAllFull(fname)
//	fmt.Println(err)
//	fmt.Println(time.Now().Sub(t1))
//	fmt.Println(ports.Storage.Load("9999999997"))
//	//Output: mBEELINE,99,D7799,77 true
//}

func ExampleGetNumInfo() {
	ports := New()
	fname := "Port_All_Full_201806280000_1569.csv"
	_ = ports.LoadPortAllFull(fname)
	t1 := time.Now()
	for i := 0; i < 500000; i++ {
		ports.GetNumInfo("9000000012")
		ports.GetNumInfo("9999999991")
	}
	fmt.Println(time.Now().Sub(t1))
	//Output: 2
}

//func ExampleLoadPortAll() {
//	//fname := "loadfile.go"
//	ports := New()
//	fname := "Port_All_201807060000_1689.csv"
//	err := ports.LoadPortAll(fname)
//	fmt.Println(err)
//	t1 := time.Now()
//	fmt.Println(time.Now().Sub(t1))
//	fmt.Println(ports.Storage.Load("9999999997"))
//	//Output: mBEELINE,D7799 true
//}

//func ExampleLoadPortIncrement() {
//	ports := New()
//	fname := "Port_Increment_201806281400_20179.csv"
//	err := ports.LoadPortIncrement(fname)
//	fmt.Println(err)
//	t1 := time.Now()
//	fmt.Println(time.Now().Sub(t1))
//	fmt.Println(ports.GetNumInfo("9946666819"))
//	//Output:
//}

//func ExampleLoadPortIncrement2() {
//	ports := New()
//	fname := "Port_All_201807060000_1689.csv"
//	err := ports.LoadPortAll(fname)
//	fmt.Println(ports.GetNumInfo("9946666819"))
//	fname = "Port_Increment_201806281400_20179.csv"
//	err = ports.LoadPortIncrement(fname)
//	t1 := time.Now()
//	fmt.Println(time.Now().Sub(t1))
//	fmt.Println(err)
//	fmt.Println(ports.GetNumInfo("9946666819"))
//	//Output: 2
//}

//func ExampleLoadReturnIncrement() {
//	ports := New()
//	fname := "Port_All_201807060000_1689.csv"
//	err := ports.LoadPortAll(fname)
//	fmt.Println(ports.GetNumInfo("9165776004"))
//	fname = "Return_Increment_201807061800_20277.csv"
//	err = ports.LoadReturnIncrement(fname)
//	t1 := time.Now()
//	fmt.Println(time.Now().Sub(t1))
//	fmt.Println(err)
//	fmt.Println(ports.GetNumInfo("9165776004"))
//	//Output: 2
//}
