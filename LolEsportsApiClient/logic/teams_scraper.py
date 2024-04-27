from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from bs4 import BeautifulSoup
from . import TIMEOUT

options = webdriver.ChromeOptions()
options.add_argument("--headless=new")
options.add_argument("--disable-gpu")
driver = webdriver.Chrome(options=options)

def get_leagues() -> str:
    driver.get("https://lolesports.com/standings")

    wait = WebDriverWait(driver, TIMEOUT)
    wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, "ul.leagues")))

    soup = BeautifulSoup(driver.page_source, "html.parser")

    leagues = soup.find_all("div", class_="name")
    print(leagues[0].text)

    return ""

def scrape_teams():
    return

def scrape_team():
    return

def teams_scraper():
    get_leagues()
    leagues = []
    return
