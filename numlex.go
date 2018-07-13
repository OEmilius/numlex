package numlex

import (
	//	"log"
	"github.com/OEmilius/numlex/numplan"
	"github.com/OEmilius/numlex/portations"
)

//central database of migrated numbers
type CDMN struct {
	NPlan *numplan.Numplan
	Ports *portations.Portations
}

func New() *CDMN {
	n := numplan.New()
	p := portations.New()
	return &CDMN{n, p}
}

//return information about phone number from portation or from number plan
func (self *CDMN) GetRNStep(num string) string {
	var rn = ""
	var ok bool
	if rn, ok = self.Ports.GetNumInfo(num); ok {
		return rn
	} else if rn, ok = self.NPlan.GetNumInfo(num); ok {
		return rn
	}
	return "not found anywhere"
}
