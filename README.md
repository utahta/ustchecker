# Ustream live status checker

[![Build Status](https://travis-ci.org/utahta/ustchecker.svg?branch=master)](https://travis-ci.org/utahta/ustchecker)

## Installing

```
$ go get -u github.com/utahta/ustchecker/cmd/ustchecker
```

## Usage

```
$ ustchecker -h
Usage of ustchecker:
  -name string
        Specifies the ustream channel name

```
```
$ ustchecker -name iss-hdev-payload
live
```

## Example

```go
package main

import (
	"log"

	"github.com/utahta/ustchecker"
)

func main() {
	c, err := ustchecker.New()
	if err != nil {
		log.Fatal(err)
	}

	live, err := c.IsLive("iss-hdev-payload")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(live)
}
```

## Contributing

1. Fork it ( https://github.com/utahta/ustchecker/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

