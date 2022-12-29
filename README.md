# example show help and exit issue with urfave/cli/v2 

While using [urfave/cli/v2](http://github.com/urfave/cli/v2) to write some command-line applications, I came across a scenario where I was using a `Before` field to check if there the correct number of arguments needed for the command to work. If there was not enough arguments, call `cli.ShowCommandHelpAndExit()`. However, no help is printed! I am offering this example in case someone knows what I might be doing wrong.

### help for the main application
```shell
$ go run ./cmd/ct --help
NAME:
   ct - A new cli application

USAGE:
   ct [global options] command [command options] [arguments...]

VERSION:
   1.0

COMMANDS:
   first    the first command
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### help for the 'first' command
```shell
$ go run ./cmd/ct first --help
NAME:
   ct first - the first command

USAGE:
   ct first command [command options] [arguments...]

COMMANDS:
   subaa    subcommand aa
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

### help for the sub-command 'subaa' of the command 'first'
```shell
$ go run ./cmd/ct first subaa --help
DBG => before "first"; args: &["subaa" "--help"]
NAME:
   ct first subaa - subcommand aa

USAGE:
   ct first subaa [command options] [arguments...]

CATEGORY:
   A

OPTIONS:
   --help, -h  show help (default: false)
```

### the failure cases that do not print help when I think they should... 🕵️‍♂️
```shell
$ go run ./cmd/ct first             
DBG => before "first"; args: &[]
DBG => not enough args, show help and exit for "first"
exit status 1
$ go run ./cmd/ct first subaa
DBG => before "first"; args: &["subaa"]
DBG => before "subaa"; args: &[]
DBG => not enough args, show help and exit for "subaa"
exit status 1
```

### when the commands run successfully ✅
```shell
$ go run ./cmd/ct first a
DBG => before "first"; args: &["a"]
FIRST complete
$ go run ./cmd/ct first subaa a
DBG => before "first"; args: &["subaa" "a"]
DBG => before "subaa"; args: &["a"]
SUBAA complete
```