package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andrewarrow/feedback/models"
)

var items = []*models.BaseModel{}

func GetPosts(sub string) []*models.BaseModel {
	//htmlString := doRedditGet()
	//htmlString, _ := os.ReadFile("/data/guns.txt")
	//htmlString, _ := os.ReadFile("guns.txt")

	//html := replaceSmartQuotes(string(htmlString))
	//doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	for i := 1; i < 51; i++ {
		var list []any
		b, _ := os.ReadFile(fmt.Sprintf("data3/data/%d.json", i))
		json.Unmarshal(b, &list)
		for _, thing := range list {
			item := thing.(map[string]any)
			href := item["href"].(string)
			if strings.HasPrefix(href, "/r/GunsNRoses/comments") {
				tokens := strings.Split(href, "/")
				id := tokens[4]
				m := map[string]any{"id_reddit": id}
				m["title"] = item["title"]
				m["href"] = item["href"]
				m["from"] = item["from"]
				items = append(items, models.NewBase(m))
			}
		}
	}

	return items
}

func processDocument(doc *goquery.Document) {
	doc.Find("a.thumbnail").Each(func(i int, p *goquery.Selection) {
		processThumbail(p)
	})
}

func processThumbail(p *goquery.Selection) {
	p.Find("img").Each(func(i int, a *goquery.Selection) {
		printImage(a, p)
	})
}

func printImage(img, a *goquery.Selection) {
	linkText := a.Text()
	href, _ := a.Attr("href")
	src, _ := img.Attr("src")
	fmt.Println(src)
	fmt.Println(href)
	fmt.Println(linkText)
}

func processTitleP(p *goquery.Selection) {
	p.Find("img").Each(func(i int, a *goquery.Selection) {
		printLink(a)
	})
}

func printLink(a *goquery.Selection) {
	linkText := a.Text()
	href, exists := a.Attr("src")
	if exists {
		if strings.HasPrefix(href, "/r/GunsNRoses/comments") {
			fmt.Printf("Link text: %s\n", linkText)
			tokens := strings.Split(href, "/")
			id := tokens[4]
			fmt.Printf(id)
			m := map[string]any{"id_reddit": id}
			m["title"] = linkText
			items = append(items, models.NewBase(m))
		}
	}
}

func doRedditGet() string {
	//"github.com/PuerkitoBio/goquery"
	url := "https://old.reddit.com/r/gunsnroses/new"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:127.0) Gecko/20100101 Firefox/127.0")
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
func doApiRedditGet() string {
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

func replaceSmartQuotes(input string) string {
	// Use strings.NewReplacer to create a replacer for smart quotes
	replacer := strings.NewReplacer(
		string('\u201C'), "\"", // Left double quote
		string('\u201D'), "\"", // Right double quote
		string('\u2018'), "'", // Left single quote
		string('\u2019'), "'", // Right single quote
		string('\u00B4'), "'", // Acute accent (´)
		"â", "'",
		"â", "\"",
		"â", "\"",
	)

	return replacer.Replace(input)
}
