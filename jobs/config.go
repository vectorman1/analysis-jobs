package jobs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Config interface {
	InitConfig() error
	GetSymbolUpdateUrl() string
	GetHistoryUpdateUrl() string
}

type JobsConfig struct {
	urls map[string]string
}

func NewJobsConfig() *JobsConfig {
	return &JobsConfig{}
}

func (j *JobsConfig) InitConfig() error {
	c := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, os.Getenv("SARUMAN_URL"), nil)
	req.Header.Add("Api-Key", os.Getenv("SARUMAN_API_KEY"))
	res, err := c.Do(req)
	if err != nil {
		return err
	}

	bodyb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bodyb, &j.urls); err != nil {
		return err
	}

	return nil
}

func (j *JobsConfig) GetSymbolUpdateUrl() string {
	return j.urls[SYMBOL_UPDATE]
}

func (j *JobsConfig) GetHistoryUpdateUrl() string {
	return j.urls[HISTORY_UPDATE]
}

