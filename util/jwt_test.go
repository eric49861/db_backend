package util

import (
	"testing"
)

//
func TestGenerateTokenWithHS256(t *testing.T) {
	if _, err := GenerateTokenWithHS256("ERIC", 0); err != nil {
		t.Errorf("TestGenerateTokenWithHS256 errors")
	}
}

//
func TestParseToken(t *testing.T) {
	tokenString, _ := GenerateTokenWithHS256("ERIC", 0)

	id, err := ParseToken(tokenString)
	if err != nil {
		t.Errorf("%s\n", err)
	} else {
		if id != 0 {
			t.Errorf("解析失败")
		}
	}
}

//
func TestRefreshToken(t *testing.T) {
	tokenString, _ := GenerateTokenWithHS256("ERIC", 0)
	_, err := RefreshToken(tokenString)
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

//
func TestIsValidTokenString(t *testing.T) {
	tokenString, _ := GenerateTokenWithHS256("ERIC", 0)
	isValid := IsValidTokenString(tokenString)
	if !isValid {
		t.Errorf("token已过期，但是新生成的token不应该立即过期")
	}
}
