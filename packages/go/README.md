# Golang Package
This is golang package for s, the repo for this project can be found
at: [S](https://github.com/harryvince/s)

## Example usage
```
package main

import (
    "github.com/harryvince/s/packages/go/s"
)

func main() {
    s.Setup()
    // OR
    s.Setup("test")
}
```
Injects all s secrets into the environment, and an environment can be specified
if required.

If an env is not specified `dev` is used by default.
