API Boilerplate
===============

A Go Restful API boilerplate

## Endpoints

* `/` - will return "OK" if the API is running;

## Development

1. Fork this project;
1. Develop your own handlers in [handlers.go](./handlers.go) or in a different file;
1. Create the Route for your handler [here](./routes.go);
1. Test your code with `Docker`.

### Docker

To start and build the container, type the following command:

```bash
$ docker-compose up -d --build
```

When the container is up and running, test if the API response looks fine:

```bash
$ curl http://127.0.0.1:8888/
```

The command above should return the string `"OK"`.

To take the stack down, run the following command:

```bash
$ docker-compose down
```

**NB**: the container listens to port `80` but, to access it from the host, you have to use port `8888`; this configuration can be changed in the [docker-compose.yml](./docker-compose.yml) file.

Author Information
------------------

This project was created in 2017 by [Davide Di Mauro](https://github.com/darkraiden).
