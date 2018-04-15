Cobra Demo
==========

Simple version of a command line parser.

The full version is complex, not implemented yet. You can see more on <https://github.com/spf13/cobra>

```
$ go-cobra-demo --help
This is a very simple demo for cobra, you can see more at: https://github.com/spf13/cobra

  Usage:
    go-cobra-demo [flags]
    go-cobra-demo [command]

  Available Commands:
    help        Help about any command
    version     Show version

  Flags:
    -h, --help   help for go-cobra-demo

  Use "go-cobra-demo [command] --help" for more information about a command.

```

Debug in editor
---------------

One way is to add commands to "Rub/Debug configuration" -> "Program arguments"

Another is add code before `cmd.Execute()` in `main.go`:

```
func main() {
    os.Args = append(os.Args, "version") // Add command here!
	cmd.Execute()
}

```
