import json
import requests
from mwrogue.esports_client import EsportsClient
from mwrogue.esports_client import EsportsClient
from . import CONSISTENT_TOURNAMENTS, TOURNAMENTS, DATETIME

site = EsportsClient("lol")

def get_active_leagues() -> list:
    print("Getting all leagues")
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
    print(f"Getting schedule for {region}")
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
        for tournament in CONSISTENT_TOURNAMENTS:
            if tournament == league["League Short"] and not "as" in league["League Short"].lower() and not "academy" in league["League Short"].lower():
                data.append(league)

    return data

def get_teams(league: str, amount: int) -> list:
    print(f"Getting data for {league} with {amount} teams")
    res = site.cargo_client.query(
        tables="TournamentRosters",
        fields="Team",
        where=f"Tournament LIKE '%2024%' AND Tournament LIKE '%{league} %' AND Tournament LIKE '%Spring%'" 
    )

    team_names = []

    for team in res:
        if len(team_names) == amount:
            break
        team_name = team["Team"]
        if team_name not in team_names:
            team_names.append(team_name)

    teams = []

    for team in team_names:
        team_data = get_team_data(team)
        teams.append(team_data)

    return teams

def get_team_data(team: str) -> dict:
    print(f"Getting data for {team}")
    parsedTeam = team.replace("'", "''")
    res = site.cargo_client.query(
        tables="Teams",
        fields="Name, Short, Image",
        where=f"Name = '%s'" % parsedTeam
    )

    print(res)

    if len(res) == 0:
        return {}

    return res[0]

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


def write_to_pantry(data, pantry_url):
    if not data:
        return "No data was provided"

    if type(data) != dict:
        return "Data was not the correct type"

    payload = json.dumps(data, indent=4)
    headers = {
      'Content-Type': 'application/json'
    }

    response = requests.request("POST", pantry_url, headers=headers, data=payload)
    
    if response.status_code != 200:
        return "Couldnt write to pantry: "+response.text

    return
