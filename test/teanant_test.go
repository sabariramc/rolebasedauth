package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
	"sabariram.com/goserverbase/utils"
	"sabariram.com/rolebasedauth/pkg/app"
	"sabariram.com/rolebasedauth/pkg/constants"
	"sabariram.com/rolebasedauth/pkg/model"
)

func BenchmarkHandleListBook(b *testing.B) {
	srv, err := app.GetDefaultApp()
	if err != nil {
		b.Fatalf("Error creating server - %v", err)
	}
	req := httptest.NewRequest("GET", "/book", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	assert.Equal(b, w.Result().StatusCode, http.StatusOK)
}

func TestHandleCreateTenant(t *testing.T) {
	srv, err := app.GetDefaultApp()
	if err != nil {
		t.Fatalf("Error creating server - %v", err)
	}
	body := &model.CreateTenantDTO{
		Name:    "fasdfs",
		BaseURL: "https://a.b.com",
		Claims: []*model.CreateClaimDTO{
			{Claim: "fads.fads", Description: "fadsf"},
		},
		AuthenticationType: []*model.CreateAuthenticationDTO{
			{Type: constants.AuthTypeBasicAuth, Configuration: model.AuthConfiguration{
				"fasdf": "fasfd",
			}},
		},
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req := httptest.NewRequest("POST", "/tenant", &buf)
	req.Header.Set("x-api-key", utils.GetEnv("TEST_API_KEY", ""))
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	blob, _ := ioutil.ReadAll(w.Body)
	fmt.Println(string(blob))
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
}
