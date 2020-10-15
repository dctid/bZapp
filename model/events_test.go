package model

import (
	"reflect"
	"testing"
)

func Test_addEventInOrder(t *testing.T) {
	type args struct {
		event  Event
		events []Event
	}
	tests := []struct {
		name string
		args args
		want []Event
	}{
		{
			name: "empty",
			args: args{
				event: Event{
					Title: "title",
					Day:   "day",
					Hour:  9,
					Min:   0,
					AmPm:  "AM",
				},
				events: []Event{},
			},
			want: []Event{
				{
					Title: "title",
					Day:   "day",
					Hour:  9,
					Min:   0,
					AmPm:  "AM",
				},
			},
		},
		{
			name: "after",
			args: args{
				event: Event{
					Title: "title",
					Day:   "day",
					Hour:  12,
					Min:   0,
					AmPm:  "PM",
				},
				events: []Event{
					{
						Title: "title",
						Day:   "day",
						Hour:  9,
						Min:   0,
						AmPm:  "AM",
					},
				},
			},
			want: []Event{
				{
					Title: "title",
					Day:   "day",
					Hour:  9,
					Min:   0,
					AmPm:  "AM",
				},
				{
					Title: "title",
					Day:   "day",
					Hour:  12,
					Min:   0,
					AmPm:  "PM",
				},
			},
		},
		{
			name: "before",
			args: args{
				event: Event{
					Title: "title",
					Day:   "day",
					Hour:  9,
					Min:   0,
					AmPm:  "AM",
				},
				events: []Event{
					{
						Title: "title",
						Day:   "day",
						Hour:  12,
						Min:   0,
						AmPm:  "PM",
					},
				},
			},
			want: []Event{
				{
					Title: "title",
					Day:   "day",
					Hour:  9,
					Min:   0,
					AmPm:  "AM",
				},
				{
					Title: "title",
					Day:   "day",
					Hour:  12,
					Min:   0,
					AmPm:  "PM",
				},
			},
		},
		{
			name: "middle",
			args: args{
				event: Event{
					Title: "title",
					Day:   "day",
					Hour:  10,
					Min:   30,
					AmPm:  "AM",
				},
				events: []Event{
					{
						Title: "title",
						Day:   "day",
						Hour:  9,
						Min:   0,
						AmPm:  "AM",
					},
					{
						Title: "title",
						Day:   "day",
						Hour:  12,
						Min:   0,
						AmPm:  "PM",
					},
				},
			},
			want: []Event{
				{
					Title: "title",
					Day:   "day",
					Hour:  9,
					Min:   0,
					AmPm:  "AM",
				},
				{
					Title: "title",
					Day:   "day",
					Hour:  10,
					Min:   30,
					AmPm:  "AM",
				},
				{
					Title: "title",
					Day:   "day",
					Hour:  12,
					Min:   0,
					AmPm:  "PM",
				},
			},
		},
		{
			name: "complex",
			args: args{
				event: Event{
					Title: "title",
					Day:   "day",
					Hour:  10,
					Min:   30,
					AmPm:  "AM",
				},
				events: []Event{
					{
						Title: "title",
						Day:   "day",
						Hour:  9,
						Min:   0,
						AmPm:  "AM",
					},
					{
						Title: "title",
						Day:   "day",
						Hour:  12,
						Min:   15,
						AmPm:  "PM",
					},
					{
						Title: "title",
						Day:   "day",
						Hour:  12,
						Min:   15,
						AmPm:  "PM",
					},
				},
			},
			want: []Event{
				{
					Title: "title",
					Day:   "day",
					Hour:  9,
					Min:   0,
					AmPm:  "AM",
				},
				{
					Title: "title",
					Day:   "day",
					Hour:  10,
					Min:   30,
					AmPm:  "AM",
				},
				{
					Title: "title",
					Day:   "day",
					Hour:  12,
					Min:   15,
					AmPm:  "PM",
				},
				{
					Title: "title",
					Day:   "day",
					Hour:  12,
					Min:   15,
					AmPm:  "PM",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddEventInOrder(tt.args.event, tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddEventInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_greaterThan(t *testing.T) {

	type args struct {
		other Event
	}
	tests := []struct {
		name  string
		event Event
		args  args
		want  bool
	}{
		{
			name: "first",
			event: Event{
				Title: "",
				Day:   "",
				Hour:  10,
				Min:   15,
				AmPm:  "AM",
			},
			args: args{
				other: Event{
					Title: "",
					Day:   "",
					Hour:  12,
					Min:   0,
					AmPm:  "PM",
				},
			},

			want: true,
		},
		{
			name: "second",
			event: Event{
				Title: "",
				Day:   "",
				Hour:  12,
				Min:   0,
				AmPm:  "PM",
			},
			args: args{
				other: Event{
					Title: "",
					Day:   "",
					Hour:  10,
					Min:   15,
					AmPm:  "AM",
				},
			},

			want: false,
		},
		{
			name: "another",
			event: Event{
				Title: "title",
				Day:   "day",
				Hour:  10,
				Min:   30,
				AmPm:  "AM",
			},
			args: args{
				other: Event{
					Title: "title",
					Day:   "day",
					Hour:  12,
					Min:   15,
					AmPm:  "PM",
				},
			},

			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.event.greaterThan(tt.args.other); got != tt.want {
				t.Errorf("greaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIndexToInsertAt(t *testing.T) {
	type args struct {
		event  Event
		events []Event
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "something",
			args: args{
				event: Event{
					Title: "title",
					Day:   "day",
					Hour:  10,
					Min:   30,
					AmPm:  "AM",
				},
				events: []Event{
					{
						Title: "title",
						Day:   "day",
						Hour:  9,
						Min:   0,
						AmPm:  "AM",
					},
					{
						Title: "title",
						Day:   "day",
						Hour:  12,
						Min:   15,
						AmPm:  "PM",
					},
					{
						Title: "title",
						Day:   "day",
						Hour:  12,
						Min:   15,
						AmPm:  "PM",
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndexToInsertAt(tt.args.event, tt.args.events); got != tt.want {
				t.Errorf("findIndexToInsertAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RemoveEvent(t *testing.T) {
	type args struct {
		id     string
		events []Event
	}
	tests := []struct {
		name string
		args args
		want []Event
	}{
		{
			name: "empty",
			args: args{
				id:     "ABC",
				events: []Event{},
			},
			want: []Event{},
		},
		{
			name: "not found",
			args: args{
				id:     "ABC",
				events: []Event{{Id: "DEF"}},
			},
			want: []Event{{Id: "DEF"}},
		},
		{
			name: "first",
			args: args{
				id:     "ABC",
				events: []Event{{Id: "ABC"}},
			},
			want: []Event{},
		},
		{
			name: "second",
			args: args{
				id:     "ABC",
				events: []Event{{Id: "DEF"}, {Id: "ABC"}},
			},
			want: []Event{{Id: "DEF"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeEvent(tt.args.id, tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
