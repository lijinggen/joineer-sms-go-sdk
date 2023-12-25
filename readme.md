# Joineer sms
Our [home page](http://sms.joineer.cloud/home) 

## Usage
```go
go get -u github.com/lijinggen/joineer-sms-go-sdk
```

## Sample

```go
client, err := NewJoineerClient("your_api_key", "your_api_secret")
if err != nil {
    log.Fatal(err)
}
err = client.Send("phone", "content")
if err != nil {
    log.Fatal(err)
}
```