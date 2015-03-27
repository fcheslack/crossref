package crossref

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Author struct {
	Family string `json:"family"`
	Given  string `json:"given"`
}

type DateTime struct {
	DateParts [][]int `json:"date-parts"`
	Timestamp int64   `json:"timestamp"`
}

func (dt DateTime) TimeTime() time.Time {
	if dt.Timestamp != 0 {
		return time.Unix(dt.Timestamp, 0)
	}
	if len(dt.DateParts[0]) == 3 {
		return time.Date(dt.DateParts[0][0], time.Month(dt.DateParts[0][1]), dt.DateParts[0][2], 0, 0, 0, 0, time.UTC)
	}
	return time.Time{}
}

type Work struct {
	Subtitle       []string `json:"subtitle"`
	Subject        []string `json:"subject"`
	Issued         DateTime `json:"issued"`
	Score          float64  `json:"score"`
	Prefix         string   `json:"prefix"`
	Author         []Author `json:"author"`
	ContainerTitle []string `json:"container-title"`
	ReferenceCount int      `json:"reference-count"`
	Page           string   `json:"page"`
	Deposited      DateTime `json:"deposited"`
	Issue          string   `json:"issue"`
	Title          []string `json:"title"`
	Type           string   `json:"type"`
	DOI            string   `json:"DOI"`
	ISSN           []string `json:"ISSN"`
	URL            string   `json:"URL"`
	Source         string   `json:"source"`
	Publisher      string   `json:"publisher"`
	Indexed        DateTime `json:"indexed"`
	Volume         string   `json:"volume"`
	Member         string   `json:"member"`
}

type WorkResponse struct {
	Status         string `json:"status"`
	MessageType    string `json:"message-type"`
	MessageVersion string `json:"message-version"`
	Work           Work   `json:"message"`
}

func GetWork(doi string) (Work, error) {
	wr := WorkResponse{}
	reqUrl := fmt.Sprintf("https://api.crossref.org/works/%s", doi)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return wr.Work, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	log.Print(string(body))
	err = json.Unmarshal(body, &wr)
	if err != nil {
		return wr.Work, err
	}
	return wr.Work, nil
}
