import json
from requests import post


class Gemini:
    URL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key="
    conversation = []

    def __init__(self, key: str, conversation=[]) -> None:
        self.URL = self.URL + key
        self.conversation = conversation

    def ask(self, question) -> tuple[str, bool]:
        userQuestion = {
            "role": "user",
            "parts": [
                {
                    "text": question,
                },
            ],
        }
        self.conversation.append(userQuestion)

        # POST request to the self.URL with headers and body
        headers = {
            "Content-Type": "application/json",
        }

        res = post(self.URL, headers=headers, json={"contents": self.conversation})
        if res.ok:
            try:
                answer = res.json()
                answer = answer["candidates"][0]["content"]
                self.conversation.append(answer)
                return answer["parts"][0]["text"], True
            except KeyError:
                return json.dumps(res.json(), False)
        else:
            return json.dumps(res.json(), False)
