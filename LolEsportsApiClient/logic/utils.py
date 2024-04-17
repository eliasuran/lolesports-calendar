import json
from mwrogue.esports_client import EsportsClient
from mwrogue.esports_client import EsportsClient
from . import TOURNAMENTS, DATETIME

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

    schedule = []

    for match in res:
        data = { "id": "", "Team1": "", "Team2": "", "DateTime": ""}

        data["id"] = match["MatchId"]
        data["Team1"] = match["Team1"]
        data["Team2"] = match["Team2"]
        data["DateTime"] = match["DateTime UTC"]

        schedule.append(data)

    return schedule

def write_to_json(data, path):
    if not data:
        return "No data was provided"

    if type(data) != dict:
        return "Data was not the correct type"

    json_data = json.dumps(data, indent=4)
    with open(path + "data.json", "w") as outfile:
        outfile.write(json_data)

    return
