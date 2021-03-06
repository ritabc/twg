package demo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ritabc/twg/stripe/v1"
)

type App struct {
	Stripe *stripe.Client
}

func (a *App) Run() {}

func TestApp(t *testing.T) {
	client, mux, teardown := stripe.TestClient(t)

	defer teardown()

	mux.HandleFunc("/v1/charges", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id":"ch_1HRnR72eZvKYlo2C0MH1drpQ","amount":2000,"description":Charge for demo purposes.","status":"succeeded"}`)
	})

	// Now inject client into your app and run your tests - they will use your local tests server using this mux
	app := App{
		Stripe: client,
	}
	app.Run()
	// ...

}
