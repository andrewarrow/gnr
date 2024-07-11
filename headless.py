from selenium import webdriver
from selenium.webdriver.firefox.options import Options
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.chrome.options import Options
import sys
import random
import time
import os
import json
from bs4 import BeautifulSoup
from parsel import Selector


def run(route, filename):
    options = Options()
    #options.headless = True
    options.add_argument("--disable-blink-features=AutomationControlled")
    options.add_argument("--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:127.0) Gecko/20100101 Firefox/127.0")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--disable-infobars")
    options.add_argument("--disable-extensions")
    options.add_argument('-headless')
    browser = webdriver.Chrome(options=options)
    browser.execute_cdp_cmd('Page.addScriptToEvaluateOnNewDocument', {
    'source': '''
        Object.defineProperty(navigator, 'webdriver', {
            get: () => undefined
        })
    '''
    })


    browser.get(route)
    wait = WebDriverWait(browser, 10)
    wait.until(EC.presence_of_element_located((By.TAG_NAME, 'body')))

    rendered_html = browser.page_source
    soup = BeautifulSoup(rendered_html, 'html.parser')

    for script in soup.find_all('script'):
        script.extract()

    soup = BeautifulSoup(str(soup), 'html.parser')

    html_content = soup.prettify()
    selector = Selector(text=html_content)

    titles = selector.css('p.title')

    results = []

    for title in titles:
        href = title.css('a::attr(href)').get()
        text = title.css('a::text').get()

        results.append({ "Link": href, "Text": text })

    with open(f"titles_{filename}.json" , 'w') as f:
      json.dump(results, f, indent=4)

    next_button_href = selector.css('span.next-button a::attr(href)').get()
    print(next_button_href)

if len(sys.argv) > 2:
  route = sys.argv[1]
  title = sys.argv[2]
  run(route, title)
