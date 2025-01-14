// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLibrary_Token_Getters(t *testing.T) {
	// setup tests
	tests := []struct {
		token *Token
		want  *Token
	}{
		{
			token: testToken(),
			want:  testToken(),
		},
		{
			token: new(Token),
			want:  new(Token),
		},
	}

	// run tests
	for _, test := range tests {
		if test.token.GetToken() != test.want.GetToken() {
			t.Errorf("GetToken is %v, want %v", test.token.GetToken(), test.want.GetToken())
		}
	}
}

func TestLibrary_Token_Setters(t *testing.T) {
	// setup types
	var l *Token

	// setup tests
	tests := []struct {
		token *Token
		want  *Token
	}{
		{
			token: testToken(),
			want:  testToken(),
		},
		{
			token: l,
			want:  new(Token),
		},
	}

	// run tests
	for _, test := range tests {
		test.token.SetToken(test.want.GetToken())

		if test.token.GetToken() != test.want.GetToken() {
			t.Errorf("SetToken is %v, want %v", test.token.GetToken(), test.want.GetToken())
		}
	}
}

func TestToken_String(t *testing.T) {
	// setup types
	l := testToken()

	want := fmt.Sprintf(`{
  Token: %s,
}`,
		l.GetToken(),
	)

	// run test
	got := l.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

// testToken is a test helper function to create a Token
// type with all fields set to a fake value.
func testToken() *Token {
	l := new(Token)

	l.SetToken("superSecretToken")

	return l
}
