```ts
// for loop
for $i = 0; $i < 10; $i++ {
    console.log($i);
}

$i = 5;

// while loop
while $i == 5 {
    std.tellraw(@a, f"print {$i}")
    $i--;
}

$x, $y = 5, 10;

// conditionals
if $x == 6 || $y - $x == 5 {
    console.log("yahoo");
} else if $y == 10 {
    console.log("10");
} else {
    console.log("other");
}

```
