# gcurl
It is a curl but with json support

## Installing
```shell
go install gcurl
```

## File command
it open json request file and do the request

Usage:
  gcurl file [file name] [flags]

Examples:

file contain example for get request

```json
{
    "method": "GET",
    "url": "https://jsonplaceholder.typicode.com/todos/1",
    "headers": {
        "accept": "application/json",
    }
}
```

file contain example for post request
```json
{
    "method": "POST",
    "url": "https://jsonplaceholder.typicode.com/posts",
    "headers": {
        "accept": "application/json",
    },
    "body": {
        "title": "foo",
            "body": "bar",
            "userId": 1
    }
}
```

## Get command
http get request

Usage:
  gcurl get [url] [flags]

Examples:

```shell
gcurl get -H '{"accept": "application/json"}' https://jsonplaceholder.typicode.com/todos/1
``` 

Flags:
  -H, --headers string   headers

## Post command
http post request

Usage:
  gcurl post [url] [flags]

Examples:

```shell
gcurl post -H '{"accept": "application/json"}' \
-B '{"title": "foo", "body": "bar", "userId": 1}' \
https://jsonplaceholder.typicode.com/posts
```


Flags:

  -B, --body string      body

  -H, --headers string   headers

