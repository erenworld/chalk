# Chalk
![chalk](https://davidwalsh.name/demo/chalk-example.png)

## API
We use ANSI SGR for personalized color.


## Install

```bash
go get github.com/erenworld/chalk
```

## Examples

### Standard colors

```go
// Print with default foreground colors
color.Cyan.Print("Prints text in cyan.")
color.Blue.Print("Prints text in blue.")

// Chain SGR paramaters
color.Green.Add(color.Bold).Println("Green with bold")
color.Red.Add(color.BgWhite, color.Underline).Printf("Red with Black background and underscore: %s\n", "format too!")
```

### Custom colors

```go
// Create and reuse color objects
c := color.Cyan.Add(color.Underline)
c.Println("Prints bold cyan.")
c.Printf("Thir prints bold cyan %s\n", "too!.")

// Create custom color objects:
c := color.New(color.fgGreen, color.bgCyan, color.Italic)
c.Print("Italic green with cyan background")
```

## Credits

- [Eren Turkoglu](https://github.com/erenworld)

## License

The MIT License (MIT) - see LICENSE.md for more details
