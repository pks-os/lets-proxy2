Need for go generate:
* Install https://github.com/gojuno/minimock to PATH
* Install https://github.com/gobuffalo/packr to PATH

Need for boulder tests:
* docker, see tests/run_docker_boulder.sh for run boulder

Fake DNS - set to IP of devel computer, allowed from docker.

Must bind docker port 4000 to local port 4000 (for integration tests).
If use docker-machine - need ```docker-machine ssh <machine-name> -L 4001:localhost:4001```

FAKE_DNS may be different for other OS.
For example for mac os - need run other image and ping host.docker.internal for see IP for host.
```bash
docker run --rm alpine ping host.docker.internal
```

https://docs.docker.com/docker-for-mac/networking/#/known-limitations-use-cases-and-workarounds

See tests/run_docker_boulder.sh for sed rate limits on own computer.
