# k8s-port-forwarding

## Table of contents
- [Description](#description)
- [Pre-requisites](#pre-requisites)
- [Using prebuilt binaries](#using-prebuilt-binaries)
  - [Auto Completion](#bonus--automatic-completion--composeforward-only-)
  - [Usage](#usage)
  - [Supported CLI Options](#usage)
- [Contributing](#contributing)

## Description
The purpose of this repository is to enhance the ease with which a developer can debug and resolve issues in the beta environment.

The approach used in this repository is to start up port-forwarding for all of a given service's dependencies (excepting databases, redis caches, etc...)

Using the `composeForward` allows you to provide a docker-compose file,
from which applications will be sourced and all the data needed to start port-forwarding collected.

## Pre-requisites
1) Install [kubectl](https://kubernetes.io/docs/tasks/tools/)

In order for this to work, you must have a valid `~/.kube/config` file. If you do not have one yet, you can create one like so:

```bash 
aws eks update-kubeconfig --name {cluster-name} --profile {aws-profile} --region {aws-region}
```


## Using prebuilt binaries
1) Head over to the Releases section of this repository and download the latest version compatible with your machine.
2) Add the downloaded binary to your $PATH variable.
> Note: If you are using Windows, make sure you run everything from within Git Bash.

### Bonus: automatic completion (composeForward only)
Because remembering the exact names of every file in your docker-compose file is a pain,
there are bash and zsh completion scripts that you can utilize to make interacting with `composeForward` more enjoyable.

1) Download the applicable `compose_forward_[bash|zsh]_completer` file.
2) In your relevant .rc file (e.g. .zshrc or .bashrc), add the following line: `source path/to/your/completer`
3) Open a new bash or zsh session, type `composeForward` then hit `TAB` (in bash hit it twice). You should see auto-completion at work!

### Usage
`composeForward [options] [app-names...]`

### Supported Options
```bash 
--file(-f)
# Usage: --file path/to/docker-compose.yml
# Required: false
# Description: >
#  Allows the user to specify the path to the docker-compose file that should be used.
#  If this flag is not provided, the docker-compose.yml file within the executable's directory is used.

--list(-l)
# Usage: --list
# Required: false
# Description: >
#   If the --list or -l flag is provided, no port-forwarding will occur, but an alphabetized list of available services/apps to use will be printed.

--omit(-o)
# Usage: --omit app-dependency-name
# Required: false
# Description: >
#  port-forwarding is not started for dependencies of app(s) with the provided omit name(s).
#  The flag may be provided multiple times to omit multiple dependencies, like so:
#  -o dep-1 -o dep-2 OR -o="dep-1 dep-2"

--service(-s)
# Usage: --service some-docker-service
# Required: false
# Description: >
#   Allows the user to provide the names of specific services (as contained in the docker-compose file)
#   for which to start port-forwarding.
#   The flag may be provided multiple times to omit multiple dependencies, like so:
#  -s svc-1 -o svc-2 OR -s="svc-1 svc-2"
```

## Contributing

### Dev Pre-requisites
1) Install [kubectl](https://kubernetes.io/docs/tasks/tools/)
2) Install [Go](https://go.dev/doc/install)

### Dev Setup
See [Setup](#setup) above.

