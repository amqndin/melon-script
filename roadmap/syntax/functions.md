
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

