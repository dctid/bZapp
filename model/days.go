package model

import "time"

const layout = "2006-01-02"

type clock interface {
	Now() time.Time
}
type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

var (
	Clock clock = RealClock{}
)

type TimeApi struct {
	Today      string
	NextBizDay string
	IsFriday   bool
}

func Days() TimeApi {
	loc, _ := time.LoadLocation("America/Detroit")
	nowEasternTimezone := Clock.Now().In(loc)
	tomorrow := nowEasternTimezone.Add(time.Hour * 24)
	isFriday := false
	if tomorrow.Weekday() == time.Saturday {
		tomorrow = tomorrow.Add(time.Hour * 48)
		isFriday = true
	}

	return TimeApi{
		Today:      nowEasternTimezone.Format(layout),
		NextBizDay: tomorrow.Format(layout),
		IsFriday:   isFriday,
	}
}
