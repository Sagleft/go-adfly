# go-adfly

Library to working with adf.ly on Golang

installation

```bash
go get github.com/Sagleft/go-adfly
```

usage example:

```go
client := goadfly.NewClient(userID, apiKeyPublic)
shortLink, err := client.ShortenLink("https://example.com")
if err != nil {
    log.Fatalln(err)
}

log.Println(shortLink)
```

---

![image](https://github.com/Sagleft/Sagleft/raw/master/image.png)

### :globe_with_meridians: [Telegram канал](https://t.me/+VIvd8j6xvm9iMzhi)
