package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Action int

const (
	BEGINS_SHIFT = iota
	FALLS_ASLEEP
	WAKES_UP
)

type TimeAction struct {
	day    string
	hour   int
	minute int
	guard  string
	action Action
}

func sortTimeActions(timeActions []TimeAction) {
	sort.Slice(timeActions, func(i, j int) bool {
		if timeActions[i].day < timeActions[j].day {
			return true
		}
		if timeActions[i].day > timeActions[j].day {
			return false
		}
		if timeActions[i].hour > timeActions[i].hour {
			return true
		}
		if timeActions[i].hour < timeActions[i].hour {
			return false
		}

		return timeActions[i].minute < timeActions[j].minute
	})
}

func main() {
	inputFileStream, _ := os.Open("./day4/input")
	defer inputFileStream.Close()

	scanner := bufio.NewScanner(inputFileStream)

	var timeActions []TimeAction

	for scanner.Scan() {
		var timeAction TimeAction

		datetimeAndAction := strings.Split(scanner.Text(), "]")
		datetime := datetimeAndAction[0]
		guardAction := datetimeAndAction[1]
		fmt.Sscanf(datetime, "[%s %d:%d]", &timeAction.day, &timeAction.hour, &timeAction.minute)

		if strings.Contains(guardAction, "begins shift") {
			timeAction.action = BEGINS_SHIFT
		}
		if strings.Contains(guardAction, "falls asleep") {
			timeAction.action = FALLS_ASLEEP
		}
		if strings.Contains(guardAction, "wakes up") {
			timeAction.action = WAKES_UP
		}

		if strings.Contains(guardAction, "Guard #") {
			re := regexp.MustCompile("[0-9]+")
			timeAction.guard = re.FindString(guardAction)
		}

		timeActions = append(timeActions, timeAction)
	}

	sortTimeActions(timeActions)

	// for _, row := range timeActions {
	// 	log.Println(row)
	// }

	log.Println("Part 1 Answer:", part1(timeActions))
}

func part1(timeActions []TimeAction) int {
	guardTotalSleep := make(map[string][]int)
	var guard, sleepiestGuard string
	var startMinuteOfSleep, sleepiestGuardFrequencyMinute, totalGuardSleep, firstFrequencyMinute, sleepiestGuardNum int

	for _, timeAction := range timeActions {
		switch timeAction.action {
		case BEGINS_SHIFT:
			guard = timeAction.guard
		case FALLS_ASLEEP:
			startMinuteOfSleep = timeAction.minute
		case WAKES_UP:
			for i := startMinuteOfSleep; i < timeAction.minute; i++ {
				guardTotalSleep[guard] = append(guardTotalSleep[guard], i)
			}
		}
	}

	for guard, minutes := range guardTotalSleep {
		if len(minutes) > totalGuardSleep {
			totalGuardSleep = len(minutes)
			sleepiestGuard = guard
		}
	}

	minutes := make(map[int]int)
	for _, v := range guardTotalSleep[sleepiestGuard] {
		minutes[v]++
	}

	for k, v := range minutes {
		if v > sleepiestGuardFrequencyMinute {
			sleepiestGuardFrequencyMinute = v
			firstFrequencyMinute = k
		}
	}

	sleepiestGuardNum, _ = strconv.Atoi(sleepiestGuard)

	return sleepiestGuardNum * firstFrequencyMinute
}
