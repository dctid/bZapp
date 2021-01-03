package view

import (
	"encoding/json"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEditGoalsModal(t *testing.T) {

	type args struct {
		model *model.Model
		metadata *model.Metadata
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				model: &model.Model{Index: 1},
				metadata: &model.Metadata{ChannelId: "fake id"},
			},
			want: test.ReadFile(t, "view/edit_goals_modal.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEditGoalsModal(tt.args.model, tt.args.metadata); !assert.EqualValues(t, format.PrettyJsonNoError(marshalNoError(got)), tt.want) {
				t.Errorf("NewEditGoalsModal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func marshalNoError(thing interface{}) string {
	marshal, _ := json.Marshal(thing)
	return string(marshal)
}

