package aoc_2018

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	guardRegex = regexp.MustCompile(`.* Guard #(\d*).*`)
)

type lines []string

func (l lines) Len() int {
	return len(l)
}

func (l lines) Less(i, j int) bool {
	ti := parseTimestampFromInputLine(l[i])
	tj := parseTimestampFromInputLine(l[j])
	return ti.Before(tj)
}

func (l lines) Swap(i, j int) {
	t := l[j]
	l[j] = l[i]
	l[i] = t
}

func parseTimestampFromInputLine(line string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04", line[1:17])
	return t
}

func sortInput(input string) lines {
	l := lines(strings.Split(input, "\n"))
	sort.Sort(l)

	return l
}

type guardLog struct {
	guardId       int
	wentToSleepAt time.Time
	wokeUpAt      time.Time
}

func (g guardLog) NumberOfMinutesSlept() int {
	return int(g.wokeUpAt.Sub(g.wentToSleepAt).Minutes())
}

func parseGuardLogs(input string) []guardLog {
	sortedInput := sortInput(input)
	var logs []guardLog

	lastGuardIdSeen := 0
	for k, v := range sortedInput {
		match := guardRegex.FindStringSubmatch(v)
		if len(match) > 0 {
			lastGuardIdSeen, _ = strconv.Atoi(match[1])
		}

		if strings.Contains(v, "wakes up") {
			previous := sortedInput[k-1]
			wentToSleep := parseTimestampFromInputLine(previous)
			wokeUp := parseTimestampFromInputLine(v)
			logs = append(logs, guardLog{guardId: lastGuardIdSeen, wentToSleepAt: wentToSleep, wokeUpAt: wokeUp})
		}
	}

	return logs
}

func Strategy1(input string) int {
	logs := parseGuardLogs(input)

	guard := mostSleepyGuard(logs)
	minute := mostFrequentSleepingMinute(logs, guard)

	return guard * minute
}

func mostFrequentSleepingMinute(logs []guardLog, guardId int) int {
	m := make(map[int]int)

	for _, v := range logs {
		if v.guardId == guardId {
			for i := v.wentToSleepAt.Minute(); i < v.wokeUpAt.Minute(); i++ {
				m[i]++
			}
		}
	}

	max := 0
	minute := 0
	for k, v := range m {
		if v > max {
			max = v
			minute = k
		}
	}

	return minute
}

func mostSleepyGuard(logs []guardLog) int {
	minutesSleptByGuard := make(map[int]int)
	for _, v := range logs {
		minutesSleptByGuard[v.guardId] += v.NumberOfMinutesSlept()
	}

	max := 0
	guard := 0
	for k, v := range minutesSleptByGuard {
		if v > max {
			guard = k
			max = v
		}
	}
	return guard
}

func Strategy2(input string) int {
	logs := parseGuardLogs(input)

	m := make(map[int]map[int]int)

	for _, v := range logs {
		if m[v.guardId] == nil {
			m[v.guardId] = make(map[int]int)
		}

		for i := v.wentToSleepAt.Minute(); i < v.wokeUpAt.Minute(); i++ {
			m[v.guardId][i]++
		}
	}

	guard := 0
	minute := 0
	minuteFreq := 0
	for k, v := range m {
		for min, freq := range v {
			if freq > minuteFreq {
				minuteFreq = freq
				minute = min
				guard = k
			}
		}
	}

	return guard * minute
}
