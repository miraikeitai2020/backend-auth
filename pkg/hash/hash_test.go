package hash

import(
	"testing"
)

func TestCreateHashString(t *testing.T) {
	result := CreateHashString("pass")
	if result == "pass" {
		t.Errorf("password is not hashed.")
	}
}
