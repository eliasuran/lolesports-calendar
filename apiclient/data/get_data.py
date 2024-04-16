from data.utils import get_active_leagues, get_schedule


def get_data():
    leagues = get_active_leagues()

    for league in leagues:
        get_schedule(league)
