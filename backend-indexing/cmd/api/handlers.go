package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `jason:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Indexing backend up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello from index")
	fmt.Println("Indexing...")
	// Here it should start indexing
}

func (app *application) Search(w http.ResponseWriter, r *http.Request) {
	search_value := r.URL.Query().Get("value")
	gte_time := time.Now().UTC().Add(-time.Minute * time.Duration(30)).Format("2006-01-02T15:04:05Z07:00")
	lt_time := time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")
	zincsearch_url := app.ZincsearchHost + "/es/" + app.ZincsearchIndex + "/_search"

	query := fmt.Sprintf(`{
	  "query": {
	    "bool": {
	      "must": [
	        {
	          "range": {
	            "@timestamp": {
	              "gte": "%v",
                "lt": "%v",
	              "format": "2006-01-02T15:04:05Z07:00"
	            }
	          }
	        },
	        {
	          "query_string": {
	            "query": "%s"
	          }
	        }
	      ]
	    }
	  },
	  "sort": [
	    "-@timestamp"
	  ],
	  "from": 0,
	  "size": 100,
	  "aggs": {
	    "histogram": {
	      "date_histogram": {
	        "field": "@timestamp",
	        "calendar_interval": "",
	        "fixed_interval": "30s"
	      }
	    }
	  }
	}`, gte_time, lt_time, search_value)

	req, err := http.NewRequest("POST", zincsearch_url, strings.NewReader(query))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Accept", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	//_ = app.writeJSON(w, http.StatusOK, body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
