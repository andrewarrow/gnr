package reddit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/andrewarrow/feedback/models"
)

func GetPosts(sub string) []*models.BaseModel {
	js := doRedditGet()
	fmt.Println(len(js))
	return nil
}

func doRedditGet() string {
	apiKey := os.Getenv("REDDIT_API")
	url := "https://api.reddit.com/r/gunsnroses/new"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response from server:", resp.Status)
		return string(b)
	}

	return string(b)
}
