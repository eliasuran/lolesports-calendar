from logic.get_data import get_data
from logic.utils import write_to_json
import os
from dotenv import load_dotenv

def main():
    load_dotenv()
    DATA_PATH = os.getenv("DATA_PATH")

    data = get_data()
    error = write_to_json(data, DATA_PATH)
    if error:
        print("Error writing to json: ", error)
        return

    print("Successfully got and wrote json data")


if __name__ == "__main__":
    main()
