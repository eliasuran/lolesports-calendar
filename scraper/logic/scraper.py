from . import TIMEOUT, URL
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from bs4 import BeautifulSoup

# SETTING UP WEBDRIVER 
options = webdriver.ChromeOptions()
options.add_argument("--headless")
options.add_argument("disable-gpu")

driver = webdriver.Chrome(options=options)

def get_schedule():
    driver.get(URL)

    wait = WebDriverWait(driver, TIMEOUT)
