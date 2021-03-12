package main

import (
	"github.com/bamzi/jobrunner"
	"github.com/vectorman1/analysis-jobs/jobs"
	"log"
	"net/http"
	"time"
)

func ScheduleJobs(analysisClient *jobs.AnalysisClient) error {
	jobrunner.Start()

	err := jobrunner.Schedule("@every 2h", *jobs.NewSymbolUpdateJob(analysisClient))
	if err != nil {
		return err
	}

	// run every day at 22:00
	err = jobrunner.Schedule("0 22 * * *", jobs.NewHistoryUpdateJob(analysisClient))
	return nil
}


func main() {
	config := jobs.NewJobsConfig()
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to get config %v", err)
		return
	}

	httpClient := &http.Client{
		Timeout: 5*time.Second,
	}

	analysisClient := jobs.NewAnalysisClient(config, httpClient)

	err = ScheduleJobs(analysisClient)
	if err != nil {
		log.Fatalf("failed to schedule jobs: %v", err)
	}

	routes := SetupRoutes()

	if err := routes.Run(":8080"); err != nil {
		log.Fatalf("failed while listening: %v", err)
	}
}
