# Chalk

![chalk](https://davidwalsh.name/demo/chalk-example.png)

## API

Color lets you use colorized outputs in terms of [ANSI Escape Codes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors). The API can be used in several ways, pick one that suits you. It has support for Windows too!

The package is under ongoing development, checkout for regular updates.

## Install

```bash
go get github.com/erenworld/chalk
```

## Examples

### Standard colors

```go
// Chain SGR paramaters DELETE
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
// OLD VERSION
// Create and reuse color objects
c := color.Cyan.Add(color.Underline)
c.Println("Prints bold cyan.")

// NEW VERSION
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
// OLD VERSION
color.Yellow.Set()
// NEW VERSION
color.Set(color.FgYellow)
fmt.Println("Existing text in your codebase will be now in Yellow")
fmt.Printf("This one %s\n", "too")
color.Unset()

// You can set custom objects too
color.New(color.FgMagenta, color.Bold).Set()

// NEW VERSION
color.Set(color.FgMagenta, color.Bold)
defer color.Unset()

fmt.Println("All text will be now bold magenta.")
```

## Credits

* [Eren Turkoglu](https://github.com/erenworld)
* Windows support via @shiena: [ansicolor](https://github.com/shiena/ansicolor)

## License

The MIT License (MIT) - see LICENSE.md for more details
