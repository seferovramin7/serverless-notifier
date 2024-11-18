package fetcher

import "time"

type Job struct {
	Success  bool `json:"success"`
	Status   int  `json:"status"`
	Response struct {
		Jobs []struct {
			Title                     string   `json:"title"`
			ComapnyURL1               string   `json:"comapnyURL1,omitempty"`
			ComapnyURL2               string   `json:"comapnyURL2,omitempty"`
			CompanyID                 string   `json:"companyId"`
			CompanyUniversalName      string   `json:"companyUniversalName"`
			CompanyName               string   `json:"companyName"`
			SalaryInsights            string   `json:"salaryInsights"`
			Applicants                int      `json:"applicants"`
			FormattedLocation         string   `json:"formattedLocation"`
			FormattedEmploymentStatus string   `json:"formattedEmploymentStatus"`
			FormattedExperienceLevel  string   `json:"formattedExperienceLevel"`
			FormattedIndustries       string   `json:"formattedIndustries"`
			JobDescription            string   `json:"jobDescription"`
			InferredBenefits          string   `json:"inferredBenefits"`
			JobFunctions              string   `json:"jobFunctions"`
			WorkplaceTypes            []string `json:"workplaceTypes"`
			CompanyData               struct {
				Name                 string `json:"name"`
				Logo                 string `json:"logo"`
				BackgroundCoverImage string `json:"backgroundCoverImage"`
				Description          string `json:"description"`
				StaffCount           int    `json:"staffCount"`
				StaffCountRange      struct {
					StaffCountRangeStart int `json:"staffCountRangeStart"`
					StaffCountRangeEnd   int `json:"staffCountRangeEnd"`
				} `json:"staffCountRange"`
				UniversalName string        `json:"universalName"`
				URL           string        `json:"url"`
				Industries    []string      `json:"industries"`
				Specialities  []interface{} `json:"specialities"`
			} `json:"company_data"`
			CompanyApplyURL string    `json:"companyApplyUrl"`
			JobPostingURL   string    `json:"jobPostingUrl"`
			ListedAt        time.Time `json:"listedAt"`
		} `json:"jobs"`
		Paging struct {
			Total int `json:"total"`
			Start int `json:"start"`
			Count int `json:"count"`
		} `json:"paging"`
	} `json:"response"`
}

func FetchJobs() ([]Job, error) {
	return []Job{}, nil
}
