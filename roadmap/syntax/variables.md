
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

# run time variables (fake players)
```ts
$x = 3
$y = 3
$z = ($x + $y) / 2
```

# compile time variables
```ts
&x = 3
&y = 3
&z = (&x + &y) / 2
```
