# Service APISIX

## Description
This ia a collection of plugins built and bundled with APISIX.
Furthermore, there is support for hosting static files, designed for SPAs.

## Plugins

All plugins are written in Go and are using the [APISIX Go Plugin Runner](https://github.com/apache/apisix-go-plugin-runner).

- [oidc plugin](docs/oidc/oidc.md)

## Static file hosting

To host static files, you need to put the static files in "/usr/share/nginx/html" inside the container.
The configuration is designed to work for SPAs, so the index.html file will be served for all requests that do not match a static file.
The server will listen on port "9889".
The configuration can be found [here](config.yaml).

## Configuration
The plugin is enabled besides the builtin plugins in the [config.yaml](config.yaml).

In order to use one or multiple plugins, you need to configure it in your apisix.yaml inside the container.

## Build
To build the image, run the following command:

```bash
docker build -t <your_image_name> .
```

## Local development & testing

To run the image locally, you can use the following command:

```bash
docker compose up -d --build
```

This will spin up a keycloak as idp for the oidc plugin besides an apisix with this [configuration](.docker-compose/apisix.yaml).
The keycloak will import a realm including a user with the name "123" and the password "123" on every startup.
