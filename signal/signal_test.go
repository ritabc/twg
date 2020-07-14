package signal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		// Fail or FailNow?
		// FailNow if we can't make the request
		t.Fatalf("http.NewRequest err = %s", err)
	}
	Handler(w, r)
	resp := w.Result()
	if resp.StatusCode != 200 {
		// Fail or FailNow?
		// If we get non 200, the rest of the test isn't super useful
		t.Fatalf("Handler() status = %d; want %d", resp.StatusCode, 200)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		// Fail or FailNow?
		// Even if we get the wrong header, we propbably still want to continue
		t.Errorf("Handler() Content-Type = %q; want %q", contentType, "application/json")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Fail or FailNow?
		// If we can't read the body, the remaining tests don't make sense
		t.Fatalf("ioutil.ReadAll(resp.Body) err = %s", err)
	}

	var p Person
	json.Unmarshal(data, &p)
	if err != nil {
		// Fail or FailNow?
		// Since if we can't unmarshal, the remaining tests don't make sense
		t.Fatalf("json.Unmarshal(resp.Body) err = %s", err)
	}

	if p.Age != 30 {
		// Fail or FailNow?
		// Even if the age is wrong, we should still check other attributes
		t.Errorf("person.Age = %d; want %d", p.Age, 30)
	}

	if p.Name != "Bob Jones" {
		// Fail or FailNow ; same as age check
		t.Errorf("person.Name = %s; want %s", p.Name, "Bob Jones")
	}
}
