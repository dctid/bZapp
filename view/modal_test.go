package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"reflect"
	"testing"
)

func TestBuildNewEventSectionBlock(t *testing.T) {
	type args struct {
		index  int
		values map[string]map[string]slack.BlockAction
	}

	values := map[string]map[string]slack.BlockAction{
		fmt.Sprintf("%s-%d", AddEventTitleInputBlock, 1): {AddEventTitleActionId: slack.BlockAction{Value: "title"}},
		fmt.Sprintf("%s-%d", AddEventDayInputBlock, 1):   {AddEventDayActionId: slack.BlockAction{SelectedOption: slack.OptionBlockObject{Value: TodayOptionValue}}},
		fmt.Sprintf("%s-%d", AddEventHoursInputBlock, 1): {AddEventHoursActionId: slack.BlockAction{SelectedOption: slack.OptionBlockObject{Text: &slack.TextBlockObject{Text: "10 AM"}}}},
		fmt.Sprintf("%s-%d", AddEventMinsInputBlock, 1):  {AddEventMinsActionId: slack.BlockAction{SelectedOption: slack.OptionBlockObject{Text: &slack.TextBlockObject{Text: "15"}}}},
	}
	tests := []struct {
		name string
		args args
		want *model.Event
	}{
		{
			name: "default", args: args{index: 1, values: values},
			want: &model.Event{Id: "FakeHash", Title: "title", Day: TodayOptionValue, Hour: 10, Min: 15, AmPm: "AM"},
		},
	}
	for _, tt := range tests {
		model.Hash = func() string {
			return "FakeHash"
		}
		t.Run(tt.name, func(t *testing.T) {
			got := BuildNewEvent(tt.args.index, tt.args.values)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildNewEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

