package modal

import (
	"github.com/slack-go/slack"
	"log"
	"reflect"
	"testing"
)

func TestAddNewEventToCorrectDay(t *testing.T) {
	type args struct {
		eventDay               string
		todaysSectionBlocks    []*slack.SectionBlock
		newEvent               *slack.SectionBlock
		tomorrowsSectionBlocks []*slack.SectionBlock
	}
	tests := []struct {
		name  string
		args  args
		want  []*slack.SectionBlock
		want1 []*slack.SectionBlock
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddNewEventToCorrectDay(tt.args.eventDay, tt.args.todaysSectionBlocks, tt.args.newEvent, tt.args.tomorrowsSectionBlocks)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNewEventToCorrectDay() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AddNewEventToCorrectDay() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
		// TODO: Add test cases.
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


func TodaySection() []*slack.SectionBlock {
	return []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "9:15 Standup", false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement("", "remove_today_1", slack.NewTextBlockObject("plain_text", "Remove", true, false)))),
}}

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
			args:  struct{ blocks []slack.Block }{blocks: NewSummaryModal(NoEventYetSection, NoEventYetSection).Blocks.BlockSet},
			want: []*slack.SectionBlock{},
			want1: []*slack.SectionBlock{},
		},
		{name: "one each",
			args: struct {
				blocks []slack.Block
			}{
				blocks: NewSummaryModal(TodaySection(), TodaySection()).Blocks.BlockSet,
			},
			want: TodaySection(),
			want1: TodaySection(),
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
		// TODO: Add test cases.
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
