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

# scoreboard syntax
```ts
// scoreboard get
obj:@s;

// scoreboard add / remove
obj:@s -= 1;
obj:@s++;

// assign
obj:@s = 10;
uwu:@s, obj:@s = 10;

// scoreboard opreation
obj:@s -= uwu:@s;
obj:@s = uwu:@s * owo:@s + (10 / twt:@s);

// store
obj:@s ?= <command> // store success
obj:@s = <command> // store result

```

# funciton syntax
```ts
// regular function
func uwu() {
    say hello world;
}
uwu();


// macro function
func uwu(owo, twt) {
    $say $(owo) $(twt);
}
uwu() with uwu::;
uwu("hello", "world");
uwu(owo: "hello", twt: "world");


// compile-time function
@lazy
func uwu(owo) {
    say $(owo);
}
uwu("hello world");


// function in a tag
@tag(minecraft:tick)
func uwu() {
    say spammming;
}
@tag(uwu) 
func uwu() {
    say can be called with #uwu();
}

// blocks. annonymous functions. 
{
    say this is a block;
}

execute run {
    say this runs a block;
}

```
