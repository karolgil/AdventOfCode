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
	guardSleepSchedules, err := getGuardsSleepSchedules(lines)
	if err != nil {
		return 0, err
	}

	bestSum := 0
	bestID := 0
	var bestSchedules [][]int
	for id, current := range guardSleepSchedules {
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

func Solution2(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}
	guardSleepSchedules, err := getGuardsSleepSchedules(lines)
	if err != nil {
		return 0, err
	}

	bestMinute := -1
	bestAsleepCount := 0
	bestID := 0
	for id, current := range guardSleepSchedules {
		for i := 0; i < 60; i++ {
			currentAsleepCount := 0
			for _, schedule := range current.schedules {
				currentAsleepCount += schedule[i]
			}
			if currentAsleepCount > bestAsleepCount {
				bestMinute = i
				bestAsleepCount = currentAsleepCount
				bestID = id
			}
		}
	}
	return bestID * bestMinute, nil
}

func getGuardsSleepSchedules(inputLines []string) (map[int]*sleepSchedule, error) {
	guardsSchedules, err := getGuardsCalendar(inputLines)
	if err != nil {
		return nil, err
	}

	guardsSleepSchedules := make(map[int]*sleepSchedule)
	for _, schedule := range guardsSchedules {
		if _, exists := guardsSleepSchedules[schedule.id]; !exists {
			guardsSleepSchedules[schedule.id] = newSleepSchedule()
		}
		asleep := 0
		for i := 0; i < 60; i++ {
			if newState, exists := schedule.events[i]; exists {
				asleep = newState
			}
			schedule.midnightSchedule = append(schedule.midnightSchedule, asleep)
			if asleep == 1 {
				guardsSleepSchedules[schedule.id].sum += 1
			}
		}
		currentSchedules := guardsSleepSchedules[schedule.id].schedules
		guardsSleepSchedules[schedule.id].schedules = append(currentSchedules, schedule.midnightSchedule)
	}
	return guardsSleepSchedules, nil
}

func getGuardsCalendar(inputLines []string) (map[time.Time]*daySchedule, error) {
	guardEntersRegex := regexp.MustCompile(`^\[([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+).*#([0-9]+) .*$`)
	fallAsleepRegex := regexp.MustCompile(`^\[([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+).*falls.*$`)
	wakeUpRegex := regexp.MustCompile(`^\[([0-9]+-[0-9]+-[0-9]+ [0-9]+:[0-9]+).*wakes.*$`)
	layout := "2006-01-02 15:04"
	guardCalendar := make(map[time.Time]*daySchedule)

	for _, line := range inputLines {
		guardEntry := guardEntersRegex.FindAllStringSubmatch(line, -1)
		if len(guardEntry) > 0 {
			dateStr := guardEntry[0][1]
			guardID := guardEntry[0][2]
			date, err := time.Parse(layout, dateStr)
			if err != nil {
				return nil, err
			}
			date = date.Add(1 * time.Hour).Truncate(24 * time.Hour)
			initDate(guardCalendar, date)
			id, err := strconv.Atoi(guardID)
			if err != nil {
				return nil, err
			}
			guardCalendar[date].id = id
		} else {
			fallAsleep := fallAsleepRegex.FindAllStringSubmatch(line, -1)
			if len(fallAsleep) > 0 {
				dateStr := fallAsleep[0][1]
				date, err := time.Parse(layout, dateStr)
				if err != nil {
					return nil, err
				}
				minute := date.Minute()
				date = date.Truncate(24 * time.Hour)
				initDate(guardCalendar, date)
				guardCalendar[date].events[minute] = 1
			} else {
				wakesUp := wakeUpRegex.FindAllStringSubmatch(line, -1)
				dateStr := wakesUp[0][1]
				date, err := time.Parse(layout, dateStr)
				if err != nil {
					return nil, err
				}
				minute := date.Minute()
				date = date.Truncate(24 * time.Hour)
				initDate(guardCalendar, date)
				guardCalendar[date].events[minute] = 0
			}
		}
	}
	return guardCalendar, nil
}

func initDate(guardCalendar map[time.Time]*daySchedule, date time.Time) {
	_, exists := guardCalendar[date]
	if !exists {
		guardCalendar[date] = newDaySchedule()
	}
}

type daySchedule struct {
	id               int
	events           map[int]int
	midnightSchedule []int
	sum              int
}

func newDaySchedule() *daySchedule {
	return &daySchedule{
		events:           make(map[int]int),
		midnightSchedule: []int{},
	}
}

type sleepSchedule struct {
	sum       int
	schedules [][]int
}

func newSleepSchedule() *sleepSchedule {
	return &sleepSchedule{
		schedules: [][]int{},
	}
}
