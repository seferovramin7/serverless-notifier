package filter

import "serverless-notifier/internal/fetcher"

func FilterJobs(jobs []fetcher.Job) ([]fetcher.Job, error) {
	return jobs, nil
}
