# Local Discovery

[![Docker Image Version](https://img.shields.io/docker/v/lucascorpion/local-discovery?sort=semver)](https://hub.docker.com/r/lucascorpion/local-discovery)
[![Docker Image Size](https://img.shields.io/docker/image-size/lucascorpion/local-discovery?sort=semver)](https://hub.docker.com/r/lucascorpion/local-discovery)
[![Docker Pulls](https://img.shields.io/docker/pulls/lucascorpion/local-discovery)](https://hub.docker.com/r/lucascorpion/local-discovery)

A simple JSON API which can be used to discover agents in a local network.

## Using the API

All endpoints only expose agents in the same network as the request origin, based on the remote client IP address.

By default the API will start on port 4000.

### Agent Schema

| Attribute      | Description |
|----------------|-------------|
| `name`         | The name of the agent application.
| `localAddress` | The local address at which the agent is running.
| `info`         | A freeform object containing application-specific info about the agent.

### List the Agents

Send a `GET` request to `/api/agents`.
This will return a list of known agent information:

```json
[
  {
    "name": "agent",
    "localAddress": "192.168.0.110:4000",
    "info": {}
  }
]
```

### Register an Agent

Send a `POST` request to `/api/agents`, with the agent information in the request body:

```json
{
  "name": "agent",
  "localAddress": "192.168.0.110:4000",
  "info": {}
}
```

This will return the newly created agent.

Note: if the agent address is the same as the address of a known agent, the known agent will be replaced with the new one.

### Remove an Agent

Send a `DELETE` request to `/api/agents`, with the agent information in the request body:

```json
{
  "name": "agent",
  "localAddress": "192.168.0.110:4000"
}
```

This will delete the agent whose `name` and `localAddress` match this info.
If no such agent is found, nothing happens.
This will return an empty.
