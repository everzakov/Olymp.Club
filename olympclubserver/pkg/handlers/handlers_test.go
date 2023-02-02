package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type SomeHanler struct{}

func (s *SomeHanler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.String(), "/123") {
		s.WriteOk(w, r)
	} else {
		ans := make(map[string]interface{})
		ans["ups"] = ":("
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&ans)
	}
}

func (s *SomeHanler) WriteOk(w http.ResponseWriter, r *http.Request) {
	ans := make(map[string]interface{})
	ans["ok"] = "ok"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ans)
}

func TestXxx(t *testing.T) {
	h := &SomeHanler{}
	ts := httptest.NewServer(h)
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL + "/123")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	dynamic := make(map[string]interface{})
	json.Unmarshal([]byte(data), &dynamic)
	okStringInterface, ok := dynamic["ok"]
	if !ok {
		t.Error("No ok in response")
		return
	}
	okString, ok := okStringInterface.(string)
	if !ok {
		t.Error("ok is not a string")
		return
	}
	if okString != "ok" {
		t.Error("ok is not ok")
	}
}
