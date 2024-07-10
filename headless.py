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
    route = f"https://www.reddit.com/r/GunsNRoses/comments/1dzpdu9/which_track_on_appetite_for_destruction_would_you/"
    options = Options()
    options.add_argument('-headless')
    browser = webdriver.Chrome(options=options)
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

run()
