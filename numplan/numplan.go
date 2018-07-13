package numplan

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"strings"
	"sync"
	//	"time"
	"strconv"
)

type Numplan struct {
	Storage *sync.Map
}

//create new sync.Map return Numplan object
func New() *Numplan {
	sm := sync.Map{}
	return &Numplan{Storage: &sm}
}

//read file line by line and put key to hashmap
func (numplan *Numplan) LoadNumberingPlan(fname string) error {
	log.Println("LoadNumberingPlan start file name=", fname)
	file, err := os.Open(fname)
	if err != nil {
		log.Println("cant open file ", fname, err)
		return err
	}
	defer file.Close()
	sm := sync.Map{}
	scanner := bufio.NewScanner(file)
	var s string
	var l []string
	for scanner.Scan() {
		s = scanner.Text()
		l = strings.Split(s, ",")
		if len(l) > 2 {
			if _, err := strconv.Atoi(l[0]); err == nil {
				pfxList := range2prefix(l[0], l[1])
				for _, pfx := range pfxList {
					//log.Println(pfx)
					sm.Store(pfx, s)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	numplan.Storage = &sm
	log.Println("LoadNumberingPlan finished")
	return nil
}

func (numplan *Numplan) GetNumInfo(num string) (string, bool) {
	for i := len(num) - 1; i > 3; i-- {
		//nplan.GetNumInfo("9999812345")log.Println(num[:i])
		if s, ok := numplan.Storage.Load(num[:i]); ok {
			return s.(string), ok
		}
	}
	return "", false

}

func (numplan *Numplan) GetNumInfoToChan(num string, ansChan chan string) {
	if info, ok := numplan.GetNumInfo(num); ok {
		ansChan <- info
	} else {
		ansChan <- "not found"
	}
}
