package alert

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// Get text from all the nested elements, returns it as a string
func getText(n *html.Node) string {
	var ret []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			ret = append(ret, strings.TrimSpace(c.Data))
		case html.ElementNode:
			getText(c)
		}
	}
	return strings.Join(ret, " ")
}

// Looks through document to find nodes
func findNodes(body, tag, class string) ([]string, error) {
	// Source adapted from the html package's examples
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var ret []string
	var find func(n *html.Node)
	find = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
			for _, a := range n.Attr {
				if a.Key == "class" {
					classes := strings.Fields(a.Val)
					for _, c := range classes {
						if c == class {
							ret = append(ret, getText(n))
						}
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			find(c)
		}
	}
	find(doc)
	return ret, nil
}

func TestApp(t *testing.T) {
	app := App{}

	type checkFn func(r *http.Response, body string) error
	// We expect no alerts, error out if there are alerts
	hasNoAlerts := func() checkFn {
		return func(r *http.Response, body string) error {
			nodes, err := findNodes(body, "div", "alert")
			if err != nil {
				return fmt.Errorf("findNodes() err = %s", err)
			}
			if len(nodes) != 0 {
				return fmt.Errorf("len(alerts)=%d; want len(alerts)=0", len(nodes))
			}
			return nil
		}
	}
	// We expect an alert, error out if none are found
	hasAlert := func(msg string) checkFn {
		return func(r *http.Response, body string) error {
			nodes, err := findNodes(body, "div", "alert")
			if err != nil {
				return fmt.Errorf("findNodes() err = %s", err)
			}
			for _, node := range nodes {
				// If we found the alert we were looking for
				if node == msg {
					return nil
				}
			}
			return fmt.Errorf("missing alert: %q", msg)
		}
	}
	tests := []struct {
		method string
		path   string
		body   io.Reader
		checks []checkFn
	}{
		{http.MethodGet, "/", nil, []checkFn{hasNoAlerts()}},
		{http.MethodGet, "/alert", nil, []checkFn{hasAlert("Stuff went wrong!")}},
		{http.MethodGet, "/many", nil, []checkFn{hasAlert("Alert Number 1"), hasAlert(("Alert Number 2"))}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.method, tt.path), func(t *testing.T) {
			w := httptest.NewRecorder()
			r, err := http.NewRequest(tt.method, tt.path, tt.body)
			if err != nil {
				t.Fatalf("http.NewRequest() err = %s", err)
			}
			app.ServeHTTP(w, r)
			res := w.Result()
			defer res.Body.Close()
			var sb strings.Builder
			io.Copy(&sb, res.Body)
			for _, check := range tt.checks {
				if err := check(res, sb.String()); err != nil {
					t.Error(err)
				}
			}
		})
	}
}
