package token_cache

import (
	"runtime"
	"path"
	"os"
	"os/user"
	"golang.org/x/oauth2"
	"encoding/json"
	"io/ioutil"
)

type TokenCache struct {
	clientTokenLocation string	`json:"-"`
	Token *oauth2.Token			`json:"token"`
	Config *oauth2.Config       `json:"config"`
}

func New(productName string) (*TokenCache, error) {
	ret := new(TokenCache)

	outPath := "";

	if runtime.GOOS == "windows" {
		outPath = path.Join("%APPDATA%", "local", productName)
	} else {
		u , err:= user.Current()
		if err!=nil {
			return nil, err
		}
		outPath = path.Join(u.HomeDir, ".config", productName)
	}


	if err := os.MkdirAll(outPath, os.ModePerm); err!=nil {
		if err != os.ErrExist {
			return nil, err
		}
	}

	token := path.Join(outPath, "token.json")
	if _, err := os.Stat(token); err != nil && !os.IsNotExist(err){
		ret.clientTokenLocation = outPath
	} else if err == nil {
		b, err := ioutil.ReadFile(token)
		if err!=nil {
			return nil, err
		} else {
			if err := json.Unmarshal(b, ret); err!=nil {
				return nil, err
			}
		}
	}

	ret.clientTokenLocation = outPath

	return ret, nil
}

func (cache TokenCache) Save() error {
	b, _ := json.MarshalIndent(cache,"","    ")
	token := path.Join(cache.clientTokenLocation, "token.json")
	return ioutil.WriteFile(token, b, os.ModePerm)
}