package numlex

import (
	"fmt"
	"math/rand"
	//"runtime"
	"strconv"
	"time"
)

//func ExampleCDMNGetRN() {
//	//runtime.GOMAXPROCS(4)
//	cdmn := New()
//	err := cdmn.NPlan.LoadNumberingPlan("numplan/Numbering_plan_201806270000_1681.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = cdmn.Ports.LoadPortAll("portations/Port_All_201807060000_1689.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	//fmt.Println(cdmn)
//	//fmt.Println(cdmn.NPlan.GetNumInfo("9999812345"))
//	rn := cdmn.GetRN("9000000030")
//	fmt.Println(rn)
//	t1 := time.Now()
//	for i := 0; i < 1000000; i++ {
//		num := rand.Intn(999999999)
//		//fmt.Println("rand int=", num)
//		rn := cdmn.GetRN("9" + strconv.Itoa(num))
//		_ = rn
//	}
//	fmt.Println(time.Now().Sub(t1))
//	//Output: 1
//}

//func ExampleCDMNGetRNStep() {
//	//runtime.GOMAXPROCS(4)
//	cdmn := New()
//	err := cdmn.NPlan.LoadNumberingPlan("numplan/Numbering_plan_201806270000_1681.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = cdmn.Ports.LoadPortAll("portations/Port_All_201807060000_1689.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	//fmt.Println(cdmn)
//	//fmt.Println(cdmn.NPlan.GetNumInfo("9999812345"))
//	rn := cdmn.GetRNStep("9000000030")
//	fmt.Println(rn)
//	t1 := time.Now()
//	for i := 0; i < 1000000; i++ {
//		num := rand.Intn(999999999)
//		//fmt.Println("rand int=", num)
//		rn := cdmn.GetRNStep("9" + strconv.Itoa(num))
//		_ = rn
//	}
//	fmt.Println(time.Now().Sub(t1))
//	//Output: 1
//}

//func ExampleCDMNGetRNcheck2GetNumInfo() {
//	cdmn := New()
//	err := cdmn.NPlan.LoadNumberingPlan("numplan/Numbering_plan_201806270000_1681.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = cdmn.Ports.LoadPortAll("portations/Port_All_201807060000_1689.csv")
//	if err != nil {
//		fmt.Println(err)
//	}
//	t1 := time.Now()
//	for i := 0; i < 1000000; i++ {
//		num := rand.Intn(999999999)
//		cdmn.NPlan.GetNumInfo("9" + strconv.Itoa(num))
//	}
//	fmt.Println("nplan get=", time.Now().Sub(t1))
//	t1 = time.Now()
//	for i := 0; i < 1000000; i++ {
//		num := rand.Intn(999999999)
//		cdmn.Ports.GetNumInfo("9" + strconv.Itoa(num))
//	}
//	fmt.Println("ports get=", time.Now().Sub(t1))
//	//Output: 2
//}

func ExampleCDMNGetRNStep() {
	//runtime.GOMAXPROCS(4)
	cdmn := New()
	err := cdmn.NPlan.LoadNumberingPlan("numplan/Numbering_plan_201806270000_1681.csv")
	if err != nil {
		fmt.Println(err)
	}
	err = cdmn.Ports.LoadPortAll("portations/Port_All_201807060000_1689.csv")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(cdmn)
	//fmt.Println(cdmn.NPlan.GetNumInfo("9999812345"))
	rn := cdmn.GetRNStep("9000000030")
	fmt.Println(rn)
	go func() {
		time.Sleep(10 * time.Millisecond)
		//err = cdmn.Ports.LoadPortAll("portations/Port_All_201807060000_1689.csv")
		err = cdmn.Ports.LoadReturnIncrement("portations/Return_Increment_201807061800_20277.csv")
		err = cdmn.Ports.LoadPortIncrement("portations/Port_Increment_201806281400_20179.csv")
	}()
	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		num := rand.Intn(999999999)
		//fmt.Println("rand int=", num)
		rn := cdmn.GetRNStep("9" + strconv.Itoa(num))
		_ = rn
	}
	//	err = cdmn.Ports.LoadPortAll("portations/Port_All_201807060000_1689.csv")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	fmt.Println(time.Now().Sub(t1))
	//Output: 1
}
