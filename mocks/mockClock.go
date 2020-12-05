package mocks

import (
	"time"
)

type mockClock struct {
	t time.Time
}
func NewMockClock(setTimeTo string) mockClock {
	return mockClock{mockATime(setTimeTo)}
}

func (m mockClock) Now() time.Time {
	return m.t
}

func mockATime(timeToMock string) time.Time {
	parse, _ := time.Parse("2006-01-02 15:04:05", timeToMock)
	return parse
}


