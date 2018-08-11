package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandler(t *testing.T) {
  // method, endpoint, body
  req, err := http.NewRequest("GET", "", nil)
  if err!=nil {
    t.Fatal(err)
  }

  recorder := httptest.NewRecorder()
  hf := http.HandlerFunc(handler)
  hf.ServeHTTP(recorder, req)

  if status := recorder.Code; status != http.StatusOK {
    t.Errorf("Handler returned incorrect status code. Expected: %v Received: %v", status, http.StatusOK)
  }

  expected := `Hello world`
  actual := recorder.Body.String()
  if actual != expected {
    t.Errorf("Handler returned incorrect body. Expected: %v Received: %v", expected, actual)
  }
}
