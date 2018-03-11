package token_cache

import (
	"runtime"
	"path"
	"os"
	"golang.org/x/oauth2"
	"encoding/json"
	"io/ioutil"
)

type TokenCache struct {
	clientTokenLocation string	`json:"-"`
	token *oauth2.Token			`json:"token,omitempty"`
	config *oauth2.Config       `json:"config,omitempty"`
}

func New(productName string) (*TokenCache, error) {
	ret := new(TokenCache)

	outPath := "";

	if runtime.GOOS == "windows" {
		outPath = path.Join("%APPDATA%", "local", productName)
	} else {
		outPath = path.Join("~", ".config", productName)
	}


	if err := os.MkdirAll(ret.clientTokenLocation, os.ModePerm); err!=nil {
		if err != os.ErrExist {
			return nil, err
		}
	}

	token := path.Join(outPath, "token.json")
	b, err := ioutil.ReadFile(token)
	if err!=nil && os.IsNotExist(err) {
		return nil, err
	}
	if err := json.Unmarshal(b, ret); err!=nil {
		return nil, err
	}

	ret.clientTokenLocation = outPath

	return ret, nil
}

func (cache *TokenCache) GetToken() (*oauth2.Token) {
	return cache.token
}

func (cache *TokenCache) SetToken(token *oauth2.Token) {
	cache.token = token
}

func (cache *TokenCache) SetConfig(config *oauth2.Config) {
	cache.config = config
}

func (cache *TokenCache) GetConfig() *oauth2.Config {
	return cache.config
}

func (cache *TokenCache) Save() error {
	b, _ := json.Marshal(cache)
	token := path.Join(cache.clientTokenLocation, "token.json")
	// Try to delete the file
	if err := os.Remove(token); !os.IsNotExist(err) {
		return err
	}

	return ioutil.WriteFile(token, b, os.ModePerm)
}