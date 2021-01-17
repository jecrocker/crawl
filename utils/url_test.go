package utils

import (
	"net/url"
	"reflect"
	"testing"
)

func TestValidateURL(t *testing.T) {
	type args struct {
		check   string
		baseURL *url.URL
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 *url.URL
	}{
		{
			name: "Absolute URL that is valid",
			args: args{
				// Using a .invalid domain as the IETF have ratified as per RFC2606
				// Other tlds that could have been used are
				// .example
				// .localhost
				// .test
				check: "https://jamescrocker.invalid/something/something",
				baseURL: &url.URL{
					Scheme: "https",
					Host:   "jamescrocker.invalid",
					Path:   "/",
				},
			},
			want: true,
			want1: &url.URL{
				Scheme: "https",
				Host:   "jamescrocker.invalid",
				Path:   "/something/something",
			},
		},
		{
			name: "Relative URL that is valid",
			args: args{
				check: "/something/something",
				baseURL: &url.URL{
					Scheme: "https",
					Host:   "jamescrocker.invalid",
					Path:   "/",
				},
			},
			want: true,
			want1: &url.URL{
				Scheme: "https",
				Host:   "jamescrocker.invalid",
				Path:   "/something/something",
			},
		},
		{
			name: "Absolute URL that is invalid",
			args: args{
				check: "https://some-other-person.invalid/something/something",
				baseURL: &url.URL{
					Scheme: "https",
					Host:   "jamescrocker.invalid",
					Path:   "/",
				},
			},
			want: false,
			want1: &url.URL{
				Scheme: "https",
				Host:   "some-other-person.invalid",
				Path:   "/something/something",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateURL(tt.args.check, tt.args.baseURL)
			if got != tt.want {
				t.Errorf("ValidateURL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
