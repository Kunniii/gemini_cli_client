package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Gemini struct
type Gemini struct {
	URL          string
	Conversation []map[string]interface{}
}

// NewGemini initializes Gemini instance
func NewGemini(key string, conversation []map[string]interface{}) *Gemini {
	return &Gemini{
		URL:          "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=" + key,
		Conversation: conversation,
	}
}

// Ask method sends a question to Gemini API
func (g *Gemini) Ask(question string) (string, bool) {
	userQuestion := map[string]interface{}{
		"role": "user",
		"parts": []map[string]string{
			{"text": question},
		},
	}
	g.Conversation = append(g.Conversation, userQuestion)

	jsonData, err := json.Marshal(map[string]interface{}{"contents": g.Conversation})
	if err != nil {
		return "", false
	}

	res, err := http.Post(g.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil || res.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		body := string(bodyBytes)
		log.Println("ERROR:", err, res.StatusCode, body)
		return "", false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", false
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", false
	}

	answer, ok := data["candidates"].([]interface{})
	if !ok || len(answer) == 0 {
		return "", false
	}

	content := answer[0].(map[string]interface{})["content"].(map[string]interface{})
	g.Conversation = append(g.Conversation, content)

	return content["parts"].([]interface{})[0].(map[string]interface{})["text"].(string), true
}
