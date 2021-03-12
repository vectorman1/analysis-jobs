package jobs

import (
	"log"
	"time"
)

type HistoryUpdateJob struct {
	analysisClient *AnalysisClient
}

func NewHistoryUpdateJob(analysisClient *AnalysisClient) *HistoryUpdateJob {
	return &HistoryUpdateJob{analysisClient: analysisClient}
}

func (j HistoryUpdateJob) Run() {
	log.Println("[HISTORY JOB] Enqueueing update job...")

	weekday := time.Now().Weekday()
	if weekday == time.Saturday ||
		weekday == time.Sunday {
		log.Println("[HISTORY JOB] Skipping job, not a working day")
	}

	err := j.analysisClient.UpdateHistories()
	if err != nil {
		log.Printf("[HISTORY JOB] Failed update job: %v", err)
		return
	}

	log.Println("[HISTORY JOB] Finished enqueueing update job")
	return
}