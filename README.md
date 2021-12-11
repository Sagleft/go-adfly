# go-adfly

Library to working with adf.ly on Golang

usage example:

```go
client := goadfly.NewClient(userID, apiKeyPublic)
shortLink, err := client.ShortenLink("https://example.com")
if err != nil {
    log.Fatalln(err)
}

log.Println(shortLink)
```
