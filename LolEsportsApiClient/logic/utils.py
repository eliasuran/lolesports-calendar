import json
import requests
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
        data = { "ID": "", "Team1": "", "Team2": "", "DateTime": ""}

        data["ID"] = match["MatchId"]
        data["Team1"] = match["Team1"]
        data["Team2"] = match["Team2"]
        data["DateTime"] = match["DateTime UTC"]

        schedule.append(data)

    return schedule

def get_all_leagues() -> list:
    res = site.cargo_client.query(
        tables="Leagues",
        fields="League, League_Short, Region",
        where="Level = 'Primary' AND IsOfficial = 'Yes'"
    )

    data = []

    for league in res:
        for tournament in TOURNAMENTS:
            if tournament == league["League Short"] and not "as" in league["League Short"].lower() and not "academy" in league["League Short"].lower():
                data.append(league)

    return data

def get_teams(league: str, teams: int) -> list:
    res = site.cargo_client.query(
        tables="TournamentRosters",
        fields="Team",
        where=f"Tournament LIKE '%2024%' AND Tournament LIKE '%{league} %' AND Tournament LIKE '%Spring%'" 
    )

    print(res)

    team_names = []

    for i in res:
        if len(team_names) == teams:
            break
        team_name = i["Team"]
        if team_name not in team_names:
            team_names.append(team_name)

    return team_names

# currently not in use because pantry
def write_to_json(data, path):
    if not data:
        return "No data was provided"

    if type(data) != dict:
        return "Data was not the correct type"

    json_data = json.dumps(data, indent=4)
    with open(path + "data.json", "w") as outfile:
        outfile.write(json_data)

    return


def write_to_pantry(data, pantry_id):
    url = "https://getpantry.cloud/apiv1/pantry/"+pantry_id+"/basket/data"

    if not data:
        return "No data was provided"

    if type(data) != dict:
        return "Data was not the correct type"

    payload = json.dumps(data, indent=4)
    headers = {
      'Content-Type': 'application/json'
    }

    response = requests.request("POST", url, headers=headers, data=payload)
    
    if response.status_code != 200:
        return "Couldnt write to pantry: "+response.text

    return
