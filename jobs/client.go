package jobs

import (
	"fmt"
	"io/ioutil"
	"net/http"
)



type Client interface {
	UpdateHistories() error
	UpdateSymbols() error
}

type AnalysisClient struct {
	client *http.Client
	config *JobsConfig
}

func NewAnalysisClient(config *JobsConfig, client *http.Client) *AnalysisClient {
	return &AnalysisClient{config: config, client: client}
}

func (a *AnalysisClient) UpdateHistories() error {
	res, err := a.client.Post(a.config.GetHistoryUpdateUrl(), "application/json", nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("response was not OK %v", res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return fmt.Errorf("body was empty: %v", res)
	}

	return nil
}

func (a *AnalysisClient) UpdateSymbols() error {
	res, err := a.client.Post(a.config.GetSymbolUpdateUrl(), "application/json", nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("response was not OK %v", res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return fmt.Errorf("body was empty: %v", res)
	}

	return nil
}

