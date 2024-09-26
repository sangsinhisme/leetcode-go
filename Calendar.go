package main

type Interval struct {
	Start int
	End   int
}

type MyCalendar struct {
	intervals []Interval
}

func Constructor() MyCalendar {
	return MyCalendar{intervals: []Interval{}}
}

func (c *MyCalendar) Book(start int, end int) bool {
	newInterval := Interval{Start: start, End: end}

	for _, existing := range c.intervals {
		if !(newInterval.End <= existing.Start || newInterval.Start >= existing.End) {
			return false
		}
	}

	c.intervals = append(c.intervals, newInterval)
	return true
}
