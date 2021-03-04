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

Use a client, like BloomRPC, gRPCox, gRPC UI or Protoman to inspect api.
Then, upload your protos dir to the client.

To inspect mongodb records go to: [http://127.0.0.1:8081/](localhost)

When you're done:

```
docker stack rm grpc-test
```
