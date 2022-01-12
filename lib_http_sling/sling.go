package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/sling"
)

const baseURL = "https://api.github.com/"

// Issue github issue (abbreviated)
type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GithubError struct {
	Message string `json:"message"`
	Errors  []struct {
		Resource string `json:"resource"`
		Field    string `json:"field"`
		Code     string `json:"code"`
	} `json:"errors"`
	DocumentationURL string `json:"documentation_url"`
}

// IssueParams github issue parameters
type IssueParams struct {
	Filter    string `url:"filter,omitempty"`
	State     string `url:"state,omitempty"`
	Labels    string `url:"labels,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
	Since     string `url:"since,omitempty"`
}

type IssueService struct {
	sling *sling.Sling
}

func NewIssueService(httpClient *http.Client) *IssueService {
	return &IssueService{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

func (s *IssueService) ListByRepo(owner, repo string, params interface{}) ([]Issue, *http.Response, error) {
	issues := new([]Issue)
	githubError := new(GithubError)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	// 注意此处一定要调用New方法来clone一个sling实例
	resp, err := s.sling.New().Get(path).QueryStruct(params).Receive(issues, githubError)
	if err == nil {
		log.Printf("err: %v\n", err.Error())
		return nil, nil, err
	}
	return *issues, resp, err
}

func main() {

}
