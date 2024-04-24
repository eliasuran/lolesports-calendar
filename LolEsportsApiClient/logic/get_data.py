from logic.utils import get_active_leagues, get_all_leagues, get_schedule


def get_data():
    data = {
        "Active_leagues": [{ "ID": "te", "Name": "test", "Schedule": [] }],
        "All_leagues": [ 
            { 
                "ID": "te", 
                "Name": "test", 
                "Logo": "imgsrc", 
                "Teams": [{ "ID": "g2", "Name": "gamers2", "Logo": "g2logo" }] 
            }
        ]
    }

    active_leagues = get_active_leagues()

    for league in active_leagues:
        schedule = get_schedule(league["OverviewPage"])
        if schedule:
            league_data = {}
            league_data["ID"] = league["Event"].split()[0]
            league_data["Name"] = league["Event"]
            league_data["Schedule"] = schedule
            data["Active_leagues"].append(league_data) 

    all_leagues = get_all_leagues()
    for league in all_leagues:
        league_data = {}
        league_data["ID"] = league["League Short"]
        league_data["Name"] = league["League"]
        league_data["Logo"] = ""
        league_data["Teams"] = []
        data["All_leagues"].append(league_data)

    return data
