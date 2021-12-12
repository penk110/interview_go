package cron_task

import (
	"github.com/robfig/cron/v3"

	"log"
)

var CronMgr *cronMgr

type cronMgr struct {
	*cron.Cron
}

func init() {
	CronMgr = &cronMgr{cron.New()}

	//CronMgr.AddFunc("* * * * *", health)
	CronMgr.AddFunc("@every 1s", health)
	CronMgr.AddFunc("* * * * *", health2)

	go CronMgr.Run()
}

func AddTask(spec string, f func()) {

	CronMgr.AddFunc(spec, f)
	return
}

type EntryID int

func RemoveTask(entryID EntryID) {
	CronMgr.Remove(cron.EntryID(entryID))
}

func ListEntries() []cron.Entry {
	return CronMgr.Entries()
}

func health() {
	log.Printf("[cronMgr] health ...")
}

func health2() {
	log.Printf("[cronMgr] health2 ...")
}
