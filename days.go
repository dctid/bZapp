package bZapp

import "time"

const layout = "2006-01-02"

type clock interface {
	Now() time.Time
}
type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

var (
	Clock  clock = RealClock{}
	loc, _       = time.LoadLocation("Us/Eastern")
)

type TimeApi struct {
	Today      string
	NextBizDay string
}

func Days() TimeApi {
	nowEasternTimezone := Clock.Now().In(loc)
	tomorrow := nowEasternTimezone.Add(time.Hour * 24)
	if tomorrow.Weekday() == time.Saturday {
		tomorrow = tomorrow.Add(time.Hour * 48)
	}

	return TimeApi{
		Today:      nowEasternTimezone.Format(layout),
		NextBizDay: tomorrow.Format(layout),
	}
}
