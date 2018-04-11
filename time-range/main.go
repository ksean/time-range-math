package time_range
import "time"


type Timerange struct {
    start time.Time
    end time.Time
}

func SubtractTimeranges(minuend []Timerange, subtrahend []Timerange) []Timerange {
	var result []Timerange

	result = simpleSubtractManyFromManyTimeranges(minuend, subtrahend)

	return result
}

func simpleSubtractManyFromManyTimeranges(minuend []Timerange, subtrahend []Timerange) []Timerange {
	var result = minuend

	for _, timerange := range subtrahend {
		result = simpleSubtractOneFromManyTimeranges(result, timerange)
	}

	return result
}

func simpleSubtractOneFromManyTimeranges(minuend []Timerange, subtrahend Timerange) []Timerange {
	var result []Timerange

	for _, timerange := range minuend {
		trimmed := simpleSubtractOneFromOneTimerange(timerange, subtrahend)
		result = append(result, trimmed...)
	}

	return result
}

func simpleSubtractOneFromOneTimerange(minuend Timerange, subtrahend Timerange) []Timerange {
	var result []Timerange

	subtrahendStartInRange := timeInRange(minuend, subtrahend.start)
	subtrahendEndInRange := timeInRange(minuend, subtrahend.end)

	// Return empty slice conditions:
	// Exact time range match
	// Subtrahend is a superset of minuend
	if (minuend.start == subtrahend.start && minuend.end == subtrahend.end) ||
		(!subtrahendStartInRange && !subtrahendEndInRange &&
		subtrahend.start.Before(minuend.start) && subtrahend.end.After(minuend.end)) {

		return result
	}

	// 4 cases:
	// 1 No intersection
	if !subtrahendStartInRange && !subtrahendEndInRange {
		result = append(result, minuend)

	// 2 Partial intersection @ end
	} else if subtrahendStartInRange && !subtrahendEndInRange {
		minuend.end = subtrahend.start
		result = append(result, minuend)

	// 3 Partial intersection @ start
	} else if !subtrahendStartInRange && subtrahendEndInRange {
		minuend.start = subtrahend.end
		result = append(result, minuend)

	// 4 Minuend UNION subtrahend == Minuend
	} else {
		var firstBisection Timerange
		var secondBisection Timerange

		firstBisection.start = minuend.start
		firstBisection.end = subtrahend.start

		secondBisection.start = subtrahend.end
		secondBisection.end = minuend.start

		result = append(result, firstBisection)
		result = append(result, secondBisection)
	}

	return result
}

func timeInRange(timerange Timerange, time time.Time) bool {
	return timerange.start.Before(time) && timerange.end.After(time)
}
