# gRPC on Node.js and MongoDB

To run project:

```
docker swarm init
```

```
docker build -t mongo-app .
```

```
docker stack deploy -c stack.yml mongo
```

Use a client, like BloomRPC, gRPCox, gRPC UI to inspect service.
Then, upload your protos dir to the client.

To inspect mongodb records go to localhost:8081
