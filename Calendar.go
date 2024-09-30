package main

type Interval struct {
	Start int
	End   int
}

type MyCalendar struct {
	intervals []Interval
}

func ConstructorMyCalendar() MyCalendar {
	return MyCalendar{intervals: []Interval{}}
}

func (c *MyCalendar) BookConflict(start int, end int) bool {
	newInterval := Interval{Start: start, End: end}

	for _, existing := range c.intervals {
		if !(newInterval.End <= existing.Start || newInterval.Start >= existing.End) {
			return false
		}
	}

	c.intervals = append(c.intervals, newInterval)
	return true
}

func (c *MyCalendar) Book(start int, end int) bool {
	newInterval := Interval{Start: start, End: end}

	conflict := ConstructorMyCalendar()
	for _, existing := range c.intervals {
		if !(newInterval.End <= existing.Start || newInterval.Start >= existing.End) {
			startConflict := existing.Start
			if newInterval.Start > existing.Start {
				startConflict = newInterval.Start
			}
			endConflict := existing.End
			if newInterval.End < existing.End {
				endConflict = newInterval.End
			}
			flag := conflict.BookConflict(startConflict, endConflict)
			if !flag {
				return false
			}
		}
	}

	c.intervals = append(c.intervals, newInterval)
	return true
}
