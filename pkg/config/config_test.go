package config

import(
	"testing"
)

func TestGetRouterAddr(t *testing.T) {
	addr := GetRouterAddr()
	if addr != ":8080" {
		t.Errorf("return value is not 8080")
	}
}

func TestGetConnectionToken(t *testing.T) {
	token := GetConnectionToken()
	if token != "miraiketai2020:miraiketai2020@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local" {
		t.Errorf("return value is not expect value")
	}
}
