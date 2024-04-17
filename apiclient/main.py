from logic.get_data import get_data
from logic.utils import write_to_json

def main():
    data = get_data()
    error = write_to_json(data)
    if error:
        print("Error writing to json: ", error)
        return

    print("Successfully got and wrote json data")


if __name__ == "__main__":
    main()
