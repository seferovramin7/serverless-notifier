package fetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"io/ioutil"
	"net/http"
	"time"
)

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

func FetchSecrets(parameterNames []string) (map[string]string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your region
	})
	if err != nil {
		return nil, err
	}

	ssmClient := ssm.New(sess)

	resp, err := ssmClient.GetParameters(&ssm.GetParametersInput{
		Names:          aws.StringSlice(parameterNames),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	secrets := make(map[string]string)
	for _, param := range resp.Parameters {
		secrets[*param.Name] = *param.Value
	}

	if len(resp.InvalidParameters) > 0 {
		return nil, errors.New(fmt.Sprintf("Invalid parameters: %v", resp.InvalidParameters))
	}

	return secrets, nil
}

func FetchJobs(query string, page int, searchLocationID string, sortBy string, functionIDsList string, postedAgo int, experience int) ([]Job, error) {
	// Fetch secrets
	secrets, err := FetchSecrets([]string{
		"/linkedin/rapidapi_host",
		"/linkedin/rapidapi_key",
		"/linkedin/rapidapi_url",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch secrets: %v", err)
	}

	apiHost := secrets["/linkedin/rapidapi_host"]
	apiKey := secrets["/linkedin/rapidapi_key"]
	apiURL := secrets["/linkedin/rapidapi_url"]

	client := &http.Client{}

	// Build the query parameters
	url := fmt.Sprintf("%s?query=%s&page=%d&searchLocationId=%s&sortBy=%s&functionIdsList=%s&postedAgo=%d&experience=%d",
		apiURL, query, page, searchLocationID, sortBy, functionIDsList, postedAgo, experience)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("x-rapidapi-host", apiHost)
	req.Header.Set("x-rapidapi-key", apiKey)

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("API request failed with status code: %d", resp.StatusCode))
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Response struct {
			Jobs []Job `json:"jobs"`
		} `json:"response"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Response.Jobs, nil
}
