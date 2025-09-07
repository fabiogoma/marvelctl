package marvel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/fabiogoma/marvelctl/internal"
	"github.com/fabiogoma/marvelctl/internal/models"
)

func NewClient() (*models.Client, error) {
	publicKey, err := internal.GetConfigByKey("public_key")
	if err != nil {
		return nil, fmt.Errorf("error fetching public key: %v", err)
	}
	privateKey, err := internal.GetConfigByKey("private_key")
	if err != nil {
		return nil, fmt.Errorf("error fetching private key: %v", err)
	}

	client := &models.Client{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		BaseUrl:    "https://gateway.marvel.com/v1/public",
		HTTPClient: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}

	return client, nil
}

func generateAuthParameters(publicKey, privateKey string) (ts, hash string) {
	ts = strconv.FormatInt(time.Now().Unix(), 10)

	data := []byte(ts + privateKey + publicKey)

	hash = fmt.Sprintf("%x", md5.Sum(data))
	return ts, hash
}

func GetCharacterByName(name string, client *models.Client) (*models.Character, error) {
	return fetchCharacter("name", name, client)
}

func SearchCharacterByName(name string, client *models.Client) (*models.Character, error) {
	return fetchCharacter("nameStartsWith", name, client)
}

func fetchCharacter(paramKey string, paramValue string, client *models.Client) (*models.Character, error) {
	ts, hash := generateAuthParameters(client.PublicKey, client.PrivateKey)

	endpoint := client.BaseUrl + "/characters"
	params := url.Values{}
	params.Set(paramKey, paramValue)
	params.Set("ts", ts)
	params.Set("apikey", client.PublicKey)
	params.Set("hash", hash)

	reqUrl := endpoint + "?" + params.Encode()
	resp, err := client.HTTPClient.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Printf("Error closing the body: %v", err)
		}
	}()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var result models.Character
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
