package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input http.Header
		want  string
		err   error
	}{
		{input: http.Header{}, want: "", err: errors.New("no authorization header included")},
		{input: http.Header{"Authorization": []string{"ApiKey 01234"}}, want: "01234", err: nil},
		{input: http.Header{"Authorization": []string{"ApiKey"}}, want: "", err: errors.New("malformed authorization header")},
		{input: http.Header{"Authorization": []string{"i"}}, want: "", err: errors.New("malformed authorization header")},
	}

	for i, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, got)
		}
		if !reflect.DeepEqual(tc.err, err) {
			t.Fatalf("test: %d: expected: %v, got: %v", i+1, tc.err, err)
		}
	}
}
