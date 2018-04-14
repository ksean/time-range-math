package time_range
import . "time"
import . "github.com/ksean/time-range-math/time-range/type"


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

	subtrahendStartInRange := isTimeInRange(minuend, subtrahend.Start)
	subtrahendEndInRange := isTimeInRange(minuend, subtrahend.End)


	/*
	let A = minuend
	Let B = subtrahend

	5 cases

	Case 1:
	A is a subset of, or equal to B
	A ⊆ B
	 */
	if (minuend.Start == subtrahend.Start && minuend.End == subtrahend.End) ||
		(!subtrahendStartInRange && !subtrahendEndInRange &&
			subtrahend.Start.Before(minuend.Start) && subtrahend.End.After(minuend.End)) {

		return result

	/*
	Case 2:
	A intersection with B is a null set (no intersection)
	A ∩ B = ∅
	 */
	} else if !subtrahendStartInRange && !subtrahendEndInRange {
		result = append(result, minuend)

	/*
	Case 3:
	A has a partial intersection with B at the end of its range
	A ∩ B != ∅
	 */
	} else if subtrahendStartInRange && !subtrahendEndInRange {
		minuend.End = subtrahend.Start
		result = append(result, minuend)

	/*
	Case 4:
	A has a partial intersection with B at the start of its range
	A ∩ B != ∅
	 */
	} else if !subtrahendStartInRange && subtrahendEndInRange {
		minuend.Start = subtrahend.End
		result = append(result, minuend)

	/*
	Case 5:
	A is a superset of B, but B != A
	A ⊃ B
	 */
	} else {
		var firstBisection Timerange
		var secondBisection Timerange

		firstBisection.Start = minuend.Start
		firstBisection.End = subtrahend.Start

		secondBisection.Start = subtrahend.End
		secondBisection.End = minuend.End

		result = append(result, firstBisection)
		result = append(result, secondBisection)
	}

	return result
}

func isTimeInRange(timerange Timerange, time Time) bool {
	return timerange.Start.Before(time) && timerange.End.After(time)
}
