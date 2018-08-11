package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "io/ioutil"
)

func TestRouter(t *testing.T) {
  r := customRouter()
  // method, endpoint, body
  mockServer := httptest.NewServer(r)
  resp, err := http.Get(mockServer.URL + "/hello")
  if err!=nil {
    t.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    t.Errorf("Handler returned incorrect status code. Expected: %v Received: %v", resp.StatusCode, http.StatusOK)
  }

  defer resp.Body.Close()
  b, err := ioutil.ReadAll(resp.Body)
  if err!=nil {
    t.Fatal(err)
  }

  expected := `Hello world`
  actual := string(b)
  if actual != expected {
    t.Errorf("Handler returned incorrect body. Expected: %v Received: %v", expected, actual)
  }
}
