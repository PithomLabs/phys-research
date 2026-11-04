package paper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// openAlexSearch queries OpenAlex for works and returns results with pagination support.
func openAlexSearch(query string, maxResults int, apiKey string) ([]SearchResult, error) {
	if maxResults <= 0 || maxResults > 30 {
		maxResults = 10
	}

	// Apply OpenAlex rate limiting - cross-process persistent
	if err := WaitForOpenAlexRateLimit(); err != nil {
		return nil, fmt.Errorf("OpenAlex rate limit: %w", err)
	}

	base := "https://api.openalex.org/works"
	var out []SearchResult
	perPage := min(maxResults, 25) // OpenAlex max per-page is 200, but we'll use 25 for reasonable batches
	fetched := 0
	cursor := "*"

	for fetched < maxResults {
		if err := WaitForOpenAlexRateLimit(); err != nil {
			return nil, fmt.Errorf("OpenAlex rate limit: %w", err)
		}

		q := fmt.Sprintf("?search=%s&per-page=%d&filter=open_access.is_oa:true&cursor=%s",
			url.QueryEscape(query), perPage, url.QueryEscape(cursor))
		urlStr := base + q

		// Add API key if provided
		if apiKey != "" {
			urlStr += "&api_key=" + url.QueryEscape(apiKey)
		}

		req, err := http.NewRequest("GET", urlStr, nil)
		if err != nil {
			return nil, fmt.Errorf("OpenAlex search: invalid URL %q: %w", urlStr, err)
		}
		req.Header.Set("User-Agent", userAgent)

		resp, err := doWithRetry(req, &http.Client{Timeout: 15 * time.Second}, DefaultRetryConfig())
		if err != nil {
			return nil, fmt.Errorf("OpenAlex search failed: %w", err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("OpenAlex search read failed: %w", err)
		}

		var payload struct {
			Results []struct {
				ID         string                   `json:"id"`
				Title      string                   `json:"title"`
				DOI        string                   `json:"doi"`
				PubYear    int                      `json:"publication_year"`
				Authors    []map[string]interface{} `json:"authorships"`
				OpenAccess map[string]interface{}   `json:"open_access"`
			} `json:"results"`
			Meta struct {
				NextCursor string `json:"next_cursor"`
			} `json:"meta"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			return nil, fmt.Errorf("OpenAlex response parse failed: %w", err)
		}

		for _, r := range payload.Results {
			authors := ""
			if len(r.Authors) > 0 {
				for _, a := range r.Authors {
					if name, ok := a["author"].(map[string]interface{}); ok {
						if n, ok2 := name["display_name"].(string); ok2 {
							if authors != "" {
								authors += ", "
							}
							authors += n
						}
					}
				}
			}
			isOA := false
			oaURL := ""
			if r.OpenAccess != nil {
				if v, ok := r.OpenAccess["is_oa"].(bool); ok && v {
					isOA = true
				}
				if r.OpenAccess != nil {
					if v, ok := r.OpenAccess["oa_url"].(string); ok {
						oaURL = v
					}
				}
			}
			doi := normalizeDOI(r.DOI)
			out = append(out, SearchResult{
				SourceType: "openalex",
				PaperID:    r.ID,
				Title:      r.Title,
				Authors:    authors,
				DOI:        doi,
				PubYear:    r.PubYear,
				IsOA:       isOA,
				OAURL:      oaURL,
				URL:        r.ID,
			})
			fetched++
			if fetched >= maxResults {
				break
			}
		}

		// Check for next page
		if payload.Meta.NextCursor == "" || payload.Meta.NextCursor == cursor {
			break
		}
		cursor = payload.Meta.NextCursor
	}

	return out, nil
}
