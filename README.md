## Find EmptyJSON

### Install

You should have installed Golang

```
$ go install github.com/nicewook/emptyjson@latest
```

### Usage

```
emptyjson --dir=testjson
```

- It will find empty JSON file and move file to `emptyjson` directory
- And, It will write the original file path to `emptyjson/emptyjson.txt` 
- Wronng JSON format will not moved to `emptyjson` directory

