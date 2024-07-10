package utilities

import (
	"fmt"
	"strings"
	"time"
)

func NewDayOfWeekType(name string) time.Weekday {
	for i := time.Monday; i <= time.Saturday; i++ {
		if strings.Contains(strings.ToLower(i.String()), strings.ToLower(name)) {
			return i
		}
	}

	// default
	return time.Sunday
}

func FormatDay() {
	// CreateExcel()
	// ReadExcel()
	fmt.Println(NewDayOfWeekType("sun").String())

	d1 := time.Now()
	d2 := d1.Add(24 * time.Hour)
	d3 := d1.Add(48 * time.Hour)
	d4 := time.Date(d3.Year(), d3.Month(), d3.Day(), 15, 10, 59, 0, d3.Location())

	fmt.Println("d1", d1, "\nd2", d2, "\nd3", d3, "\nd4", d4)
}

func StringToDate(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

func ToTimestamp(timeInMilis int64) time.Time {
	return time.UnixMilli(timeInMilis)
}

func TimestampToDateFormat(timestamp int64) string {
	return time.UnixMilli(timestamp).Format("2006-01-02 15:04:05")
}
