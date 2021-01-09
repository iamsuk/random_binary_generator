# random_binary_generator
A oackage to console random binary in your terminal

##Document
This repository provides tools package which contains random.go file.
random.go file has Random_binary func which takes two arguments.

```
func Random_binary(duration int, iter int){}
```

### duration (type integer)
one binary(0 or 1) will created every duration micro seconds.

### iter (type integer)
the number of binary you want to create.
If 0 is provided, binary will be created endlessly.

## Example
your-app
  |-main.go
  |-tools
    |-random.go
```
//in main.go
package main (module name is also main)

import "main/tools"

func main() {
  Random_binary(5,1000) //will create 1000 characters binary every 5 micro seconds.
}
```

