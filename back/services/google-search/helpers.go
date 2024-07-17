package googlesearch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/RafaZeero/brand_monitor/types"
	"github.com/RafaZeero/brand_monitor/utils"
)

func replaceSpacesWith(text, replace string) string {
	re := regexp.MustCompile(`\s+`)

	return re.ReplaceAllString(text, replace)
}

func SearchFor(query string) (*types.GoogleSearchApiResponse, error) {
	url := fmt.Sprintf(
		"https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s",
		utils.EnvOrFatal("GOOGLE_API_KEY"),
		utils.EnvOrFatal("GOOGLE_CX"),
		replaceSpacesWith(query, "+"),
	)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data types.GoogleSearchApiResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
