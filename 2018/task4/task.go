package task1

import (
	"regexp"
	"strconv"
	"time"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}

	guardEntersRegex := regexp.MustCompile(`^\[([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+)\].*#([0-9]+) .*$`)
	fallAsleepRegex := regexp.MustCompile(`^\[([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+)\].*falls.*$`)
	wakeUpRegex := regexp.MustCompile(`^\[([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+)\].*wakes.*$`)
	layout := "2006-01-02 15:04"
	guardCalendar := make(map[time.Time]*daySchedule)

	for _, line := range lines {
		guardEntry := guardEntersRegex.FindAllStringSubmatch(line, -1)
		if len(guardEntry) > 0 {
			dateStr := guardEntry[0][1]
			guardID := guardEntry[0][2]
			date, err := time.Parse(layout, dateStr)
			if err != nil {
				return 0, err
			}
			date = date.Add(1 * time.Hour).Truncate(24 * time.Hour)
			_, exists := guardCalendar[date]
			if !exists {
				guardCalendar[date] = newDaySchedule()
			}
			id, err := strconv.Atoi(guardID)
			if err != nil {
				return 0, err
			}
			guardCalendar[date].id = id
		} else {
			fallAsleep := fallAsleepRegex.FindAllStringSubmatch(line, -1)
			if len(fallAsleep) > 0 {
				dateStr := fallAsleep[0][1]
				date, err := time.Parse(layout, dateStr)
				if err != nil {
					return 0, err
				}
				minute := date.Minute()
				date = date.Truncate(24 * time.Hour)
				_, exists := guardCalendar[date]
				if !exists {
					guardCalendar[date] = newDaySchedule()
				}
				guardCalendar[date].events[minute] = 1
			} else {
				wakesUp := wakeUpRegex.FindAllStringSubmatch(line, -1)
				dateStr := wakesUp[0][1]
				date, err := time.Parse(layout, dateStr)
				if err != nil {
					return 0, err
				}
				minute := date.Minute()
				date = date.Truncate(24 * time.Hour)
				_, exists := guardCalendar[date]
				if !exists {
					guardCalendar[date] = newDaySchedule()
				}
				guardCalendar[date].events[minute] = 0
			}
		}
	}
	longest := make(map[int]*longestSleep)
	for _, schedule := range guardCalendar {
		if _, exists := longest[schedule.id]; !exists {
			longest[schedule.id] = &longestSleep{schedules: [][]int{}}
		}
		asleep := 0
		for i := 0; i < 60; i++ {
			if newState, exists := schedule.events[i]; exists {
				asleep = newState
			}
			schedule.midnightSchedule = append(schedule.midnightSchedule, asleep)
			if asleep == 1 {
				longest[schedule.id].sum += 1
			}
		}
		currentSchedules := longest[schedule.id].schedules
		longest[schedule.id].schedules = append(currentSchedules, schedule.midnightSchedule)
	}
	bestSum := 0
	bestID := 0
	var bestSchedules [][]int
	for id, current := range longest {
		if current.sum > bestSum {
			bestSum = current.sum
			bestID = id
			bestSchedules = current.schedules
		}
	}

	bestMinute := -1
	bestAsleepCount := 0
	for i := 0; i < 60; i++ {
		currentAsleepCount := 0
		for _, schedule := range bestSchedules {
			currentAsleepCount += schedule[i]
		}
		if currentAsleepCount > bestAsleepCount {
			bestMinute = i
			bestAsleepCount = currentAsleepCount
		}
	}

	return bestID * bestMinute, nil
}

func bToI(b bool) int {
	if b {
		return 1
	}
	return 0
}

type daySchedule struct {
	id               int
	events           map[int]int
	midnightSchedule []int
	length           int
	sum              int
}

func newDaySchedule() *daySchedule {
	return &daySchedule{
		events:           make(map[int]int),
		midnightSchedule: []int{},
	}
}

type longestSleep struct {
	sum       int
	schedules [][]int
}
