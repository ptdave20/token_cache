Token Cache
===========
This is an attempt to make token caching a little simpler with Oauth2 token storage.

How to use
----------
`go get github.com/ptdave20/token_cache`
```go
    ctx := context.Background()
	tCache, err:= token_cache.New("backupGroups")
	if err!=nil {
		panic(err)
	}


	if tCache.Config == nil {
		// Read the client file
		b, err := ioutil.ReadFile("client.json")
		if err != nil {
			panic(err)
		}

		config, err := google.ConfigFromJSON(b, admin.AdminDirectoryGroupReadonlyScope, admin.AdminDirectoryGroupMemberReadonlyScope)
		if err != nil {
			panic(err)
		}

		tCache.Config = config
		tCache.Save()
	}

	if tCache.Token == nil {
		// Get the client
		url := tCache.Config.AuthCodeURL("code", oauth2.AccessTypeOffline)
		fmt.Printf("Please visit the following url: \r\n\r\n%s\r\nEnter the code given to you here:", url)
		var code string
		if _, err := fmt.Scanln(&code); err != nil {
			panic(err)
		}

		tok, err := tCache.Config.Exchange(context.Background(), code)
		if err != nil {
			panic(err)
		}

		tCache.Token = tok
		tCache.Save()
	}
```

Contibuting
-----------
Fork the code and make a pull request stating any changes you made and why.

License
-------
Copyright 2018 David Marchbanks

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.