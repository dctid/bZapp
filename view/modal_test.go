package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"reflect"
	"strings"
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
		want model.Event
	}{
		{name: "default", args: args{index: 1, values: values},
			want: model.Event{Id: "FakeHash", Title: "title", Day: TodayOptionValue, Hour: 10, Min: 15, AmPm: "AM"},
		},
	}
	for _, tt := range tests {
		model.Hash = func() string {
			return "FakeHash"
		}
		t.Run(tt.name, func(t *testing.T) {
			got := BuildNewEvent(tt.args.index, tt.args.values)
			if got != tt.want {
				t.Errorf("BuildNewEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractModel(t *testing.T) {
	type args struct {
		blocks []slack.Block
	}
	tests := []struct {
		name string
		args args
		want model.Model
	}{
		{name: "empty",
			args: args{blocks: NewSummaryModal(model.Model{}).Blocks.BlockSet},
			want: model.Model{
				Events: model.Events{
					TodaysEvents:    []model.Event{},
					TomorrowsEvents: []model.Event{},
				},
				Goals: model.Goals{
					"Customer Questions?": {},
					"Learnings":           {},
					"Other":               {},
					"Questions?":          {},
					"Team Needs":          {},
				},
			},
		},
		{name: "one each",
			args: args{
				blocks: NewSummaryModal(
					model.Model{
						Index: 0,
						Events: model.Events{
							TodaysEvents: []model.Event{{
								Id:    "fake event id 1",
								Title: "Standup",
								Day:   TodayOptionValue,
								Hour:  9,
								Min:   15,
								AmPm:  "AM",
							},
							},
							TomorrowsEvents: []model.Event{{
								Id:    "fake event id 2",
								Title: "Standdown",
								Day:   TomorrowOptionValue,
								Hour:  10,
								Min:   30,
								AmPm:  "AM",
							},
							},
						},
						Goals: nil,
					},
				).Blocks.BlockSet,
			},
			want: model.Model{
				Events: model.Events{
					TodaysEvents: []model.Event{
						{
							Id:    "fake event id 1",
							Title: "Standup",
							Day:   "today",
							Hour:  9,
							Min:   15,
							AmPm:  "AM",
						}},
					TomorrowsEvents: []model.Event{
						{
							Id:    "fake event id 2",
							Title: "Standdown",
							Day:   "tomorrow",
							Hour:  10,
							Min:   30,
							AmPm:  "AM",
						},
					},
				},
				Goals: model.Goals{
					"Customer Questions?": {},
					"Learnings":           {},
					"Other":               {},
					"Questions?":          {},
					"Team Needs":          {},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractModel(tt.args.blocks)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractModel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToEventsWithRemoveButton(t *testing.T) {
	type args struct {
		events model.Events
	}
	tests := []struct {
		name  string
		args  args
		want  []slack.Block
		want1 []slack.Block
	}{
		{
			name: "empty",
			args: args{
				events: model.Events{},
			},
			want:  NoEventYetSection,
			want1: NoEventYetSection,
		},
		{
			name: "today has one",
			args: args{
				events: model.Events{
					TodaysEvents: []model.Event{{Id: "Today Event", Title: "today event", Day: TodayOptionValue, Hour: 9, Min: 12, AmPm: "AM"}},
				},
			},
			want:  []slack.Block{EventSectionWithRemoveButton(TodayOptionValue, 0, "Today Event", "today event", "9 AM", "12")},
			want1: NoEventYetSection,
		},
		{
			name: "tomorrow has two",
			args: args{
				events: model.Events{
					TomorrowsEvents: []model.Event{
						{Id: "Event 1", Title: "tomorrow event", Day: TomorrowOptionValue, Hour: 9, Min: 12, AmPm: "AM"},
						{Id: "Event 2", Title: "tomorrow event 2", Day: TomorrowOptionValue, Hour: 11, Min: 1, AmPm: "AM"},
					},
				},
			},
			want: NoEventYetSection,
			want1: []slack.Block{
				EventSectionWithRemoveButton(TomorrowOptionValue, 0, "Event 1", "tomorrow event", "9 AM", "12"),
				EventSectionWithRemoveButton(TomorrowOptionValue, 1, "Event 2", "tomorrow event 2", "11 AM", "01"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConvertToEventsBlocks(true, tt.args.events)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToEventsWithRemoveButton() got = %v\n, want %v\n", got, tt.want)

			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ConvertToEventsWithRemoveButton() \ngot1 = %v\n, want1 %v\n", got1, tt.want1)
			}
		})
	}
}

func TestConvertToEventsWithoutRemoveButton(t *testing.T) {
	type args struct {
		events model.Events
	}
	tests := []struct {
		name  string
		args  args
		want  []slack.Block
		want1 []slack.Block
	}{
		{
			name: "empty",
			args: args{
				events: model.Events{},
			},
			want:  NoEventYetSection,
			want1: NoEventYetSection,
		},
		{
			name: "today has one",
			args: args{
				events: model.Events{
					TodaysEvents:    []model.Event{{Id: "todays id", Title: "today event", Day: TodayOptionValue, Hour: 9, Min: 12, AmPm: "AM"}},
					TomorrowsEvents: []model.Event{},
				},
			},
			want:  []slack.Block{EventSectionWithoutRemoveButton("todays id", "today event", "9 AM", "12")},
			want1: NoEventYetSection,
		},
		{
			name: "tomorrow has two",
			args: args{
				events: model.Events{
					TodaysEvents: []model.Event{},
					TomorrowsEvents: []model.Event{
						{Id: "tomorrows 1", Title: "tomorrow event", Day: TomorrowOptionValue, Hour: 9, Min: 12, AmPm: "AM"},
						{Id: "tomorrows 2", Title: "tomorrow event 2", Day: TomorrowOptionValue, Hour: 11, Min: 1, AmPm: "AM"},
					},
				},
			},
			want: NoEventYetSection,
			want1: []slack.Block{
				EventSectionWithoutRemoveButton("tomorrows 1", "tomorrow event", "9 AM", "12"),
				EventSectionWithoutRemoveButton("tomorrows 2", "tomorrow event 2", "11 AM", "01"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConvertToEventsBlocks(false, tt.args.events)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToEventsBlocks() got = %v\n, want %v\n", got, tt.want)

			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ConvertToEventsBlocks() \ngot1 = %v\n, want1 %v\n", got1, tt.want1)
			}
		})
	}
}

func EventSectionWithRemoveButton(day string, index int, id string, title string, hour string, mins string) *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		slack.NewAccessory(
			slack.NewButtonBlockElement(
				RemoveEventActionId,
				fmt.Sprintf("remove_%s_%s", day, id),
				slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
			),
		),
		slack.SectionBlockOptionBlockID(id),
	)
}

func EventSectionWithoutRemoveButton(id string, title string, hour string, mins string) *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		nil,
		slack.SectionBlockOptionBlockID(id),
	)
}

func TestExtractInputIndex(t *testing.T) {
	type args struct {
		blocks []slack.Block
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero index",
			args: args{blocks: []slack.Block{
				slack.NewDividerBlock(),
				&slack.InputBlock{Type: slack.MBTInput, BlockID: "block_id-0"},
			}},
			want: 0,
		},
		{
			name: "one index",
			args: args{blocks: []slack.Block{
				slack.NewDividerBlock(),
				&slack.InputBlock{Type: slack.MBTInput, BlockID: "block_id-1"},
			}},
			want: 1,
		},
		{
			name: "missing index",
			args: args{blocks: []slack.Block{
				slack.NewDividerBlock(),
				&slack.InputBlock{Type: slack.MBTInput, BlockID: "block_id"},
			}},
			want: 0,
		},
		{
			name: "missing input",
			args: args{blocks: []slack.Block{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractInputIndex(tt.args.blocks); got != tt.want {
				t.Errorf("ExtractInputIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func eventSectionWithoutRemoveButton(id string, title string, hour string, mins string) *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		nil,
		slack.SectionBlockOptionBlockID(id),
	)
}
