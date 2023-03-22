# k8s-port-forwarding

## Table of contents
- [Description](#description)
  - [compose-forward](#compose-forward)
  - [manual-forward](#manual-forward)
- [Using prebuilt binaries](#using-prebuilt-binaries)
  - [auto-completion](#bonus--automatic-completion--composeforward-only-)
- [Pre-requisites](#pre-requisites)
- [Setup](#setup)
- [Running](#running)
  - [Running compose-forward](#run-compose-forward)
  - [Running manual-forward](#run-manual-forward)
- [Add an app to manual-forward](#adding-an-application-to-manual-forward)
## Description
The purpose of this repository is to enhance the ease with which a developer can debug and resolve issues in the beta environment.

The approach used in this repository is to start up port-forwarding for all of a given service's dependencies (excepting databases, redis caches, etc...)

There are two approaches implemented in this repository, the two separate applications can be found in the `cmd` directory:

(Quick gotcha. This port-forwarding occurs via IPv4. Java services you are running locally will utilize IPv6,
which means you could end up running a service on a given port using both IPv4 and IPv6 simultaneously.)

### compose-forward
Using the compose-forward approach allows you to provide a docker-compose file,
from which applications will be sourced and all the data needed to start port-forwarding collected. More detail on this approach is given below.

### manual-forward
Using the manual-forward approach, you can define each individual application manually, tailored to your needs. This requires more work from the get go, but does allow for greater customization.

## Using prebuilt binaries
1) Head over to the Releases section of this repository and download the latest version compatible with your machine.
2) Add the downloaded binary to your $PATH variable.

### Bonus: automatic completion (composeForward only)
Because remembering the exact names of every file in your docker-compose file is a pain,
there are bash and zsh completion scripts that you can utilize to make interacting with `composeForward` more enjoyable.

1) Download the applicable `compose_forward_[bash|zsh]_completer` file.
2) In your relevant .rc file (e.g. .zshrc or .bashrc), add the following line: `source path/to/your/completer`
3) Open a new bash or zsh session, type `composeForward` then hit `TAB` (in bash hit it twice). You should see auto-completion at work!

## Pre-requisites
1) Install [kubectl](https://kubernetes.io/docs/tasks/tools/)
2) Install [Go](https://go.dev/doc/install)

## Setup
> Note: If you are using Windows, make sure you run everything from within Git Bash.

In order for this to work, you must have a valid `~/.kube/config` file. If you do not have one yet, you can create one like so:

```bash 
aws eks update-kubeconfig --name {cluster-name} --profile {aws-profile} --region {aws-region}

# For example, to setup your kube config file for beta cloud services, you would do this:
aws eks update-kubeconfig --name beta-cloud-services --profile dev_access --region us-east-1
```

## Running
As mentioned previously, there are two ways to use this service. Each approach allows the following arguments to be provided:

```bash 
--omit(-o)
# Usage: --omit app-dependency-name
# Required: false
# Description: >
#  port-forwarding is not started for dependencies of app(s) with the provided omit name(s).
#  The flag may be provided multiple times to omit multiple dependencies, like so:
#  -o dep-1 -o dep-2 OR -o="dep-1 dep-2"
```

### Run compose-forward
The docker-compose route allows for the following additional arguments:
```bash 
--service(-s)
# Usage: --service some-docker-service
# Required: false
# Description: >
#   Allows the user to provide the names of specific services (as contained in the docker-compose file)
#   for which to start port-forwarding.
#   The flag may be provided multiple times to omit multiple dependencies, like so:
#  -s svc-1 -o svc-2 OR -s="svc-1 svc-2"

--list(-l)
# Usage: --list
# Required: false
# Description: >
#   If the --list or -l flag is provided, no port-forwarding will occur, but an alphabetized list of available services/apps to use will be printed.

--file(-f)
# Usage: --file path/to/docker-compose.yml
# Required: false
# Description: >
#  Allows the user to specify the path to the docker-compose file that should be used.
#  If this flag is not provided, the docker-compose.yml file within the executable's directory is used.
```

1) Checkout the docker-compose branch: `git checkout docker-compose`
2) Run `go mod tidy`
3) Run `go build -o composeForward cmd/compose-forward/*`
4) Run `./forward [options] [app-names...]`

### Run manual-forward
1) Checkout the docker-compose branch: `git checkout manual`
2) Run `go mod tidy`
3) Run `go build -o manualForward  cmd/manual-forward/*`
4) Run `./forward --app {service name in docker-compose file} [--omit]`

# Adding an application to manual-forward
If you are using the manual-forward application route, you will undoubtedly need to add your own application and/or services. 
The following details the process of doing so:

1) Determine which services the application you are adding depends on.
2) Add your service to the `manual/service/vars.go` file. If I were adding a k8s service named "magic", my addition would look like this:
    ```go
    package service
    
    var (
        Magic = Service{
            K8sName:      "service/magic",  // The name of the k8s service. 
            K8sNamespace: "magic-services", // The namespace that the k8s service belongs to.
            DefaultPort:  5013,             // The local port that port-forwarding should go through.
            ForwardToPort: 8080,            // The port on the k8s service to forward to. If not provided, defaults to 80.
        }
    )
    ```
3) Add your application to the `manual/app/{applicationName}.go`. If I were adding an application named "magician", my addition would like this:
    ```go
   // filename: magician.go
   package app
   
   import "github.com/chad-bekmezian-snap/k8s-port-forwarding/manual/service"
   
   var Magician = App{
        {
            Service: service.Magic,
            DefaultPort: 7373, // Can be used to override the DefaultPort specified in service.Magic
        },
   }
   ```
4) Expose your new application to the user by adding an entry to the `appNameToApp` map in `cmd/manual-forward/main.go`:
    ```go
    package main
   
   import "github.com/chad-bekmezian-snap/k8s-port-forwarding/manual/app"     
   
    var appNameToApp = map[string]app.App{
        "magician":   app.Magician,
    }
    ```

5) Finally, run your new application like so:
   `go run cmd/manual-forward/* --app magician`