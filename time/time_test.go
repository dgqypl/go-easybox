package time

import (
	easyboxt "github.com/dgqypl/go-easybox/testing"
	"testing"
	"time"
)

var timeString = "2022-08-09 14:38:29"
var seconds int64 = 1660027109
var milliseconds int64 = 1660027109123

type secondsFormatCase struct {
	In  int64
	Out string
}

func TestSecondsFormat(t *testing.T) {
	easyboxt.FuncAssertion(t, &secondsFormatCase{seconds, timeString}, SecondsFormat)
}

type millisecondsFormatCase struct {
	In  int64
	Out string
}

func TestMillisecondsFormat(t *testing.T) {
	easyboxt.FuncAssertion(t, &millisecondsFormatCase{milliseconds, timeString}, MillisecondsFormat)
}

type secondsToTimeCase struct {
	In  int64
	Out time.Time
}

func TestSecondsToTime(t *testing.T) {
	_time := time.Unix(seconds, 0)
	easyboxt.FuncAssertion(t, &secondsToTimeCase{seconds, _time}, SecondsToTime)
}

type millisecondsToTimeCase struct {
	In  int64
	Out time.Time
}

func TestMillisecondsToTime(t *testing.T) {
	_time := time.Unix(0, milliseconds*1e6)
	easyboxt.FuncAssertion(t, &millisecondsToTimeCase{milliseconds, _time}, MillisecondsToTime)
}
