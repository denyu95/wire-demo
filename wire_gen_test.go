package main

import (
	"fmt"
	"testing"
)

type TestClient struct {
}

func NewTestClient() *TestClient {
	return &TestClient{}
}

func (m *TestClient) GetById(id string) string {
	return fmt.Sprintf("TestClient: some data %v", id)
}

func TestGetData(t *testing.T) {
	ds := NewTestClient()
	app := NewApp(ds)
	app.GetData("123")
}
