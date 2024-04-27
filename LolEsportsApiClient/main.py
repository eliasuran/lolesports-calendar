from logic.get_data import get_data
from logic.utils import get_all_leagues, get_teams, write_to_pantry
import os
from dotenv import load_dotenv

def main():
    teams = get_teams("LEC", 10)
    print(teams)
    return

    load_dotenv()
    PANTRY_ID = os.getenv("PANTRY_ID")

    data = get_data()
    error = write_to_pantry(data, PANTRY_ID)
    if error:
        print("Error writing to pantry: ", error)
        return


    print("Successfully got and wrote json data")


if __name__ == "__main__":
    main()
