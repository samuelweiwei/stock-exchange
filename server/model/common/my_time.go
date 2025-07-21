package common

import (
	"database/sql/driver"
	"time"
)

type MyTime time.Time

func NewMyTime(t time.Time) *MyTime {
	return (*MyTime)(&t)
}

func (t *MyTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	var (
		n time.Time
	)
	n, err = time.Parse("\"2006-01-02 15:04:05\"", string(data))
	*t = MyTime(n)
	return err
}
func (t *MyTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*t).Format("\"2006-01-02 15:04:05\"")), nil
}

func (t *MyTime) Value() (driver.Value, error) {
	return []byte(time.Time(*t).Format(TimeFormat)), nil
}

var (
	TimeFormat = "2006-01-02 15:04:05"
)
