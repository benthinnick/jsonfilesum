# JSONFileSum

A test program that I did as part of an interview process. Streams json file, deserializes into objects and then processes those objects in chunks using goroutines. In this implementation it just sums the values


## How to use
Two command-line arguments
-f for filepath
-gr for number of goroutines that process objects

>go run main.go -f="bigData.json" -gr=5

## Test Files
Created by https://json-generator.com/
```
[
  '{{repeat(100, 5)}}',
  {
    a: '{{integer(-1000, 1000)}}',
    b: '{{integer(-1000, 1000)}}'
  }
]
```
