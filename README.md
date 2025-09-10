# Chalk

![chalk](https://davidwalsh.name/demo/chalk-example.png)

## API

We use ANSI SGR for personalized color.

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

- [Eren Turkoglu](https://github.com/erenworld)

## License

The MIT License (MIT) - see LICENSE.md for more details
