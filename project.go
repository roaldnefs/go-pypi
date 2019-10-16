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
		Author                 string      `json:"author"`
		AuthorEmail            string      `json:"author_email"`
		BugtrackURL            string      `json:"bugtrack_url"`
		Classifiers            []string    `json:"classifiers"`
		Description            string      `json:"description"`
		DescriptionContentType interface{} `json:"description_content_type"`
		DocsURL                interface{} `json:"docs_url"`
		DownloadURL            string      `json:"download_url"`
		Downloads              struct {
			LastDay   int `json:"last_day"`
			LastMonth int `json:"last_month"`
			LastWeek  int `json:"last_week"`
		} `json:"downloads"`
		HomePage        string      `json:"home_page"`
		Keywords        string      `json:"keywords"`
		License         string      `json:"license"`
		Maintainer      interface{} `json:"maintainer"`
		MaintainerEmail interface{} `json:"maintainer_email"`
		Name            string      `json:"name"`
		PackageURL      string      `json:"package_url"`
		Platform        string      `json:"platform"`
		ProjectURL      string      `json:"project_url"`
		ProjectUrls     struct {
			Download string `json:"Download"`
			Homepage string `json:"Homepage"`
		} `json:"project_urls"`
		ReleaseURL     string      `json:"release_url"`
		RequiresDist   interface{} `json:"requires_dist"`
		RequiresPython interface{} `json:"requires_python"`
		Summary        string      `json:"summary"`
		Version        string      `json:"version"`
	} `json:"info"`
	LastSerial int `json:"last_serial"`
	// Releases []struct {
	// 	CommentText string `json:"comment_text"`
	// 	Digests     struct {
	// 		Md5    string `json:"md5"`
	// 		Sha256 string `json:"sha256"`
	// 	} `json:"digests"`
	// 	Downloads     int    `json:"downloads"`
	// 	Filename      string `json:"filename"`
	// 	HasSig        bool   `json:"has_sig"`
	// 	Md5Digest     string `json:"md5_digest"`
	// 	Packagetype   string `json:"packagetype"`
	// 	PythonVersion string `json:"python_version"`
	// 	Size          int    `json:"size"`
	// 	UploadTime    string `json:"upload_time"`
	// 	URL           string `json:"url"`
	// } `json:"releases"`
	Urls []struct {
		CommentText string `json:"comment_text"`
		Digests     struct {
			Md5    string `json:"md5"`
			Sha256 string `json:"sha256"`
		} `json:"digests"`
		Downloads     int    `json:"downloads"`
		Filename      string `json:"filename"`
		HasSig        bool   `json:"has_sig"`
		Md5Digest     string `json:"md5_digest"`
		Packagetype   string `json:"packagetype"`
		PythonVersion string `json:"python_version"`
		Size          int    `json:"size"`
		UploadTime    string `json:"upload_time"`
		URL           string `json:"url"`
	} `json:"urls"`
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