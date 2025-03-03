package domain

import (
	"testing"
	"time"
)

func TestTodo_updateStatus(t *testing.T) {
	type fields struct {
		Id          string
		Title       string
		Description string
		Status      Status
		LastUpdate  time.Time
	}
	tests := []struct {
		name           string
		fields         fields
		expectedStatus Status
	}{
		{name: "test Created to In progress", fields: fields{Status: Created}, expectedStatus: InProgress},
		{name: "test In progress to Done", fields: fields{Status: InProgress}, expectedStatus: Done},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todo := &Todo{
				Id:          tt.fields.Id,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Status:      tt.fields.Status,
				LastUpdate:  tt.fields.LastUpdate,
			}
			todo.UpdateStatus()
		})
	}
}

func Test_from(t *testing.T) {
	type args struct {
		title       string
		description string
	}
	tests := []struct {
		name string
		args args
		want Todo
	}{
		{name: "checking the values should be there", args: args{title: "demo", description: "testing"}, want: Todo{Status: Created, Title: "demo", Description: "testing", LastUpdate: time.Now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TodoFrom(tt.args.title, tt.args.description)

			now := time.Now()

			if tt.want.Status != got.Status {
				t.Errorf("TodoFrom() = %v, want %v", got, tt.want)
			}
			if tt.want.Title != got.Title {
				t.Errorf("TodoFrom() = %v, want %v", got, tt.want)
			}
			if tt.want.Description != got.Description {
				t.Errorf("TodoFrom() = %v, want %v", got, tt.want)
			}
			if now.Sub(got.LastUpdate) > time.Second {
				t.Errorf("TodoFrom() = %v, want %v", got, tt.want)
			}

		})
	}
}
