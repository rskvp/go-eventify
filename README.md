# Eventify

Eventify is a lightweight low-code platform for running workflows on tiny servers.

## Try it

`Eventify` has a ready-to-use `Docker` image that bundles both the `UI` and the `Server`.

You'll need to install [`Docker`](https://docs.docker.com/engine/install/) to be able to run `Eventify`.

To run `Eventify`, execute the following command:

```bash
docker run -d -p 8080:8080 ghcr.io/assalielmehdi/go-eventify:main
```

**Note**: `Eventify` runs by default on port `8080`. Make sure the port is not used by another program.

## Authors

- El Mehdi Assali - <assalielmehdi@gmail.com>