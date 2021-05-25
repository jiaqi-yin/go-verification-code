package utils

import (
	"time"
)

func TimeDiffInSecBetweenNowAndTomorrow() (time.Duration, error) {
	location, err := time.LoadLocation("Australia/Melbourne")
	if err != nil {
		return 0, err
	}

	now := time.Now().In(location)
	tomorrow := now.AddDate(0, 0, 1)
	layout := "2006-01-02"
	tmrTime, err := time.ParseInLocation(layout, tomorrow.Format(layout), location)
	if err != nil {
		return 0, err
	}

	diff := tmrTime.Sub(now)
	return time.Duration(diff.Seconds()) * time.Second, nil
}
