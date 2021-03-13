
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const connect = require('./db');

const host = process.env.GRPC_HOST;
const port = process.env.GRPC_PORT;

const packageDefinition = protoLoader.loadSync('../protos/people.proto', {
  keepCase: true,
  longs: String,
  enums: String,
  arrays: true
});

const peopleProto = grpc.loadPackageDefinition(packageDefinition);

module.exports = async () => {

  const server = new grpc.Server();
  const { Person } = await connect();

  server.addService(peopleProto.PeopleService.service, {
    getAll: async (_, callback) => {
      const people = await Person.find();
      callback(null, { people });
    },
    remove: async (caller, callback) => {
      await Person.deleteOne({ _id: caller.request.id })
      callback(null, {});
    },
    insert: async (caller, callback) => {
      let person = new Person(caller.request);
      await person.save();
      callback(null, person);
    }
  });

  server.bindAsync(`${host}:${port}`,
    grpc.ServerCredentials.createInsecure(),
    (err) => {
      if (err) {
        console.log(err);
        return;
      }
      console.log(`gRPC server now running at http://${host}:${port}`);
      server.start();
    });
};
