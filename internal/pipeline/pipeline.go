package pipeline

import (
	"serverless-notifier/internal/database"
	"serverless-notifier/internal/fetcher"
	"serverless-notifier/internal/filter"
	"serverless-notifier/internal/notifier"
)

func Run(query string, page int, locationID string, sortBy string, functionIDs string, postedAgo int, experience int) error {

	jobs, err := fetcher.FetchJobs(query, page, locationID, sortBy, functionIDs, postedAgo, experience)

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
