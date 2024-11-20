package database

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"serverless-notifier/internal/fetcher"
	"time"
)

var dynamoClient *dynamodb.Client
var tableName = "linkedin_jobs"

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	dynamoClient = dynamodb.NewFromConfig(cfg)
}

func CheckDatabase(jobs []fetcher.Job) ([]fetcher.Job, error) {
	var newJobs []fetcher.Job

	for _, job := range jobs {
		input := &dynamodb.GetItemInput{
			TableName: aws.String(tableName),
			Key: map[string]types.AttributeValue{
				"JobID": &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobPostingURL},
			},
		}

		result, err := dynamoClient.GetItem(context.TODO(), input)
		if err != nil {
			return nil, err
		}

		// If the item doesn't exist, add it to the new jobs list
		if result.Item == nil {
			newJobs = append(newJobs, job)
		}
	}

	return newJobs, nil
}

//func StoreJobIDs(jobs []fetcher.Job) error {
//	for _, job := range jobs {
//		input := &dynamodb.PutItemInput{
//			TableName: aws.String(tableName),
//			Item: map[string]types.AttributeValue{
//				"JobID":   &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobPostingURL},
//				"JobData": &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobPostingURL}, // Adjust this based on your Job structure
//			},
//		}
//
//		_, err := dynamoClient.PutItem(context.TODO(), input)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func StoreJobIDs(jobs []fetcher.Job) error {
	// Mock data structure matching the fetcher.Job definition
	mockJobs := []fetcher.Job{
		{
			Success: true,
			Status:  200,
			Response: struct {
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
			}{
				Jobs: []struct {
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
				}{
					{
						Title:                     "Software Engineer",
						ComapnyURL1:               "https://example.com/company1",
						ComapnyURL2:               "https://example.com/company2",
						CompanyID:                 "12345",
						CompanyUniversalName:      "ExampleUniversalName",
						CompanyName:               "Example Company",
						SalaryInsights:            "Competitive Salary",
						Applicants:                10,
						FormattedLocation:         "Remote, USA",
						FormattedEmploymentStatus: "Full-time",
						FormattedExperienceLevel:  "Mid-Senior level",
						FormattedIndustries:       "Software Development",
						JobDescription:            "Develop and maintain software applications.",
						InferredBenefits:          "Health, Dental, Vision",
						JobFunctions:              "Engineering, Development",
						WorkplaceTypes:            []string{"Remote", "Hybrid"},
						CompanyData: struct {
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
						}{
							Name:                 "Example Company",
							Logo:                 "https://example.com/logo.png",
							BackgroundCoverImage: "https://example.com/cover.png",
							Description:          "An innovative company.",
							StaffCount:           500,
							StaffCountRange: struct {
								StaffCountRangeStart int `json:"staffCountRangeStart"`
								StaffCountRangeEnd   int `json:"staffCountRangeEnd"`
							}{
								StaffCountRangeStart: 100,
								StaffCountRangeEnd:   1000,
							},
							UniversalName: "ExampleUniversalName",
							URL:           "https://example.com",
							Industries:    []string{"Software", "Technology"},
							Specialities:  []interface{}{"Development", "Consulting"},
						},
						CompanyApplyURL: "https://example.com/apply",
						JobPostingURL:   "https://example.com/job12345",
						ListedAt:        time.Now(),
					},
				},
				Paging: struct {
					Total int `json:"total"`
					Start int `json:"start"`
					Count int `json:"count"`
				}{
					Total: 1,
					Start: 0,
					Count: 1,
				},
			},
		},
	}

	// Use the mockJobs instead of input jobs
	for _, job := range mockJobs {
		// Validate that JobPostingURL is non-empty
		if job.Response.Jobs[0].JobPostingURL == "" {
			return fmt.Errorf("missing JobPostingURL for job: %+v", job)
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item: map[string]types.AttributeValue{
				"url":     &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobPostingURL},
				"JobData": &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobDescription}, // Replace with actual JobData
			},
		}

		_, err := dynamoClient.PutItem(context.TODO(), input)
		if err != nil {
			return err
		}
	}

	return nil
}
