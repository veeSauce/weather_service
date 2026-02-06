// package main

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHandleFunc(t *testing.T) {

// 	w := httptest.NewRecorder()

// 	inputForm(w, nil)

// 	desiredCode := http.StatusOK
// 	if w.Code != desiredCode {
// 		t.Errorf("Expected status code %d, got %d", desiredCode, w.Code)
// 	}

// 	desiredBody := []byte("Hello, World!\n")
// 	if !bytes.Equal(desiredBody, w.Body.Bytes()) {
// 		t.Errorf("Expected body %s, got %s", desiredBody, w.Body.String())
// 	}

// }
