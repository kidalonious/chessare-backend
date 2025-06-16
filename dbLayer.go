package main 
import (
	"os"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var client *resty.Client
var userExtension string = "rest/v1/user"
var gameExtension string = "rest/v1/game"

type User struct {
	Username string `json:"username"`
	Rating int `json:"rating"`
}

type Game struct {
	Gameid int `json:"gameid"`
	Whiteplayer string `json:"whiteplayer"`
	Blackplayer string `json:"blackplayer"`
	Winner string `json:"winner"`
	Opening string `json:"opening"`
	Gamemoves string `json:"gamemoves"`
	Result string `json:"result"`
}

func GetGameMap(g Game) map[string]string {
	returnMap := make(map[string]string)
	returnMap["whiteplayer"] = g.Whiteplayer
	returnMap["blackplayer"] = g.Blackplayer
	returnMap["winner"] = g.Winner
	returnMap["opening"] = g.Opening
	returnMap["gamemoves"] = g.Gamemoves
	returnMap["result"] = g.Result
	return returnMap
}

func CreateClient() (*resty.Client, error) {
	if client != nil {
		return client, nil
	}
	godotenv.Load()
	apikey := os.Getenv("DBAPIKEY")
	url := os.Getenv("DBURL")

	if apikey == "" || url == "" {
		return nil, fmt.Errorf("apikey or url does not exist")
	}

	client = resty.New().
		SetBaseURL(url).
		SetHeader("apikey", apikey).
		SetHeader("Authorization", "Bearer "+apikey).
		SetHeader("Content-Type", "application/json").
		SetHeader("Prefer", "return=presentation")
	return client, nil
}

func GetUser(username string) (*User, error) {
	client, err := CreateClient()
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}
	var result []User

	response, err := client.R().
		SetQueryParam("username", "ilike." + username + "%").
		SetResult(&result).
		Get(userExtension)
	if err != nil {
		return nil, fmt.Errorf("error with get method")
	}

	if response.IsError() {
		return nil, fmt.Errorf("response gave error code %s", response.Status())
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no user found with username: %s", username)
	}

	return &result[0], nil
}

func GetGamesByPlayer(username string) ([]Game, error) {
	client, err := CreateClient()
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}
	var result []Game
	response, err := client.R(). 
		SetQueryParams(map[string]string{
			"or": fmt.Sprintf("(whiteplayer.eq.%s,blackplayer.eq.%s)", username, username),
			"limit": "10",
		}).
		SetResult(&result).
		Get(gameExtension)
	if err != nil {
		return nil, fmt.Errorf("request gave error: %w", err)
	}
	if response.IsError() {
		return nil, fmt.Errorf("response gave error: %s", response.Status())
	}

	if len(result) == 0 {
		return result, fmt.Errorf("no game exists with username: %s", username)
	}

	return result, nil
}