from logic import REGIONS
from logic.utils import get_active_leagues, get_all_leagues, get_schedule, get_teams


def get_data():
    print("Starting program..")
    data = { "Active_leagues": [], "All_leagues": [] }

    active_leagues = get_active_leagues()

    for league in active_leagues:
        schedule = get_schedule(league["OverviewPage"])
        if schedule:
            league_data = {}
            league_data["ID"] = league["Event"].split()[0]
            league_data["Name"] = league["Event"]
            league_data["Schedule"] = schedule
            data["Active_leagues"].append(league_data) 

    print("Getting all leagues")
    all_leagues = get_all_leagues()
    for league in all_leagues:
        print(f"Adding data for {league["League"]}")
        league_data = {}
        league_data["ID"] = league["League Short"]
        league_data["Name"] = league["League"]
        league_data["Logo"] = ""
        amount = REGIONS[league["League Short"]]
        teams = get_teams(league["League Short"], amount)
        league_data["Teams"] = teams
        data["All_leagues"].append(league_data)

    return data
