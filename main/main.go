package main

import (
	"fmt"
	"log"

	"github.com/OEmilius/numlex"
	"github.com/OEmilius/numlex/webserver"
)

func main() {
	log.Println("start")
	serv := webserver.NewServer(":8080")
	webserver.CDMN = numlex.New()
	err := webserver.CDMN.NPlan.LoadNumberingPlan("../numplan/Numbering_plan_201806270000_1681.csv")
	if err != nil {
		fmt.Println(err)
	}
	err = webserver.CDMN.Ports.LoadPortAll("../portations/Port_All_201807060000_1689.csv")
	if err != nil {
		fmt.Println(err)
	}
	go serv.StartWebServer()
	fmt.Println(webserver.CDMN.GetRNStep("9000000030"))
	fmt.Println("for exit pres any key")
	fmt.Scanln()
}
