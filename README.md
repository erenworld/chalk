# Chalk

![chalk](https://davidwalsh.name/demo/chalk-example.png)

chalk is a lightweight Go package for styling terminal output using ANSI colors and attributes.

## Install

```bash
go get github.com/erenworld/chalk
```

> Please note that this is a clone of Fatih arslan's project. All credit goes to him.

### Features
* Foreground and background colors (FgRed, BgBlue, etc...)
* High-intensity colors (FgHiRed, FgHiBlue, etc.)
* Text attributes like `Bold`
* Color disabling support (via NoColor or DisableColor())
* Printable and formattable methods (Print, Printf, Sprintf, etc.)
* Cross-platform color support using github.com/mattn/go-colorable

## Examples

### Standard colors

```go
// Chain SGR parameters
color.Green.Add(color.Bold).Println("Green with bold")
color.Red.Add(color.BgWhite, color.Underline).Printf("Red with Black background and underscore: %s\n", "format too!")

// These are using by default foreground colors.
color.Red("We have red")
color.Yellow("Yellow color too!")
color.Magenta("And many others ..")

// Windows supported too! Just don't forget to change the output to color.Output
fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))
```

### Mix and reuse colors

```go
d := color.New(color.FgCyan).Add(color.Underline)
d.Println("Prints bold cyan.")

// Create custom color objects:
c := color.New(color.fgGreen, color.bgCyan, color.Italic)
c.Print("Italic green with cyan background")

// Mix up foreground and background colors, create new mixes!
red := color.New(color.FgRed)

boldRed := red.Add(color.Bold)
boldRed.Println("This will print text in bold red.")

whiteBg := red.Add(color.BgWhite)
whiteBg.Println("Red text with White background.")

```

### Insert into noncolor strings

```go
// Create Sprint__ functions to mix strings with other non-colorized strings:
yellow := New(FgYellow).SprintFunc()
red := New(FgRed).SprintFunc()
fmt.Printf("this is a %s and this is %s.\n", yellow("warning"), red("error"))

info := New(FgWhite, BgGreen).SprintFunc()
fmt.Printf("this %s rocks!\n", info("package"))
```

### Plug into existing code

```go
color.Set(color.FgYellow)
fmt.Println("Existing text in your codebase will be now in Yellow")
fmt.Printf("This one %s\n", "too")
color.Unset()

color.Set(color.FgMagenta, color.Bold)
defer color.Unset()

fmt.Println("All text will be now bold magenta.")
```

## Credits

* [Eren Turkoglu](https://github.com/erenworld)
* Windows support via @mattn: [colorable](https://github.com/mattn/go-colorable)

## License

The MIT License (MIT) - see [LICENSE.md](./LICENSE) for more details
