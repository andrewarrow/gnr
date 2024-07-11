from selenium import webdriver
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import sys
import json
from bs4 import BeautifulSoup
from parsel import Selector

def fetch_page(route):
    print(route)
    options = Options()
    options.add_argument("--disable-blink-features=AutomationControlled")
    options.add_argument("--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:127.0) Gecko/20100101 Firefox/127.0")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--disable-infobars")
    options.add_argument("--disable-extensions")
    options.add_argument('-headless')
    
    browser = webdriver.Chrome(options=options)

    browser.get(route)
    wait = WebDriverWait(browser, 10)
    wait.until(EC.presence_of_element_located((By.TAG_NAME, 'body')))

    rendered_html = browser.page_source
    browser.quit()
    
    return rendered_html

def parse_page(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')

    for script in soup.find_all('script'):
        script.extract()

    soup = BeautifulSoup(str(soup), 'html.parser')
    selector = Selector(text=soup.prettify())

    results = []
    tops = selector.css('div.thing')
    for thing in tops:
        title = thing.css('p.title')
        href = title.css('a::attr(href)').get()
        text = title.css('a::text').get()
        tagline = thing.css('p.tagline')
        fromUser = tagline.css('a::text').get()
        results.append({ "href": href, "title": text, "from": fromUser })

    next_button_href = selector.css('span.next-button a::attr(href)').get()
    
    return results, next_button_href

def run(start_route, filename):
    route = start_route
    count = 1
    all_results = []

    while route:
        html_content = fetch_page(route)
        results, next_button_href = parse_page(html_content)

        with open(f"data/{count}.json", 'w') as f:
            json.dump(results, f, indent=4)

        count += 1
        route = next_button_href

if __name__ == "__main__":
    if len(sys.argv) > 1:
        start_route = sys.argv[1]
        filename = "results"
        run(start_route, filename)
    else:
        print("Please provide the initial route as a command-line argument.")

