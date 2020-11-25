# Local Discovery

A simple API which can be used to discover agents in a local network.

## Using the API

All endpoints only expose agents in the same network as the request origin, based on the remote client IP address.

### Agent Schema

| Attribute  | Description |
|------------|-------------|
| `name`     | The name of the agent application. This will likely be the same for all agents.
| `version`  | The version of the agent application.
| `address`  | The local address (ip and port) at which the agent is running.
| `platform` | The platform the agent is running on, e.g. `windows`, `linux` or `mac`.
| `hostname` | The name of the host the agent is running on.

### List the Agents

Send a `GET` request to `/api/agents`. This will return a JSON list of known agent information:

```json
[
  {
    "name": "agent",
    "version": "1.0.0",
    "address": "192.168.0.110:4000",
    "platform": "windows",
    "hostname": "my-pc"
  }
]
```

### Register an Agent

Send a `POST` request to `/api/agents`, with the agent information as JSON in the request body:

```json
{
  "name": "agent",
  "version": "1.0.0",
  "address": "192.168.0.110:4000",
  "platform": "windows",
  "hostname": "my-pc"
}
```

This will return the list of known agents in the local network (see "List the Agents" above).

If the agent address is the same as the address of a known agent, the known agent will be replaced with the new one.

## Configuration

The discovery server can be configured through environment variables.

| Variable          | Default | Description |
|-------------------|---------|-------------|
| `LOG_LEVEL`       | `INFO`  | The log level (`DEBUG`, `INFO`, `WARN`, `ERROR` or `OFF`).
| `PORT`            | 5000    | The port the server listens on.
| `KEEP_AGENT_TIME` | 600     | The time (in seconds) to keep agents in the registry. 
