package modal

import (
	"github.com/slack-go/slack"
	"reflect"
	"testing"
)

func TestBuildNewEventSectionBlock(t *testing.T) {
	type args struct {
		values map[string]map[string]slack.BlockAction
	}

	values := map[string]map[string]slack.BlockAction{
		AddEventTitleInputBlock : {AddEventTitleActionId: slack.BlockAction{Value: "title"}},
		AddEventDayInputBlock: {AddEventDayActionId: slack.BlockAction{SelectedOption: slack.OptionBlockObject{Value: TodayOptionValue}}},
		AddEventHoursInputBlock: {AddEventHoursActionId: slack.BlockAction{SelectedOption: slack.OptionBlockObject{Text: &slack.TextBlockObject{Text: "10 AM"}}}},
		AddEventMinsInputBlock: {AddEventMinsActionId: slack.BlockAction{SelectedOption: slack.OptionBlockObject{Text: &slack.TextBlockObject{Text: "15"}}}},
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 *slack.SectionBlock
	}{
		{name: "default", args: args{values: values} , want: "today",
			want1: EventSectionWithoutRemoveButton("title", "10 AM", "15"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BuildNewEventSectionBlock(tt.args.values)
			if got != tt.want {
				t.Errorf("BuildNewEventSectionBlock() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BuildNewEventSectionBlock() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}