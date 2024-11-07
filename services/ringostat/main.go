package ringostat

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RingostatResponse struct {
	// Define fields based on the API's response JSON structure
	Data string `json:"data"`
}

func main() {
	// Replace with your actual Ringostat API key
	apiKey := os.Getenv("RINGOSTAT_API_KEY")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.ringostat.com/v1/some_endpoint", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error closing response body:", err)
		}
	}(resp.Body)

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var ringostatResp RingostatResponse
	err = json.Unmarshal(body, &ringostatResp)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Output the data
	fmt.Println("Response from Ringostat:", ringostatResp.Data)
}
