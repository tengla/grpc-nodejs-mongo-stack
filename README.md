# gRPC on Node.js and MongoDB

To run project:

```
docker swarm init
```

```
docker build -t grpc-server .
```

```
docker stack deploy -c stack.yml grpc-test
```

Use a client, like BloomRPC, gRPCox, gRPC UI to inspect service.
Then, upload your protos dir to the client.

To inspect mongodb records go to: [http://localhost:8081/](http://localhost:8081/)

When you're done:

```
docker stack rm grpc-test
```
