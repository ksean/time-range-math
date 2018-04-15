package impl

import (
	. "time"
	. "github.com/ksean/time-range-math/time-range/type"
)

type SortedAlgebraOfSets struct {}

type timeNode struct {
	id        int
	isStart   bool
	isMinuend bool
	time      Time
	next      *timeNode
}

func (s SortedAlgebraOfSets) Subtract(minuend []Timerange, subtrahend []Timerange) []Timerange {
	var result []Timerange

	var root *timeNode

	root = insertSortTimeranges(root, minuend, true)
	root = insertSortTimeranges(root, subtrahend, false)

	result = parseTimeNodes(root)

	return result
}

func insertSortTimeranges(root *timeNode, timeranges []Timerange, isMinuend bool) *timeNode {

	for id, timerange := range timeranges {

		// Initialize
		if root == nil {
			end := timeNode{
				 id:        id,
				 isStart:   false,
				 isMinuend: isMinuend,
				 time:      timerange.End,
				 next:      nil,
			}

			root = &timeNode{
				id:        id,
				isStart:   true,
				isMinuend: isMinuend,
				time:      timerange.Start,
				next:      &end,
			}
		} else {
			root = insertTimerange(root, timerange, id, isMinuend)
		}
	}

	return root
}

func insertTimerange(root *timeNode, timerange Timerange, id int, isMinuend bool) *timeNode {

	startNode := timeNode{
		id:        id,
		isStart:   true,
		isMinuend: isMinuend,
		time:      timerange.Start,
		next:      nil,
	}

	endNode := timeNode{
		id:        id,
		isStart:   false,
		isMinuend: isMinuend,
		time:      timerange.End,
		next:      nil,
	}

	insertTimeNode(root, startNode)
	insertTimeNode(root, endNode)

	return root
}

func insertTimeNode(root *timeNode, insert timeNode) *timeNode {

	cursor := root
	prev := root

	// Replace root if node is new earliest time
	if insert.time.Before(cursor.time) {
		insert.next = cursor
		return &insert
	}

	// Otherwise insert sort it
	for {

		if insert.time.Before(cursor.time) {
			prev.next = &insert
			insert.next = cursor
			break
		}

		// Insert at end
		if cursor.next == nil {
			cursor.next = &insert
			break
		}

		prev = cursor
		cursor = cursor.next
	}

	return root
}

func parseTimeNodes(root *timeNode) []Timerange {

	var minuendNodes []timeNode
	var subtrahendNodes []timeNode
	var result []Timerange
	var lastEndTime *Time

	cursor := root

	for cursor != nil {

		if lastEndTime == nil {
			lastEndTime = &cursor.time
		}

		if cursor.isStart {

			if cursor.isMinuend {

				minuendNodes = append(minuendNodes, *cursor)

				// Subtrahend nodes
			} else {

				if len(subtrahendNodes) == 0 {

					for _, minuendNode := range minuendNodes {

						latest := latestTime(*lastEndTime, minuendNode.time)

						if latest == cursor.time {
							continue
						}

						result = append(result, Timerange{
							Start: latest,
							End:   cursor.time,
						})
					}
				}

				subtrahendNodes = append(subtrahendNodes, *cursor)
			}

			// END nodes
		} else {

			if cursor.isMinuend {

				if len(subtrahendNodes) == 0 {
					startNodeIndex := indexOfStartNode(minuendNodes, *cursor)

					result = append(result, Timerange{
						Start: latestTime(*lastEndTime, minuendNodes[startNodeIndex].time),
						End:   cursor.time,
					})
				}

				lastEndTime = &cursor.time
				minuendNodes = remove(minuendNodes, indexOfStartNode(minuendNodes, *cursor))

				// Subtrahend nodes
			} else {

				if len(subtrahendNodes) == 0 {

					for _, minuendNode := range minuendNodes {
						result = append(result, Timerange{
							Start: latestTime(*lastEndTime, minuendNode.time),
							End:   cursor.time,
						})
					}
				}

				lastEndTime = &cursor.time
				subtrahendNodes = remove(subtrahendNodes, indexOfStartNode(subtrahendNodes, *cursor))
			}


		}

		cursor = cursor.next
	}

	return result
}

func indexOfStartNode(haystack []timeNode, needle timeNode) int {

	for index, hay := range haystack {

		if needle.id == hay.id &&
			needle.isMinuend == hay.isMinuend {

			return index
		}
	}

	return -1
}

func latestTime(t1 Time, t2 Time) Time {

	if t1.After(t2) {

		return t1

	} else {

		return t2
	}
}

func remove(nodes []timeNode, s int) []timeNode {

	return append(nodes[:s], nodes[s+1:]...)
}
