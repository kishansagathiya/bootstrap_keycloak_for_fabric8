package bootstrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const configFilePath = "./configuration.json"

type Config struct {
	Realm         string `json:"realm"`
	Enabled       bool   `json:"enabled"`
	ID            string `json:"id"`
	AdminPassword string `json:"adminpassword"`
	BaseURL       string `json:"baseurl"`
}

type usertoken struct {
	AccessToken string `json:"access_token"`
}

func ParseConfigJSON() (Config, error) {
	file, e := ioutil.ReadFile(configFilePath)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var c Config
	err := json.Unmarshal(file, &c)

	fmt.Println(err)
	return c, err
}

func GetKeyCloakAdminToken() (string, error) {
	user := "admin"
	config, err := ParseConfigJSON()
	if err != nil {
		return "", err
	}
	pwd := config.AdminPassword

	data := url.Values{}
	data.Add("username", user)
	data.Add("password", pwd)
	data.Add("grant_type", "password")
	data.Add("client_id", "admin-cli")

	url := strings.TrimSuffix(config.BaseURL, "/") + "/auth/realms/master/protocol/openid-connect/token"

	client := http.Client{}
	body := []byte(data.Encode())
	req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body)))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	b := buf.Bytes()
	status := resp.StatusCode
	if status < 200 || status >= 400 {
		return "", fmt.Errorf("Failed to get the KeyCloak openid-connect token. Status %d from POST to %s", status, url)
	}
	var u usertoken
	err = json.Unmarshal(b, &u)
	if err != nil {
		return "", err
	}
	token := u.AccessToken
	if len(token) == 0 {
		return "", fmt.Errorf("Missing `access_token` property from KeyCloak openid-connect response %s", string(b))
	}
	return token, nil
}

func CreateRealm(token string) (bool, error) {
	config, err := ParseConfigJSON()
	if err != nil {
		return false, err
	}

	url := strings.TrimSuffix(config.BaseURL, "/") + "/auth/admin/realms"

	client := http.Client{}
	jsonString := `{"enabled":` + strconv.FormatBool(config.Enabled) + `,"id":"` + config.ID + `","realm":"` + config.Realm + `"}`
	fmt.Println(jsonString)
	body := []byte(jsonString)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(body)))
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	status := resp.StatusCode
	if status < 200 || status >= 400 {
		return false, fmt.Errorf("Failed to create the realm. Status %d from POST to %s", status, url)
	}

	return true, nil
}
