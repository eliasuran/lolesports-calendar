from . import TIMEOUT, URL
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
from bs4 import BeautifulSoup

# SETTING UP WEBDRIVER 
options = webdriver.ChromeOptions()
options.add_argument("--headless=new")
options.add_argument("--disable-gpu")

driver = webdriver.Chrome(options=options)

def get_schedule():
    driver.get(URL)

    wait = WebDriverWait(driver, TIMEOUT)
    wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, ".simplebar-content")))

    page_source = driver.page_source

    soup = BeautifulSoup(page_source, "html.parser")

    game = soup.find(string="BLG")
    if game:
        print(game.text)
        return

    print("didnt find nuttin")
