package view

import(
	"testing"
)

func TestMakeErrorResponse(t *testing.T) {
	result := MakeErrorResponse(400, "hoge")
	if result.Code != 400 {
		t.Errorf("code value is not `400`")
	}
	if result.Message != "Bad Request" {
		t.Errorf("message value is not `Bad Request`")
	}
	if result.Description != "hoge" {
		t.Errorf("description value is not `hoge`")
	}
	result = MakeErrorResponse(403, "hoge")
	if result.Code != 403 {
		t.Errorf("code value is not `403`")
	}
	if result.Message != "Forbidden" {
		t.Errorf("message value is not `Forbidden`")
	}
	if result.Description != "hoge" {
		t.Errorf("description value is not `hoge`")
	}
	result = MakeErrorResponse(500, "hoge")
	if result.Code != 500 {
		t.Errorf("code value is not `id`")
	}
	if result.Message != "Internal Server Error" {
		t.Errorf("message value is not `Internal Server Error`")
	}
	if result.Description != "hoge" {
		t.Errorf("description value is not `hoge`")
	}
}
