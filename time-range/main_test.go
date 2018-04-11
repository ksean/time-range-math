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
	range1End := getTimeFromTodayUsingHoursMinutesSeconds(10, 0, 0)
	range2Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	range2End := getTimeFromTodayUsingHoursMinutesSeconds(9, 30, 0)

	minuendRange := Timerange{range1Start, range1End}
	subtrahendRange := Timerange{range2Start, range2End}

	var minuendRanges []Timerange
	var subtrahendRanges []Timerange

	minuendRanges = append(minuendRanges, minuendRange)
	subtrahendRanges = append(subtrahendRanges, subtrahendRange)

	result := SubtractTimeranges(minuendRanges, subtrahendRanges)

	expectedRangeStart := getTimeFromTodayUsingHoursMinutesSeconds(9, 30, 0)
	expectedRangeEnd := getTimeFromTodayUsingHoursMinutesSeconds(10, 0, 0)

	expectedRange := Timerange{expectedRangeStart, expectedRangeEnd}

	var expectedRanges []Timerange

	expectedRanges = append(expectedRanges, expectedRange)

	if !reflect.DeepEqual(expectedRanges, result) {
		t.Errorf("Expected: %s, Actual: %s", timerangesToString(expectedRanges), timerangesToString(result))
	}
}

func TestMinusSameRange(t *testing.T) {

	range1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	range1End := getTimeFromTodayUsingHoursMinutesSeconds(10, 0, 0)
	range2Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	range2End := getTimeFromTodayUsingHoursMinutesSeconds(10, 0, 0)

	minuendRange := Timerange{range1Start, range1End}
	subtrahendRange := Timerange{range2Start, range2End}

	var minuendRanges []Timerange
	var subtrahendRanges []Timerange

	minuendRanges = append(minuendRanges, minuendRange)
	subtrahendRanges = append(subtrahendRanges, subtrahendRange)

	result := SubtractTimeranges(minuendRanges, subtrahendRanges)


	var expectedRanges []Timerange

	if !reflect.DeepEqual(expectedRanges, result) {
		t.Errorf("Expected: %s, Actual: %s", timerangesToString(expectedRanges), timerangesToString(result))
	}
}

func TestMinusNonIntersectingRange(t *testing.T) {

	range1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	range1End := getTimeFromTodayUsingHoursMinutesSeconds(9, 30, 0)
	range2Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 30, 0)
	range2End := getTimeFromTodayUsingHoursMinutesSeconds(15, 0, 0)

	minuendRange := Timerange{range1Start, range1End}
	subtrahendRange := Timerange{range2Start, range2End}

	var minuendRanges []Timerange
	var subtrahendRanges []Timerange

	minuendRanges = append(minuendRanges, minuendRange)
	subtrahendRanges = append(subtrahendRanges, subtrahendRange)

	result := SubtractTimeranges(minuendRanges, subtrahendRanges)

	expectedRangeStart := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	expectedRangeEnd := getTimeFromTodayUsingHoursMinutesSeconds(9, 30, 0)

	expectedRange := Timerange{expectedRangeStart, expectedRangeEnd}

	var expectedRanges []Timerange

	expectedRanges = append(expectedRanges, expectedRange)

	if !reflect.DeepEqual(expectedRanges, result) {
		t.Errorf("Expected: %s, Actual: %s", timerangesToString(expectedRanges), timerangesToString(result))
	}
}

