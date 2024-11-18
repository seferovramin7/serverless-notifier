package database

import "serverless-notifier/internal/fetcher"

func CheckDatabase(jobs []fetcher.Job) ([]fetcher.Job, error) {
	return jobs, nil
}

func StoreJobIDs(jobs []fetcher.Job) error {
	return nil
}
