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

def run():
    route = f"https://old.reddit.com/r/GunsNRoses/new/"
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
    formatted_html = soup.prettify()
    print(formatted_html)
    #random_sleep_time = 9
    #time.sleep(random_sleep_time)

run()