func TestManyRangesMinusOneRange(t *testing.T) {

	minuend1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	minuend1End := getTimeFromTodayUsingHoursMinutesSeconds(9, 30, 0)
	minuend2Start := getTimeFromTodayUsingHoursMinutesSeconds(10, 0, 0)
	minuend2End := getTimeFromTodayUsingHoursMinutesSeconds(10, 30, 0)

	subtrahendStart := getTimeFromTodayUsingHoursMinutesSeconds(9, 15, 0)
	subtrahendEnd := getTimeFromTodayUsingHoursMinutesSeconds(10, 15, 0)

	minuend1Range := Timerange{minuend1Start, minuend1End}
	minuend2Range := Timerange{minuend2Start, minuend2End}
	subtrahendRange := Timerange{subtrahendStart, subtrahendEnd}

	var minuendRanges []Timerange
	var subtrahendRanges []Timerange

	minuendRanges = append(minuendRanges, minuend1Range, minuend2Range)
	subtrahendRanges = append(subtrahendRanges, subtrahendRange)

	result := SubtractTimeranges(minuendRanges, subtrahendRanges)

	expectedRange1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	expectedRange1End := getTimeFromTodayUsingHoursMinutesSeconds(9, 15, 0)
	expectedRange2Start := getTimeFromTodayUsingHoursMinutesSeconds(10, 15, 0)
	expectedRange2End := getTimeFromTodayUsingHoursMinutesSeconds(10, 30, 0)

	expected1Range := Timerange{expectedRange1Start, expectedRange1End}
	expected2Range := Timerange{expectedRange2Start, expectedRange2End}

	var expectedRanges []Timerange

	expectedRanges = append(expectedRanges, expected1Range, expected2Range)

	if !reflect.DeepEqual(expectedRanges, result) {
		t.Errorf("Expected: %s, Actual: %s", timerangesToString(expectedRanges), timerangesToString(result))
	}
}

func TestManyRangesMinusManyRanges(t *testing.T) {

	minuend1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	minuend1End := getTimeFromTodayUsingHoursMinutesSeconds(11, 0, 0)
	minuend2Start := getTimeFromTodayUsingHoursMinutesSeconds(13, 0, 0)
	minuend2End := getTimeFromTodayUsingHoursMinutesSeconds(15, 0, 0)

	subtrahend1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 0, 0)
	subtrahend1End := getTimeFromTodayUsingHoursMinutesSeconds(9, 15, 0)
	subtrahend2Start := getTimeFromTodayUsingHoursMinutesSeconds(10, 0, 0)
	subtrahend2End := getTimeFromTodayUsingHoursMinutesSeconds(10, 15, 0)
	subtrahend3Start := getTimeFromTodayUsingHoursMinutesSeconds(12, 30, 0)
	subtrahend3End := getTimeFromTodayUsingHoursMinutesSeconds(16, 0, 0)

	minuend1Range := Timerange{minuend1Start, minuend1End}
	minuend2Range := Timerange{minuend2Start, minuend2End}
	subtrahend1Range := Timerange{subtrahend1Start, subtrahend1End}
	subtrahend2Range := Timerange{subtrahend2Start, subtrahend2End}
	subtrahend3Range := Timerange{subtrahend3Start, subtrahend3End}

	var minuendRanges []Timerange
	var subtrahendRanges []Timerange

	minuendRanges = append(minuendRanges, minuend1Range, minuend2Range)
	subtrahendRanges = append(subtrahendRanges, subtrahend1Range, subtrahend2Range, subtrahend3Range)

	result := SubtractTimeranges(minuendRanges, subtrahendRanges)

	expectedRange1Start := getTimeFromTodayUsingHoursMinutesSeconds(9, 15, 0)
	expectedRange1End := getTimeFromTodayUsingHoursMinutesSeconds(10, 0, 0)
	expectedRange2Start := getTimeFromTodayUsingHoursMinutesSeconds(10, 15, 0)
	expectedRange2End := getTimeFromTodayUsingHoursMinutesSeconds(11, 0, 0)

	expected1Range := Timerange{expectedRange1Start, expectedRange1End}
	expected2Range := Timerange{expectedRange2Start, expectedRange2End}

	var expectedRanges []Timerange

	expectedRanges = append(expectedRanges, expected1Range, expected2Range)

	if !reflect.DeepEqual(expectedRanges, result) {
		t.Errorf("Expected: %s, Actual: %s", timerangesToString(expectedRanges), timerangesToString(result))
	}
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
