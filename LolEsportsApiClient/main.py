from logic.get_data import get_data
from logic.utils import write_to_pantry
import os
from dotenv import load_dotenv

def main():
    load_dotenv()
    PANTRY_URL = os.getenv("PANTRY_URL")

    data = get_data()
    error = write_to_pantry(data, PANTRY_URL)
    if error:
        print("Error writing to pantry: ", error)
        return


    print("Successfully got and wrote json data")


if __name__ == "__main__":
    main()
