package metrics

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tweet struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

type TweetResponse struct {
	Message string `json:"message"`
	Tweet   Tweet  `json:"tweet"`
}

func SimulatedMetrics() {
	if !strings.EqualFold(os.Getenv("SIMULATED_METRICS"), "true") {
		log.Println("[WARNING] SIMULATED_METRICS disabled, exiting.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	baseInterval := 2 * time.Second
	if val := os.Getenv("SIMULATED_METRICS_INTERVAL"); val != "" {
		if dur, err := time.ParseDuration(val); err == nil {
			baseInterval = dur
		}
	}

	maxTweets := 5
	if n, err := strconv.Atoi(os.Getenv("SIMULATED_METRICS_TWEETS")); err == nil && n > 0 {
		maxTweets = n
	}

	maxUsers := 5
	if n, err := strconv.Atoi(os.Getenv("SIMULATED_METRICS_USERS")); err == nil && n > 0 {
		maxUsers = n
	}

	log.Printf("[INFO] Max tweets: %d, Max users: %d", maxTweets, maxUsers)

	go func() {
		var tweetIDs []string
		createdUsers := 0
		client := &http.Client{Timeout: 5 * time.Second}

		for len(tweetIDs) < maxTweets || createdUsers < maxUsers {
			if len(tweetIDs) < maxTweets {
				tweetText := fmt.Sprintf("Simulated tweet %s", time.Now().Format("150405"))
				body := fmt.Sprintf(`{"text":"%s","userId":"1","username":"simuser"}`, tweetText)
				resp, err := doRequestWithResponse(client, "POST", "http://localhost:8080/api/tweets", body)
				if err == nil {
					raw, _ := io.ReadAll(resp.Body)
					resp.Body.Close()
					var tResp TweetResponse
					if json.Unmarshal(raw, &tResp) == nil && tResp.Tweet.ID != "" {
						tweetIDs = append(tweetIDs, tResp.Tweet.ID)
						log.Printf("[INFO] Tweet created: %s (%d/%d)", tResp.Tweet.ID, len(tweetIDs), maxTweets)
					}
				}
				time.Sleep(randomSleep(200, 700))
			}

			if createdUsers < maxUsers {
				newUser := fmt.Sprintf("slurmuser_%d", time.Now().UnixNano())
				newPass := fmt.Sprintf("pass_%d", rand.Intn(100000))
				regBody := fmt.Sprintf(`{"username":"%s","email":"%s@example.com","password":"%s"}`, newUser, newUser, newPass)
				doRequest(client, "POST", "http://localhost:8080/api/register", regBody)
				log.Printf("[INFO] Registered user: %s (%d/%d)", newUser, createdUsers+1, maxUsers)

				loginPass := newPass
				if rand.Float64() < 0.4 {
					loginPass = fmt.Sprintf("wrongpass_%d", rand.Intn(100000))
					log.Printf("[INFO] Simulating failed login for user: %s", newUser)
				}

				loginBody := fmt.Sprintf(`{"username":"%s","password":"%s"}`, newUser, loginPass)
				doRequest(client, "POST", "http://localhost:8080/api/login", loginBody)

				createdUsers++
				time.Sleep(randomSleep(200, 700))
			}
		}

		log.Printf("[INFO] User and tweet creation complete. Continuing GET requests indefinitely.")

		for {
			doRequest(client, "GET", "http://localhost:8080/api/tweets", "")
			doRequest(client, "GET", "http://localhost:8080/api/users", "")
			doRequest(client, "GET", "http://localhost:8080/api/404", "")

			for _, tid := range tweetIDs {
				if rand.Float64() < 0.2 {
					continue
				}
				doRequest(client, "POST", fmt.Sprintf("http://localhost:8080/api/tweets/%s/like", tid), "")
				doRequest(client, "POST", fmt.Sprintf("http://localhost:8080/api/tweets/%s/share", tid), "")
				doRequest(client, "POST", fmt.Sprintf("http://localhost:8080/api/tweets/%s/slurm", tid), "")
				time.Sleep(randomSleep(200, 700))
			}

			time.Sleep(baseInterval + time.Duration(rand.Intn(3000))*time.Millisecond)
		}
	}()
}

func randomSleep(minMs, maxMs int) time.Duration {
	return time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond
}

func doRequest(client *http.Client, method, url, body string) {
	_, err := doRequestWithResponse(client, method, url, body)
	if err != nil {
		log.Printf("[ERROR] Request failed: %v", err)
	}
}

func doRequestWithResponse(client *http.Client, method, url, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		log.Printf("[ERROR] Can't create request: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Request error: %v", err)
		return nil, err
	}

	if resp.StatusCode >= 400 {
		log.Printf("[WARN] Got HTTP %d for %s %s", resp.StatusCode, method, url)
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		return nil, nil
	}

	return resp, nil
}
