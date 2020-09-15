package stripe_test

import (
	"fmt"
	"net/http"
	"testing"

	stripe "github.com/ritabc/twg/stripe/v1"
)

func TestApp(t *testing.T) {
	client, mux, teardown := stripe.TestClient(t)
	defer teardown()

	mux.HandleFunc("/v1/charges", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id":"ch_1HRnR72eZvKYlo2C0MH1drpQ", "amount":2000, "description":"Charge for demo puposes.","status":"succeeded"}`)
	})

	charge, err := client.Charge(123, "doesnt_matter", "something else")
	if err != nil {
		t.Errorf("Charge() err = %s; want nil", err)
	}
	if charge.ID != "ch_1HRnR72eZvKYlo2C0MH1drpQ" {
		t.Errorf("Charge() id = %s; want %s", charge.ID, "ch_1HRnR72eZvKYlo2C0MH1drpQ")
	}
}
