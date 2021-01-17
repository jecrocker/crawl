package utils

import (
	"sync"
	"testing"
)

func TestTracker_HasItem(t *testing.T) {
	type fields struct {
		mu   *sync.Mutex
		data map[string]string
	}
	type args struct {
		check string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Testing an item that does exist",
			fields: fields{
				mu: &sync.Mutex{},
				data: map[string]string{
					"testing": "testing",
				},
			},
			args: args{
				check: "testing",
			},
			want: true,
		},
		{
			name: "Testing an item that doesn't exist",
			fields: fields{
				mu: &sync.Mutex{},
				data: map[string]string{
					"testing": "testing",
				},
			},
			args: args{
				check: "testing123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracker := &Tracker{
				mu:   tt.fields.mu,
				data: tt.fields.data,
			}
			if got := tracker.HasItem(tt.args.check); got != tt.want {
				t.Errorf("Tracker.HasItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTracker_AddURL(t *testing.T) {
	type fields struct {
		mu     *sync.Mutex
		data   map[string]string
		queued int
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add a URL",
			fields: fields{
				mu: &sync.Mutex{},
				data: map[string]string{},
				queued: 0,
			},
			args: args{
				url: "testURL",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tracker := &Tracker{
				mu:     tt.fields.mu,
				data:   tt.fields.data,
				queued: tt.fields.queued,
			}
			tracker.AddURL(tt.args.url)
			if got := tracker.HasItem(tt.args.url); !got {
				t.Errorf("Tracker.AddURL() failed to add url to data")
			}
		})
	}
}
