from logic.utils import get_active_leagues, get_schedule


def get_data():
    data = {
        "Leagues": [
            { "League": "test", "Schedule": [] }
        ]
    }

    leagues = get_active_leagues()

    for league in leagues:
        schedule = get_schedule(league["OverviewPage"])
        if schedule:
            league_data = { "League": "", "Schedule": [] }
            league_data["League"] = league["Event"]
            league_data["Schedule"] = schedule
            data["Leagues"].append(league_data) 

    return data
