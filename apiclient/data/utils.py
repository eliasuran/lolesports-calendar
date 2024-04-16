from mwrogue.esports_client import EsportsClient
from mwrogue.esports_client import EsportsClient
from data import TOURNAMENTS, DATETIME

site = EsportsClient("lol")

def get_active_leagues() -> list:
    res = site.cargo_client.query(
        tables="CurrentLeagues",
        fields="Event, OverviewPage",
    )

    data = []

    for league in res:
        for tournament in TOURNAMENTS:
            if tournament in league["Event"] and not "as" in league["Event"].lower() and not "academy" in league["Event"].lower():
                data.append(league)

    return data

def get_schedule(region: str):
    res = site.cargo_client.query(
        tables="MatchSchedule",
        fields="Team1, Team2, DateTime_UTC, OverviewPage, MatchId",
        where="DateTime_UTC >= '%s' AND OverviewPage = '%s'" % (DATETIME, region)
    )

    print(res)
