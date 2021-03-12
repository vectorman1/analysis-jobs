package jobs

import (
	"log"
	"time"
)

type SymbolUpdateJob struct {
	analysisClient *AnalysisClient
}

func NewSymbolUpdateJob(analysisClient *AnalysisClient) *SymbolUpdateJob {
	return &SymbolUpdateJob{analysisClient: analysisClient}
}

func (j SymbolUpdateJob) Run() {
	log.Println("[SYMBOL JOB] Starting recalculation job")
	start := time.Now()

	err := j.analysisClient.UpdateSymbols()
	if err != nil {
		log.Printf("[HISTORY JOB] Failed update job:\n - elapsed: %v\n", time.Since(start).Seconds())
		return
	}

	log.Printf("[SYMBOL JOB] Finished recalculation job.\n - elapsed:%v\n", time.Since(start).Seconds())
}
