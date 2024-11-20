


### Build binary
```bash
cd cmd/lambda/main

GOOS=linux GOARCH=amd64 go build -o main main.go

zip main.zip main
```

### Local invoke
```bash
sam build   

sam local invoke LinkedInJobNotifierFunction --event event.json
```