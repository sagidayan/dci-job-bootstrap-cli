## DCI Job bootstrap

A simple CLI tool to bootstrap a DCI job with sane defaults and a run script. WIP

```

Usage:
  dci-bootstrap [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  new         Will create a new [subcommand]. (for example a job)

Flags:
  -h, --help     help for dci-bootstrap
  -t, --toggle   Help message for toggle

Use "dci-bootstrap [command] --help" for more information about a command.
```

#### new
```
Will create a new [subcommand]. (for example a job)

Usage:
  dci-bootstrap new [command]

Available Commands:
  job         Create a new Job folder

Flags:
  -h, --help   help for new

Use "dci-bootstrap new [command] --help" for more information about a command.

```

##### job
```
Create a new Job folder

Usage:
  dci-bootstrap new job [flags]

Flags:
  -f, --force             Force create. Might override existing jobs
  -h, --help              help for job
  -n, --name string       Job Name
      --tags strings      Job Tags. Comma seperated.
  -t, --topic DCI_Topic   DCI topic. i.e. OCP-4.11

```
