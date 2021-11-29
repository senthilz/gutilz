# gutilz

## Overview
Frequently used utility functions

```go
import "github.com/senthilz/gutilz"
```

## Index

- [func CreateFolder(folder string, forceCreate int) error](<#func-createfolder>)
- [func Dumper(data ...interface{}) (string, error)](<#func-dumper>)


## func [CreateFolder](<https://github.com/senthilz/gutilz/blob/main/gutilz.go#L13>)

```go
func CreateFolder(folder string, forceCreate int) error
```

CreateFolder \- create or force re\-create a folder

## func [Dumper](<https://github.com/senthilz/gutilz/blob/main/gutilz.go#L41>)

```go
func Dumper(data ...interface{}) (string, error)
```

Dumper \- similar to Data::Dumper module in Perl usage: Dumper\(data\, map\[string\]\[string\]\{"indent":"\\t"\, "prefix":""\}\)
