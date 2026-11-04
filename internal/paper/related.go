package paper

import (
	"encoding/json"
	"fmt"
)

func Related(paperID, direction string, maxResults int) ([]SearchResult, error) {
	return relatedPapers(paperID, direction, maxResults)
}

func relatedPapers(paperID string, direction string, maxResults int) ([]SearchResult, error) {
	if maxResults <= 0 || maxResults > 30 {
		maxResults = 10
	}
	var endpoint string
	switch direction {
	case "references", "ref":
		endpoint = fmt.Sprintf("https://api.openalex.org/works/%s/referenced_works", paperID)
	case "cited", "citing", "both":
		endpoint = fmt.Sprintf("https://api.openalex.org/works/%s/cited_by_works", paperID)
	default:
		endpoint = fmt.Sprintf("https://api.openalex.org/works/%s/referenced_works", paperID)
	}

	q := fmt.Sprintf("?per-page=%d", maxResults)
	urlStr := endpoint + q

	resp, err := safeFetch(urlStr, 15)
	if err != nil {
		return nil, fmt.Errorf("OpenAlex related lookup failed: %w", err)
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
	}

	var singleWork struct {
		ID         string                 `json:"id"`
		Title      string                 `json:"title"`
		DOI        string                 `json:"doi"`
		PubYear    int                    `json:"publication_year"`
		OpenAccess map[string]interface{} `json:"open_access"`
	}

	if err := json.Unmarshal(resp, &payload); err == nil && len(payload.Results) > 0 {
		var out []SearchResult
		for _, r := range payload.Results {
			authors := ""
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
			isOA := false
			oaURL := ""
			if r.OpenAccess != nil {
				if v, ok := r.OpenAccess["is_oa"].(bool); ok && v {
					isOA = true
				}
				if v, ok := r.OpenAccess["oa_url"].(string); ok {
					oaURL = v
				}
			}
			out = append(out, SearchResult{
				SourceType: "openalex",
				PaperID:    r.ID,
				Title:      r.Title,
				Authors:    authors,
				DOI:        r.DOI,
				PubYear:    r.PubYear,
				IsOA:       isOA,
				OAURL:      oaURL,
				URL:        r.ID,
			})
		}
		return out, nil
	}

	if err := json.Unmarshal(resp, &singleWork); err == nil && singleWork.ID != "" {
		isOA := false
		oaURL := ""
		if singleWork.OpenAccess != nil {
			if v, ok := singleWork.OpenAccess["is_oa"].(bool); ok && v {
				isOA = true
			}
			if v, ok := singleWork.OpenAccess["oa_url"].(string); ok {
				oaURL = v
			}
		}
		return []SearchResult{
			{
				SourceType: "openalex",
				PaperID:    singleWork.ID,
				Title:      singleWork.Title,
				DOI:        singleWork.DOI,
				PubYear:    singleWork.PubYear,
				IsOA:       isOA,
				OAURL:      oaURL,
				URL:        singleWork.ID,
			},
		}, nil
	}

	return nil, fmt.Errorf("OpenAlex related lookup returned unexpected response shape (status verified but structure unrecognized)")
}
