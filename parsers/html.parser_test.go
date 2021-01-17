package parsers

import (
	"testing"

	"golang.org/x/net/html"
)

func TestHTMLProvider_pullHref(t *testing.T) {
	type args struct {
		node *html.Node
	}
	tests := []struct {
		name string
		h    *HTMLProvider
		args args
		want string
	}{
		{
			name: "Get href from an html node",
			h:    &HTMLProvider{},
			args: args{
				node: &html.Node{
					Attr: []html.Attribute{
						{
							Key: "src",
							Val: "some other value",
						},
						{
							Key: "href",
							Val: "https://jamescrocker.invalid/some/url",
						},
					},
				},
			},
			want: "https://jamescrocker.invalid/some/url",
		},
		{
			name: "Invalid tag returns empty string",
			h:    &HTMLProvider{},
			args: args{
				node: &html.Node{
					Attr: []html.Attribute{
						{
							Key: "src",
							Val: "some other value",
						},
					},
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTMLProvider{}
			if got := h.pullHref(tt.args.node); got != tt.want {
				t.Errorf("HTMLProvider.pullHref() = %v, want %v", got, tt.want)
			}
		})
	}
}
