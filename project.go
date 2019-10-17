package pypi

import (
	"fmt"
)

// ProjectService handles communication with the project related endpoint of
// the PyPi API.
type ProjectService struct {
	client *Client
}

// Project represents a PyPi project.
type Project struct {
	Info struct {
		Author                 string        `json:"author"`
		AuthorEmail            string        `json:"author_email"`
		BugtrackURL            string        `json:"bugtrack_url"`
		Classifiers            []interface{} `json:"classifiers"`
		Description            string        `json:"description"`
		DescriptionContentType interface{}   `json:"description_content_type"`
		DocsURL                interface{}   `json:"docs_url"`
		DownloadURL            string        `json:"download_url"`
		Downloads              struct {
			LastDay   int `json:"last_day"`
			LastMonth int `json:"last_month"`
			LastWeek  int `json:"last_week"`
		} `json:"downloads"`
		HomePage        string      `json:"home_page"`
		Keywords        string      `json:"keywords"`
		License         string      `json:"license"`
		Maintainer      string      `json:"maintainer"`
		MaintainerEmail string      `json:"maintainer_email"`
		Name            string      `json:"name"`
		PackageURL      string      `json:"package_url"`
		Platform        string      `json:"platform"`
		ProjectURL      string      `json:"project_url"`
		ReleaseURL      string      `json:"release_url"`
		RequiresDist    interface{} `json:"requires_dist"`
		RequiresPython  interface{} `json:"requires_python"`
		Summary         string      `json:"summary"`
		Version         string      `json:"version"`
	} `json:"info"`
	LastSerial int `json:"last_serial"`
	Releases map[string]interface{} `json:"releases"`
	Urls []interface{} `json:"urls"`
}

// GetProject retrieves information about a PyPi project.
func (s *ProjectService) GetProject(name string) (*Project, *Response, error) {

	u :=  fmt.Sprintf("pypi/%s/json", pathEscape(name))

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// GetRelease retrieves information about a specific version of a PyPi project.
func (s *ProjectService) GetRelease(name string, version string) (*Project, *Response, error) {

	u :=  fmt.Sprintf("pypi/%s/%s/json", pathEscape(name), pathEscape(version))

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}