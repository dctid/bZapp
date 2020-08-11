package modal

import (
	"github.com/slack-go/slack"
	"log"
	"reflect"
	"testing"
)

func TestAddNewEventToDay(t *testing.T) {
	type args struct {
		blocks   []slack.Block
		eventDay string
		newEvent *slack.SectionBlock
	}
	tests := []struct {
		name  string
		args  args
		want  []*slack.SectionBlock
		want1 []*slack.SectionBlock
	}{
		{
			name: "add to today when empty",
			args: args{
				blocks:   BuildEventBlocks(NoEventYetSection, NoEventYetSection),
				eventDay: TodayOptionValue,
				newEvent: EventSectionWithoutRemoveButton("retro", "1 PM", "44"),
			},
			want:  []*slack.SectionBlock{EventSectionWithoutRemoveButton("retro", "1 PM", "44")},
			want1: NoEventYetSection,
		},
		{
			name: "add to tomorrow when empty",
			args: args{
				blocks:   BuildEventBlocks(NoEventYetSection, NoEventYetSection),
				eventDay: TomorrowOptionValue,
				newEvent: EventSectionWithoutRemoveButton("retroed", "3 PM", "22"),
			},
			want:  NoEventYetSection,
			want1: []*slack.SectionBlock{EventSectionWithoutRemoveButton("retroed", "3 PM", "22")},
		},
		{
			name: "add to today when not empty",
			args: args{
				blocks:   BuildEventBlocks([]*slack.SectionBlock{EventSectionWithRemoveButton(TodayOptionValue, 0, "standup", "9 AM", "07")}, NoEventYetSection),
				eventDay: TodayOptionValue,
				newEvent: EventSectionWithoutRemoveButton("retro", "1 PM", "44"),
			},
			want: []*slack.SectionBlock{
				EventSectionWithoutRemoveButton("standup", "9 AM", "07"),
				EventSectionWithoutRemoveButton("retro", "1 PM", "44"),
			},
			want1: NoEventYetSection,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddNewEventToDay(tt.args.blocks, tt.args.eventDay, tt.args.newEvent)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNewEventToDay() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AddNewEventToDay() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExtractEvents(t *testing.T) {
	type args struct {
		blocks []slack.Block
	}
	tests := []struct {
		name  string
		args  args
		want  []*slack.SectionBlock
		want1 []*slack.SectionBlock
	}{
		{name: "empty",
			args:  args{blocks: NewSummaryModal(NoEventYetSection, NoEventYetSection).Blocks.BlockSet},
			want:  []*slack.SectionBlock{},
			want1: []*slack.SectionBlock{},
		},
		{name: "one each",
			args: args{
				blocks: NewSummaryModal(
					[]*slack.SectionBlock{EventSectionWithoutRemoveButton("Standup", "9 AM", "15")},
					[]*slack.SectionBlock{EventSectionWithoutRemoveButton("Standdown", "10 AM", "30")},
				).Blocks.BlockSet,
			},
			want:  []*slack.SectionBlock{EventSectionWithoutRemoveButton("Standup", "9 AM", "15")},
			want1: []*slack.SectionBlock{EventSectionWithoutRemoveButton("Standdown", "10 AM", "30")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExtractEvents(tt.args.blocks)
			log.Printf("got: %v, got1: %v", len(got), len(got1))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractEvents() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ExtractEvents() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReplaceEmptyEventsWithNoEventsYet(t *testing.T) {
	type args struct {
		todaysSectionBlocks    []*slack.SectionBlock
		tomorrowsSectionBlocks []*slack.SectionBlock
	}
	tests := []struct {
		name  string
		args  args
		want  []*slack.SectionBlock
		want1 []*slack.SectionBlock
	}{
		{name: "both empty", args: args{
			todaysSectionBlocks:    []*slack.SectionBlock{},
			tomorrowsSectionBlocks: []*slack.SectionBlock{},
		},
			want:  []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoEventsText, false, false), nil, nil)},
			want1: []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoEventsText, false, false), nil, nil)},
		},
		{name: "today empty", args: args{
			todaysSectionBlocks:    []*slack.SectionBlock{},
			tomorrowsSectionBlocks: []*slack.SectionBlock{EventSectionWithoutRemoveButton("title", "10 AM", "15")},
		},
			want:  []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoEventsText, false, false), nil, nil)},
			want1: []*slack.SectionBlock{EventSectionWithoutRemoveButton("title", "10 AM", "15")},
		},
		{name: "tomorrow empty", args: args{
			todaysSectionBlocks:    []*slack.SectionBlock{EventSectionWithoutRemoveButton("title", "10 AM", "15")},
			tomorrowsSectionBlocks: []*slack.SectionBlock{},
		},
			want:  []*slack.SectionBlock{EventSectionWithoutRemoveButton( "title", "10 AM", "15")},
			want1: []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoEventsText, false, false), nil, nil)},
		},
		{name: "none empty", args: args{
			todaysSectionBlocks:    []*slack.SectionBlock{EventSectionWithoutRemoveButton( "title", "10 AM", "15")},
			tomorrowsSectionBlocks: []*slack.SectionBlock{EventSectionWithoutRemoveButton( "title2", "11 AM", "30")},
		},
			want:  []*slack.SectionBlock{EventSectionWithoutRemoveButton( "title", "10 AM", "15")},
			want1: []*slack.SectionBlock{EventSectionWithoutRemoveButton( "title2", "11 AM", "30")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ReplaceEmptyEventsWithNoEventsYet(tt.args.todaysSectionBlocks, tt.args.tomorrowsSectionBlocks)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReplaceEmptyEventsWithNoEventsYet() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReplaceEmptyEventsWithNoEventsYet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRemoveEvent(t *testing.T) {
	type args struct {
		blocks   []slack.Block
		actionId string
	}
	tests := []struct {
		name  string
		args  args
		want  []*slack.SectionBlock
		want1 []*slack.SectionBlock
	}{
		{
			name: "remove first today",
			args: args{
				blocks:   BuildEventBlocks([]*slack.SectionBlock{EventSectionWithRemoveButton(TodayOptionValue, 0, "standup", "9 AM", "07")}, NoEventYetSection),
				actionId: "remove_today_0",
			},
			want:  NoEventYetSection,
			want1: NoEventYetSection,
		},
		{
			name: "remove first tomorrow",
			args: args{
				blocks:   BuildEventBlocks(NoEventYetSection, []*slack.SectionBlock{EventSectionWithRemoveButton(TomorrowOptionValue, 0, "standup", "9 AM", "07")}),
				actionId: "remove_tomorrow_0",
			},
			want:  NoEventYetSection,
			want1: NoEventYetSection,
		},
		{
			name: "remove second today",
			args: args{
				blocks:   BuildEventBlocks([]*slack.SectionBlock{EventSectionWithRemoveButton(TodayOptionValue, 0, "standup", "9 AM", "07"), EventSectionWithRemoveButton(TodayOptionValue, 1, "retro", "3 PM", "00")}, NoEventYetSection),
				actionId: "remove_today_1",
			},
			want:  []*slack.SectionBlock{EventSectionWithRemoveButton(TodayOptionValue, 0, "standup", "9 AM", "07")},
			want1: NoEventYetSection,
		},
		{
			name: "remove first today with second",
			args: args{
				blocks:   BuildEventBlocks([]*slack.SectionBlock{EventSectionWithRemoveButton(TodayOptionValue, 0, "standup", "9 AM", "07"), EventSectionWithRemoveButton(TodayOptionValue, 1, "retro", "3 PM", "00")}, NoEventYetSection),
				actionId: "remove_today_0",
			},
			want:  []*slack.SectionBlock{EventSectionWithRemoveButton(TodayOptionValue, 1, "retro", "3 PM", "00")},
			want1: NoEventYetSection,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RemoveEvent(tt.args.blocks, tt.args.actionId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveEvent() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RemoveEvent() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestConvertToEventsWithRemoveButton(t *testing.T) {
	type args struct {
		todayEvents    []*slack.SectionBlock
		tomorrowEvents []*slack.SectionBlock
	}
	tests := []struct {
		name  string
		args  args
		want  []*slack.SectionBlock
		want1 []*slack.SectionBlock
	}{
		{
			name: "empty",
			args: args{
				todayEvents:    []*slack.SectionBlock{},
				tomorrowEvents: []*slack.SectionBlock{},
			},
			want:  []*slack.SectionBlock{},
			want1: []*slack.SectionBlock{},
		},
		{
			name: "today has one",
			args: args{
				todayEvents:    []*slack.SectionBlock{EventSectionWithoutRemoveButton("today event", "9 AM", "12")},
				tomorrowEvents: []*slack.SectionBlock{},
			},
			want:  []*slack.SectionBlock{EventSectionWithRemoveButton(TodayOptionValue, 0, "today event", "9 AM", "12")},
			want1: []*slack.SectionBlock{},
		},
		{
			name: "tomorrow has two",
			args: args{
				todayEvents: []*slack.SectionBlock{},
				tomorrowEvents: []*slack.SectionBlock{
					EventSectionWithoutRemoveButton("tomorrow event", "9 AM", "12"),
					EventSectionWithoutRemoveButton("tomorrow event 2", "11 AM", "01"),
				},
			},
			want: []*slack.SectionBlock{},
			want1: []*slack.SectionBlock{
				EventSectionWithRemoveButton(TomorrowOptionValue, 0, "tomorrow event", "9 AM", "12"),
				EventSectionWithRemoveButton(TomorrowOptionValue, 1, "tomorrow event 2", "11 AM", "01"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConvertToEventsWithRemoveButton(tt.args.todayEvents, tt.args.tomorrowEvents)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToEventsWithRemoveButton() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ConvertToEventsWithRemoveButton() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
