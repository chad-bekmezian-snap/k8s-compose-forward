# cs-port-forwarding

## Description
The purpose of this repository is to enhance the ease with which a developer can debug and resolve issues in the beta environment.

The approach used in this repository is to start up port-forwarding for all of a given service's dependencies (excepting databases, redis caches, etc...)

There are two approaches implemented in this repository, the two separate applications can be found in the `cmd` directory:

### compose-forward
Using the compose-forward approach allows you to provide a docker-compose file,
from which applications will be sourced and all the data needed to start port-forwarding collected. More detail on this approach is given below.

### manual approach
Using the manual-forward approach, you can define each individual application manually, tailored to your needs. This requires more work from the get go, but does allow for greater customization.

## Pre-requisites
1) Install [kubectl](https://kubernetes.io/docs/tasks/tools/)
2) Install [Go](https://go.dev/doc/install)

## Setup
In order for this to work, you must have a valid `~/.kube/config` file. If you do not have one yet, you can create one like so:

```bash 
aws eks update-kubeconfig --name {cluster-name} --profile {aws-profile} --region {aws-region}

# For example, to setup your kube config file for beta cloud services, you would do this:
aws eks update-kubeconfig --name beta-cloud-services --profile dev_access --region us-east-1
```

## Running
As mentioned previously, there are two ways to use this service. Each approach allows the following arguments to be provided:

```bash 
--app(-a)
# Usage: --app my-service-name
# Required: true
# Description: >
#  port-forwarding is started for all services upon which the provided app(s) depend. 
#  The flag may be provided multiple times to start multiple apps, like so:
#  -a app-1 -a app-2 OR -a="app-1 app-2"

--omit(-o)
# Usage: --omit app-dependency-name
# Required: false
# Description: >
#  port-forwarding is not started for dependencies of app(s) with the provided omit name(s).
#  The flag may be provided multiple times to omit multiple dependencies, like so:
#  -o dep-1 -o dep-2 OR -o="dep-1,dep-2"
```

### Run compose-forward
The docker-compose route allows for one more argument:
```bash 
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
4) Run `./forward --app {service name in docker-compose file} [--omit|--file]`

### Run manual-forward
1) Checkout the docker-compose branch: `git checkout manual`
2) Run `go mod tidy`
3) Run `go build -o manualForward  cmd/manual-forward/*`
4) Run `./forward --app {service name in docker-compose file} [--omit]`
