package model

import (
	"github.com/dctid/bZapp/mocks"
	"reflect"
	"testing"
)

func TestDays(t *testing.T) {
	tests := []struct {
		name           string
		wantToday      string
		wantNextBizDay string
		timeToTest     string
	}{
		{
			name:           "test a friday",
			wantToday:      "2020-12-04",
			wantNextBizDay: "2020-12-07",
			timeToTest:     "2020-12-04 09:01:29",
		},
		{
			name:           "test a not friday",
			wantToday:      "2020-12-02",
			wantNextBizDay: "2020-12-03",
			timeToTest:     "2020-12-02 09:08:29",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clock = mocks.NewMockClock(tt.timeToTest)
			got := Days()
			if !reflect.DeepEqual(got.Today, tt.wantToday) {
				t.Errorf("Days().Today = %v, want %v", got.Today, tt.wantToday)
			}
			if !reflect.DeepEqual(got.NextBizDay, tt.wantNextBizDay) {
				t.Errorf("Days().NextBizDay = %v, want %v", got.NextBizDay, tt.wantNextBizDay)
			}
		})
	}
}
