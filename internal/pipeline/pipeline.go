package pipeline

import (
	"serverless-notifier/internal/database"
	"serverless-notifier/internal/fetcher"
	"serverless-notifier/internal/filter"
	"serverless-notifier/internal/notifier"
)

func Run() error {

	jobs, err := fetcher.FetchJobs()

	if err != nil {
		return err
	}

	filteredJobs, err := filter.FilterJobs(jobs)
	if err != nil {
		return err
	}

	newJobs, err := database.CheckDatabase(filteredJobs)
	if err != nil {
		return err
	}

	err = notifier.SendNotification(newJobs)
	if err != nil {
		return err
	}

	return database.StoreJobIDs(newJobs)
}
