# gRPC with Node.js, golang and MongoDB

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

To inspect mongodb records go to: [localhost](http://127.0.0.1:8081/)

When you're done:

```
docker stack rm grpc-test
```

For more info about Node.js and gRPC, refer to https://blog.logrocket.com/creating-a-crud-api-with-node-express-and-grpc/
