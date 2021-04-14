# incforge DevOps Task

### Introduction
This repository contains a simple application written in Golang that allows a user to save and get notes from a MongoDB collection.

Your task is to containerise the Golang Application contained in this repository. As this application has a hard dependency on MongoDB, it would also make sense to containerise that and provide a secure way of allowing the two containers to communicate.

**Bonus points** if you are able to achieve this by creating a Kubernetes manifest!

### Build the Project

The following steps assume that you are using a Linux based system. You should be left with an executable called `if-devops-test`

```bash
$ git clone https://github.com/IncForge/if-devops-test.git
$ cd if-devops-test
$ go build .
```

### Running the Application

In order to run the application, you will need to ensure that you have set the MongoDB Connection String environment variable.

The Application is exposed on TCP Port 8000.

The Application has two endpoints:

* **GET /notes** (get a list of all notes persisted to the database)
* **POST /notes** (create a new note)

To create a new note, you will need to POST a JSON body like the below:

```json
{
	"note": "This is a new Note!"
}
```

```bash
$ MONGODB_CONNECTION_STRING="my-connection-string" ./if-devops-test

::1 - - [14/Apr/2021:10:04:01 +0100] "GET /notes HTTP/1.1" 200 194
::1 - - [14/Apr/2021:10:07:10 +0100] "POST /notes HTTP/1.1" 200 42
```

### Useful Resources

* [Golang - Installation Guide](https://golang.org/doc/install)
* [Docker - Getting started](https://docs.docker.com/get-started)
* [MongoDB - Docker Hub Image](https://hub.docker.com/_/mongo)