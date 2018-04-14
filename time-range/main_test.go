package time_range

import (
	"testing"
	. "time"
	"reflect"
	. "github.com/ksean/time-range-math/time-range/type"
	"strings"
	"strconv"
	. "github.com/ksean/time-range-math/time-range/impl"
	"os"
)

var testDate = Unix(0, 0)
var implementations = []AlgebraOfSets{}

func TestMain(m *testing.M) {
	implementations = append(implementations, SimpleAlgebraOfSets{})
	retCode := m.Run()
	os.Exit(retCode)
}

// Test given examples
func TestMinusPartialIntersection(t *testing.T) {

	minuendRanges := []Timerange{getTimerangeFromString("9:00-10:00")}
	subtrahendRanges := []Timerange{getTimerangeFromString("9:00-9:30")}

	expectedResult := []Timerange{getTimerangeFromString("9:30-10:00")}

	testAll(t, expectedResult, minuendRanges, subtrahendRanges)
}

func TestMinusSameRange(t *testing.T) {

	minuendRanges := []Timerange{getTimerangeFromString("9:00-10:00")}
	subtrahendRanges := []Timerange{getTimerangeFromString("9:00-10:00")}

	var expectedResult []Timerange

	testAll(t, expectedResult, minuendRanges, subtrahendRanges)
}

func TestMinusNonIntersectingRange(t *testing.T) {

	minuendRanges := []Timerange{getTimerangeFromString("9:00-9:30")}
	subtrahendRanges := []Timerange{getTimerangeFromString("9:30-15:00")}

	expectedResult := []Timerange{getTimerangeFromString("9:00-9:30")}

	testAll(t, expectedResult, minuendRanges, subtrahendRanges)
}

func TestManyRangesMinusOneRange(t *testing.T) {

	minuendRanges := []Timerange{
		getTimerangeFromString("9:00-9:30"),
		getTimerangeFromString("10:00-10:30"),
	}

	subtrahendRanges := []Timerange{getTimerangeFromString("9:15-10:15")}

	expectedResult := []Timerange{
		getTimerangeFromString("9:00-9:15"),
		getTimerangeFromString("10:15-10:30"),
	}

	testAll(t, expectedResult, minuendRanges, subtrahendRanges)
}

func TestManyRangesMinusManyRanges(t *testing.T) {

	minuendRanges := []Timerange{
		getTimerangeFromString("9:00-11:00"),
		getTimerangeFromString("13:00-15:00"),
	}

	subtrahendRanges := []Timerange{
		getTimerangeFromString("9:00-9:15"),
		getTimerangeFromString("10:00-10:15"),
		getTimerangeFromString("10:15-12:30"),
		getTimerangeFromString("12:30-16:00"),
	}

	expectedResult := []Timerange{
		getTimerangeFromString("9:15-10:00"),
		getTimerangeFromString("10:15-11:00"),
	}

	testAll(t, expectedResult, minuendRanges, subtrahendRanges)
}

// Test specific helpers ****************************************************************************

func testAll(t *testing.T, expected []Timerange, minuend []Timerange, subtrahend []Timerange) {

	for _, implementation := range implementations {
		actual := implementation.Subtract(minuend, subtrahend)

		assertEqual(t, expected, actual)
	}
}

func assertEqual(t *testing.T, expected []Timerange, actual []Timerange) {

	// Check for equality of two slices of Timeranges
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"Expected: %+v\nActual: %+v",
			expected,
			actual,
		)
	}
}

func getTimerangeFromString(timerange string) Timerange {

	times := strings.Split(timerange, "-")

	return Timerange{
		Start:getTimeFromString(times[0]),
		End: getTimeFromString(times[1]),
	}
}

func getTimeFromString(time string) Time {

	unitsOfTime := strings.Split(time, ":")

	hours, _ := strconv.Atoi(unitsOfTime[0])
	minutes, _ := strconv.Atoi(unitsOfTime[1])

	return getTimeFromHoursMinutes(hours, minutes)

}

func getTimeFromHoursMinutes(hours int, minutes int) Time {

	return Date(testDate.Year(), testDate.Month(), testDate.Day(), hours, minutes, 0, 0, UTC)
}
