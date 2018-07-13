package portations

import (
	"bufio"
	//	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Portations struct {
	Storage *sync.Map
}

//create new sync.Map return Portations object
func New() *Portations {
	sm := sync.Map{}
	return &Portations{Storage: &sm}
}

//return number portaion info
func (ports *Portations) GetNumInfo(num string) (string, bool) {
	if s, ok := ports.Storage.Load(num); ok {
		return s.(string), ok
	}
	return "", false

}

//return number portaion info to channel
func (ports *Portations) GetNumInfoToChan(num string, ansChan chan string) {
	if s, ok := ports.Storage.Load(num); ok {
		//return s.(string), ok
		ansChan <- s.(string)
	} else {
		ansChan <- ""
	}
}

//change new sync Map load Port_all_full
//port_all_full
//Number,OwnerId,MNC,Route,RegionCode,PortDate,RowCount,NPId,DonorId,RangeHolderId,OldRoute,OldMNC,ProcessType
//9000000000,mMEGAFON,02,D2502,25,2016-09-07T12:00:17+03:00,7978178,1000000006441231,mTELE2,mTELE2,,,ShortTimePort
func (ports *Portations) LoadPortAllFull(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		log.Println("cant open file ", fname, err)
		return err
	}
	defer file.Close()
	sm := sync.Map{}
	scanner := bufio.NewScanner(file)
	//m := make(map[int]string)
	var s string
	var l []string
	for scanner.Scan() {
		s = scanner.Text()
		l = strings.Split(s, ",")
		if len(l) > 3 {
			//ports.Storage.Store(l[0], s)
			//sm.Store(l[0], parsePortAllFull(l))
			sm.Store(l[0], strings.Join([]string{l[1], l[3]}, ","))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	ports.Storage = &sm
	return nil
}

//change new sync Map load Port_all
//Port_all
//Number,OwnerId,MNC,Route,RegionCode,PortDate,RowCount
//9000000000,mMEGAFON,02,D2502,25,2016-09-07T12:00:17+03:00,7504433
func (ports *Portations) LoadPortAll(fname string) error {
	log.Println("LoadPortAll start file name=", fname)
	file, err := os.Open(fname)
	if err != nil {
		log.Println("cant open file ", fname, err)
		return err
	}
	defer file.Close()
	sm := sync.Map{}
	scanner := bufio.NewScanner(file)
	//m := make(map[int]string)
	var s string
	var l []string
	for scanner.Scan() {
		s = scanner.Text()
		l = strings.Split(s, ",")
		if len(l) > 3 {
			//ports.Storage.Store(l[0], s)
			//sm.Store(l[0], parsePortAllFull(l))
			sm.Store(l[0], strings.Join([]string{l[1], l[3]}, ","))
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	ports.Storage = &sm
	log.Println("LoadPortAll finished")
	return nil
}

//load PortIncrement.csv file change Store in ports
//NPId,Number,RecipientId,DonorId,RangeHolderId,OldRoute,NewRoute,OldMNC,NewMNC,RegionCode,PortDate,RowCount
//1000000015391205,9946666819,mBEELINE,mMTT,mMTT,D7742,D7799,42,99,77,2018-06-28T11:59:46+03:00,252
func (ports *Portations) LoadPortIncrement(fname string) error {
	//	NPId,            Number,    RecipientId,DonorId,RangeHolderId,OldRoute,NewRoute,OldMNC,NewMNC,RegionCode,PortDate,RowCount
	//1000000015391205,9946666822,mBEELINE,   mMTT,   mMTT,         D7742,   D7799,   42,    99,    77,        2018-06-28T12:00:18+03:00,
	//0                 1           2         3       4             5        6        7       8      9          10
	file, err := os.Open(fname)
	if err != nil {
		log.Println("cant open file ", fname, err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var s string
	var l []string
	for scanner.Scan() {
		s = scanner.Text()
		l = strings.Split(s, ",")
		if len(l) > 2 {
			//ports.Storage.Store(l[0], s)
			//sm.Store(l[0], parsePortAllFull(l))
			//log.Println(l)
			ports.Storage.Store(l[1], strings.Join([]string{l[2], l[6]}, ","))
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

//remove portations from store
func (ports *Portations) LoadReturnIncrement(fname string) error {
	log.Println("LoadReturnIncrement file name=", fname)
	file, err := os.Open(fname)
	if err != nil {
		log.Println("cant open file ", fname, err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var s string
	var l []string
	for scanner.Scan() {
		s = scanner.Text()
		l = strings.Split(s, ",")
		if len(l) > 2 {
			//ports.Storage.Store(l[0], s)
			//sm.Store(l[0], parsePortAllFull(l))
			//log.Println(l)
			ports.Storage.Delete(l[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	log.Println("LoadReturnIncrement finished")
	return nil
}
