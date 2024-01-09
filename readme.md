# Golang Project Layout

<p id="description">A Golang project structure including a simple example.</p>
<p>
    <a href="https://tip.golang.org/doc/go1.21"><img src="https://badgen.net/badge/golang/v1.21/3988FB" alt="shields"></a>
    <a href="https://github.com/uber-go/fx"><img src="https://badgen.net/badge/uber.fx/v1.20/39d353" alt="shields"></a>
</p>

## Goals

The main goals of this project are:
 - Show service structure for Golang projects with a simple example
 - Show how to use DI in Golang by `go.uber.org/fx` and how it can be convenient
 - Show how easy it is to work when project structure has boundaries between layers

## Features

Here are some of the project's features:

 - Ease of scalability for the project
 - Clear description of the project's domain area 
 - No need to worry about errors with cyclic imports
 - Support for modularity
 - Easy to use


## Project Structure

I like splitting the project into several layers, each of them has own responsibilities and boundaries.

For me, any project has three main layers:
  - `CMD` - entry point for the application, in this layer we describe how to run the application
  - `Internal` - the main layer of the application, in this layer we describe the domain area and business logic of the application's
  - `Config` - the layer of the application that is responsible for the configuration of the application's, I call this layer *foundation*


    I want to pay attention to the fact that the project structure is not a dogma, it is just a recommendation.


In *layer relationship diagram* you can see that the project structure has boundaries between layers and layers can address to each other.

`CMD` layer contact with `Internal` and `Config` layers, `Internal` layer contact only with `Config` layer. `Config` layer does not contact with other layers.

**Layer relationship diagram:**
```mermaid
stateDiagram-v2
    [*] --> CMD
    CMD --> Internal
    CMD --> Config
    Internal --> Config
```

Each layer has detailed description in the `readme.md` file inside the layer directory.

- [CMD Layer](cmd/readme.md)
- [Internal Layer](internal/readme.md)
- [Config Layer](config/readme.md)

For the root directory, I would like to use files that are responsible for the service and help to use this service.
In my example, I use `Makefile` and `docker-compose.yml` files. Usually you want to open root directory and use something quickly.
Also in the golang service you can use `go.mod` and `go.sum` files.


## Project Example

In this project, I show you how to use the project structure and how to use DI in Golang.
For example, I use a simple payment service that have two applications: `inbox` and `outbox`.

The `inbox` application responsible for handling requests for creating *payment transaction* from the user, 
and the `outbox` application responsible for handling requests for showing *payment transaction* from the user.

For this example, I use the `http` server, to show how you can use versioning and why you should use `uber.fx`.

P.S. I'm not describe the application logic in details, I only try to show the structure of the application and how to use DI in Golang. Sometimes you can see not optimized code or solutions, because the goal was to show project structure.



