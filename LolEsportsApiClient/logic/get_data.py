from logic.utils import get_active_leagues, get_schedule


def get_data():
    data = {}

    leagues = get_active_leagues()

    for league in leagues:
        schedule = get_schedule(league["OverviewPage"])
        if schedule:
            data[league["Event"]] = schedule

    return data
