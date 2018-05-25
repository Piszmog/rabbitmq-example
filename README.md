# RabbitMQ Example
Example if a Pub/Sub model in Go. It is also deployable to Cloud Foundry.

## Packages
* `github.com/goolge/uuid`
  * Generates UUIDs
* `go.uber.org/zap`
  * The logging framework

## Version Control
Using [vgo](https://github.com/golang/go/wiki/vgo) for version control.

## Cloud Ready
As part of this Go example, I wanted to see what the effort would be to deploy to Cloud Foundry, specifically [Pivotal Cloud Foundry](https://run.pivotal.io/).

### Manifest YAML
The manifest file is required for pushing the application to Cloud Foundry.

In the `manifest.yml`, it specifies the application's name, the buildpack (Cloud Foundry can figure out on its own), 
the memory (I am still testing the least amount of memory required), instances, environment variables, and any services to bind to.

#### Rabbitmq Service

### Pushing
To push to Cloud Foundry, change directories to where the `manifest.yml` file is located and run 
`cf push` (ensure you are logged into Cloud Foundry thru the __cli__ and have the __cli__ installed).

Cloud Foundry will build the application and run it.

## Local running
Simply run the `main()` function. 