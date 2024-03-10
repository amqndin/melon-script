
# nbt syntax 
```ts
// data get 
storage::uwu
storage::uwu * 2 // scale specified

// data modify
storage::uwu = "Hello, World!" // set value
storage::uwu = storage::owo // set from
storage::uwu = storage::twt[0:1] // set string
// arrays
storage::uwu >> "Hello, World!" // prepend 
storgae::uwu << storage::owo // append 
storage::uwu ^3 "Hello, World!" // insert value

// store 
storage::uwu ?= <command> // store success
storage::uwu = <command> // store result scale 1 int
storage::uwu = (float) <command> // store result scale 1 float
storage::uwu = 2 * (byte) <command> // store result scale 2 byte

// data merge
storage:: += {key: "value"}

// data remove
storage::uwu.del()

```
