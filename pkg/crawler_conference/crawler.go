package crawler_conference

import (
	"encoding/json"
	"io"
	"io/ioutil"
	. "my_ieee/internel/utils"
	"net/http"
	"strings"
	"time"
)

// Todo(yuting): make all urls into config file
// Todo(yuting): make a conference-to-conference_number map

func get_issue_number(conference_number string) int {
	// Concatenate url
	meta_data_url := Str_Concate("https://ieeexplore.ieee.org/rest/publication/home/metadata?pubid=", conference_number)
	// Build http request header
	req, err := http.NewRequest("GET", meta_data_url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	referer := Str_Concate("https://ieeexplore.ieee.org/xpl/conhome/", conference_number, "/proceeding")
	req.Header.Set("Referer", referer)
	// Send http GET request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// Read response body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// Parse json
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	// Get issue number
	m1, ok := data.(map[string]interface{})
	if !ok {
		panic("data is not map[string]interface{}")
	}
	m2, ok := m1["currentIssue"].(map[string]interface{})
	if !ok {
		panic("m1[\"currentIssue\"] is not map[string]interface{}")
	}
	issue_number, ok := m2["issueNumber"].(float64)
	if !ok {
		panic("m2[\"issueNumber\"] is not float64")
	}
	return int(issue_number)
}

func DoCrawler(conference_number string) [2][]string {
	// 1. Get the issue number
	issue_number := Str_Itoa(get_issue_number(conference_number))

	// 2. Compose url
	result_titles := make([]string, 0)
	result_links := make([]string, 0)
	page_number := 1
	for {
		toc_url := Str_Concate("https://ieeexplore.ieee.org/rest/search/pub/", conference_number, "/issue/", issue_number, "/toc?count=100&sortType=asc&pageNumber=", Str_Itoa(page_number))
		referer := Str_Concate("https://ieeexplore.ieee.org/xpl/conhome/", conference_number, "/proceeding?pageNumber=", Str_Itoa(page_number))
		payload := Str_Concate(`{"pageNumber":`, Str_Itoa(page_number), `,"punumber":"`, conference_number, `","isnumber":`, issue_number, `}`)
		// 3. Build http request header
		req, err := http.NewRequest("POST", toc_url, strings.NewReader(payload))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Referer", referer)
		// 4. Send http POST request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		// 5. Read response body
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		// 6. Parse json
		var data interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}
		m1, ok := data.(map[string]interface{})
		if !ok {
			break
		}

		// 7. Get titles and links
		records, ok := m1["records"].([]interface{})
		if !ok {
			break
		}

		for _, record := range records {
			record, ok := record.(map[string]interface{})
			if !ok {
				panic("record is not map[string]interface{}")
			}
			title := record["highlightedTitle"].(string)
			link := Str_Concate("https://ieeexplore.ieee.org/stampPDF/getPDF.jsp?tp=&arnumber=", record["articleNumber"].(string), "&ref=")
			result_titles = append(result_titles, title)
			result_links = append(result_links, link)
		}

		page_number++
		// 8. Sleep 1 second to avoid being banned
		time.Sleep(1 * time.Second)
	}
	return [2][]string{result_titles, result_links}
}
