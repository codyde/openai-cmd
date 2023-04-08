package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	Messages       []Message `json:"messages"`
	MaxTokens      int       `json:"max_tokens"`
	NumCompletions int       `json:"n"`
	Model          string    `json:"model"`
	Stream         bool      `json:"stream"`
	Temperature    float64   `json:"temperature"`
}

type ResponseMessage struct {
	Content string `json:"content"`
}

type CompletionResponse struct {
	Choices []struct {
		Message ResponseMessage `json:"message"`
	} `json:"choices"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY environment variable is not set.")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./openai_completions \"your prompt here\"")
		os.Exit(1)
	}

	client := &http.Client{}
	prompt := strings.Join(os.Args[1:], " ")

	requestBody := CompletionRequest{
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens:      3800,
		NumCompletions: 1,
		Model:          "gpt-3.5-turbo",
		Stream:         false,
		Temperature:    0.7,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		loadingChars := []string{"-", "\\", "|", "/"}
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\rLoading %s", loadingChars[i])
				i = (i + 1) % len(loadingChars)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	resp, err := client.Do(req)
	done <- true // Signal that the API call is completed
	if err != nil {
		fmt.Println("Error making API request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	wg.Wait() // Wait for the loading indicator goroutine to finish
	fmt.Println("\rLoading completed")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: API request failed with status code %d\n", resp.StatusCode)
		fmt.Println(string(body))
		os.Exit(1)
	}

	var completionResponse CompletionResponse
	err = json.Unmarshal(body, &completionResponse)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		os.Exit(1)
	}

	generatedText := completionResponse.Choices[0].Message.Content

	fmt.Println("Answer:")
	fmt.Println(generatedText)
}
