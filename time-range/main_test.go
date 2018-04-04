package time_range

import (
	"testing"
	"time"
	"reflect"
	"bytes"
	"fmt"
	"strings"
)


var today = time.Now()

// Test given examples
func TestMinusPartialIntersection(t *testing.T) {

	range1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	range1End := getTimeFromTodayUsingHoursMinutesSeconds(12, 0, 0)
	range2Start := getTimeFromTodayUsingHoursMinutesSeconds(11, 30, 0)
	range2End := getTimeFromTodayUsingHoursMinutesSeconds(14, 0, 0)

	minuendRange := Timerange{range1Start, range1End}
	subtrahendRange := Timerange{range2Start, range2End}

	var minuendRanges []Timerange
	var subtrahendRanges []Timerange

	minuendRanges = append(minuendRanges, minuendRange)
	subtrahendRanges = append(subtrahendRanges, subtrahendRange)

	result := SubtractTimeranges(minuendRanges, subtrahendRanges)

	expectedRangeStart := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	expectedRangeEnd := getTimeFromTodayUsingHoursMinutesSeconds(11, 30, 0)

	expectedRange := Timerange{expectedRangeStart, expectedRangeEnd}

	var expectedRanges []Timerange

	expectedRanges = append(expectedRanges, expectedRange)

	if !reflect.DeepEqual(expectedRanges, result) {
		t.Errorf("Expected: %s, Actual: %s", timerangesToString(expectedRanges), timerangesToString(result))
	}
}

func TestMinusSameRange(t *testing.T) {

	t.Errorf("Not implemented")
}

func TestMinusNonIntersectingRange(t *testing.T) {

	t.Errorf("Not implemented")
}

func TestManyRangesMinusOneRange(t *testing.T) {

	t.Errorf("Not implemented")
}

func TestManyRangesMinusManyRanges(t *testing.T) {

	t.Errorf("Not implemented")
}

func getTimeFromTodayUsingHoursMinutesSeconds(hours int, minutes int, seconds int) time.Time {
	return time.Date(today.Year(), today.Month(), today.Day(), hours, minutes, seconds, 0, time.UTC)
}

func timerangesToString(timeranges []Timerange) string {
	var buffer bytes.Buffer

	for _, v := range timeranges {
		buffer.WriteString(fmt.Sprintf("(start: %s, end: %s),", v.start.String(), v.end.String()))
	}

	return strings.TrimRight(buffer.String(), ",")
}
