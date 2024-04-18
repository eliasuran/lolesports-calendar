from logic.utils import get_active_leagues, get_schedule


def get_data():
    data = {
        "Leagues": [
            { "ID": "te", "Name": "test", "Schedule": [] }
        ]
    }

    leagues = get_active_leagues()

    for league in leagues:
        schedule = get_schedule(league["OverviewPage"])
        if schedule:
            league_data = {}
            league_data["ID"] = league["Event"].split()[0]
            league_data["Name"] = league["Event"]
            league_data["Schedule"] = schedule
            data["Leagues"].append(league_data) 

    return data
