
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const connect = require('./db');

const packageDefinition = protoLoader.loadSync('./protos/people.proto', {
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

  await server.bindAsync("127.0.0.1:30043",
    grpc.ServerCredentials.createInsecure(),
    (err, port) => {
      if (err) {
        console.log(err);
        return;
      }
      console.log(`Server running at http://127.0.0.1:${port}`);
      server.start();
    });
};
