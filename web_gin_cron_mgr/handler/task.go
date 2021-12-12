package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"

	"github.com/penk110/interview_go/web_gin_cron_mgr/cron_task"
)

var cronTaskH = &CronTaskH{}

type CronTaskH struct {
}

func GetCronTaskH() *CronTaskH {
	return cronTaskH
}

type ListEntryResp struct {
	ID   cron.EntryID `json:"id"`
	Next time.Time    `json:"next"`
	Prev time.Time    `json:"prv"`
}

func (cth *CronTaskH) ListEntry(c *gin.Context) {
	var (
		entries  []cron.Entry
		listResp []*ListEntryResp
		entry    cron.Entry
		i        int
	)
	entries = cron_task.ListEntries()
	listResp = make([]*ListEntryResp, len(entries))
	for i, entry = range entries {
		var resp = &ListEntryResp{
			ID:   entry.ID,
			Next: entry.Next,
			Prev: entry.Prev,
		}

		listResp[i] = resp
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200010,
		"msg":  "ok",
		"data": listResp,
	})

	return
}

func (cth *CronTaskH) DelEntry(c *gin.Context) {
	var (
		entryID int
		err     error
	)

	entryID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200020,
			"msg":  err.Error(),
		})
		return
	}
	cron_task.CronMgr.Remove(cron.EntryID(entryID))

	c.JSON(http.StatusOK, gin.H{
		"code": 200010,
		"msg":  "ok",
		"data": "success",
	})

	return
}
