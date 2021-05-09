# Codename

> an [RFC1178](https://tools.ietf.org/html/rfc1178) implementation to generate pronounceable, sometimes even memorable, _"superheroe like"_ codenames, consisting of a random combination of adjective and noun.


## Usage

Codename it's a [package](https://golang.org/doc/code#ImportingRemote), so all you need to do is import it into your code (:point_right: [Try it!](https://play.golang.org/p/mVJl2ZL9pwA)):

```go
package main

import (
	"fmt"
	"github.com/lucasepe/codename"
)

func main() {
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 8; i++ {
		name := codename.Generate(rng, 0, false)
		fmt.Println(name)
	}
}
```

This is how the output looks like (since it's random your will be different).

```txt
absolute-karatecha
moving-colleen
game-nova
fine-madrox
pro-penguin
keen-morbius
firm-iron
refined-epoch
```

You can request the addition of a token to create even more entropy [Try it!](https://play.golang.org/p/60sK2OD8bAH)):

```go
package main

import (
	"fmt"
	"github.com/lucasepe/codename"
)

func main() {
	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 8; i++ {
		name := codename.Generate(rng, 4, false)
		fmt.Println(name)
	}
}
```

note the token (with the specified length) added at the end:

```txt
hopeful-toad-men-133b
blessed-man-thing-2bdc
unique-starfox-4271
full-butterfly-2470
accepted-santa-claus-e24e
merry-belphegor-65da
willing-medusa-cdf4
adapting-nightstar-f626
```