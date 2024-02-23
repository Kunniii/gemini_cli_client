from art import tprint
import os
from getpass import getpass
from Gemini import Gemini


def loadAPIKey() -> str:
    # Check if the file "gemini_key" exists at ~/.gemini_key
    # If not, prompt user to input the key, create file and save the key in file
    if not os.path.isfile(os.path.join(os.path.expanduser("~"), ".gemini_key")):
        key = getpass("Please enter your API key: ")
        with open(os.path.join(os.path.expanduser("~"), ".gemini_key"), "w") as f:
            f.write(key)
    else:
        with open(os.path.join(os.path.expanduser("~"), ".gemini_key"), "r") as f:
            key = f.read()
    return key


def main():
    key = loadAPIKey()
    gemini = Gemini(key)
    while True:
        question = input("\n>>>>>>> ")
        if question == "!BYE":
            break
        answer, ok = gemini.ask(question)
        if not ok:
            print(f"SYSTEM: Something went wrong")
        else:
            print(f"GEMINI: {answer}")


if __name__ == "__main__":
    tprint("\nGEMINI\n", font="")
    print(f"\tTo exit program, enter !BYE\n")
    main()
