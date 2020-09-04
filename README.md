# mojikiban-go

mojikiban-go is a CLI and go client library for the [Mojikiban API](https://mojikiban.ipa.go.jp/search/help/api).
 
# Install

```bash
$ go get github.com/yassun/mojikiban-go
```

# Usage

Search

```Go
r := []rune("字")
o := SearchOptions{
  UCS: fmt.Sprintf("0x%x", r[0]), // "0x8fbb"
}
client := NewClient()
res, _ := client.Search(context.Background(), o)

// &{Status:success Find:true Results:[{MJMojizukeimei:MJ010057 ..... ] Count:1}
fmt.Printf("%+v\n", res)
```

Info

```Go
o := GetOptions{
  MJMojizukeimei: "MJ010057",
}
client := NewClient()
res, _ := client.Get(context.Background(), o)

// &{Version:005.02 Date:2018-01-26 MJCharInfo:[{MJMojizukeimei:MJ010057,...}]
fmt.Printf("%+v\n", res)
```
