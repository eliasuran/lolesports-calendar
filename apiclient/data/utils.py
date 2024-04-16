from mwrogue.esports_client import EsportsClient
from mwrogue.esports_client import EsportsClient
from data import TOURNAMENTS

site = EsportsClient("lol")

def get_active_leagues() -> list:
    res = site.cargo_client.query(
        tables="CurrentLeagues",
        fields="Event",
    )

    data = []

    for league in res:
        for tournament in TOURNAMENTS:
            if tournament in league["Event"] and not "as" in league["Event"].lower() and not "academy" in league["Event"].lower():
                data.append(league["Event"])

    return data

def get_schedule(region: str):
    res = site.cargo_client.query(
        tables="MatchSchedule",
        fields="Event",
    )
