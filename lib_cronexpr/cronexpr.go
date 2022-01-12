package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		t          string
		tArray     []string
		spec       string
		expression *cronexpr.Expression
		now        time.Time
		now2       time.Time
		ts         []time.Time
		lastRunAt  time.Time
		lastHour   string

		//year   int
		//month  int
		//day    int
		//hour   int
		//minute int
		//second int
	)

	now = time.Now()
	//year = now.Year()
	//month = int(now.Month())
	//day = now.Day()
	//hour = now.Hour()
	//minute = now.Minute()
	//second = now.Second()

	t = "0,5,9,12"
	tArray = strings.Split(t, ",")
	sort.Strings(tArray)
	spec = fmt.Sprintf("0 0 %s 2,3 8 * 2021", t)

	expression = cronexpr.MustParse(spec)
	ts = expression.NextN(now, 10)
	fmt.Println(ts)

	lastHour = GetLastHour(tArray, now.Hour())
	lastRunAt, _ = time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("2021-08-03 %s:00:00", lastHour), time.Local)

	fmt.Println(lastHour, lastRunAt)

	now2 = now.AddDate(0, 0, 1)
	fmt.Println(GetLastNextTime(tArray, now, now2))

	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-08-03 23:59:59", time.Local)
	fmt.Println(ti, ti.Add(1e9))

	ti2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-08-03 02:00:01", time.Local)
	lastRunAt, _ = time.ParseInLocation("2006-01-02 15:04:05", "2021-08-03 02:00:00", time.Local)

	spec = "0 0 0,1,2,19 2,3 8 * 2021"
	expression = cronexpr.MustParse(spec)
	nextTime := expression.Next(ti2)
	fmt.Println("2021-08-02 19:00:01 nextTime===> ", nextTime)
	fmt.Println("2021-08-02 19:00:01 nextTime compare  ===> ", nextTime.After(lastRunAt))

}

func GetLastHour(tArray []string, hn int) string {
	var (
		t  string
		ti int
		tr int
	)

	for _, t = range tArray {
		ti, _ = strconv.Atoi(t)
		if ti == hn {
			return t
		}

		// hn = 10 | ti = 8 tr = 0 | ti > tr -> tr = 8
		// hn = 10 | ti = 9 tr = 8 | ti > tr -> tr = 9
		if ti < hn && ti > tr {
			tr = ti
		}
	}

	return strconv.Itoa(tr)
}

func padTimeMeta(meta string) string {
	if len(meta) == 1 {
		meta = "0" + meta
	}
	return meta
}

func GetLastNextTime(tArray []string, now1 time.Time, now2 time.Time) time.Time {
	var (
		nextTime time.Time
		//nextTimeSpec string
		t  string
		ti int
		tr int
		tn int
	)

	tn = now1.Hour()
	for _, t = range tArray {
		ti, _ = strconv.Atoi(t)
		if ti == tn { // 等于当前小时，则是第二天的当前整点
			tr = ti
			break
		}
		if ti < tn && ti > tr {
			tr = ti
		}
	}

	//nextTimeSpec = fmt.Sprintf("%d-%d-%d %d:00:00", now2.Year(), int(now2.Month()), now2.Day(), tr)
	nextTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2021-08-03 9:00:00", time.Local)

	return nextTime
}
