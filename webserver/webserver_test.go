package webserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OEmilius/numlex"
)

func TestServerhomeHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveHome)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `example usage http://localhost:8080/?num=9000000030`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \r\ngot %v \r\nwant %v",
			rr.Body.String(), expected)
	}
}

func TestServerhomeHandler2(t *testing.T) {
	//serv := NewServer(":8080")
	CDMN = numlex.New()
	err := CDMN.NPlan.LoadNumberingPlan("../numplan/Numbering_plan_201806270000_1681.csv")
	if err != nil {
		fmt.Println(err)
	}
	err = CDMN.Ports.LoadPortAll("../portations/Port_All_201807060000_1689.csv")
	if err != nil {
		fmt.Println(err)
	}
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/?num=9000000030", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveHome)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `mBEELINE,D2599`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \r\ngot %v \r\nwant %v",
			rr.Body.String(), expected)
	}
}
