const mongo = require('mongoose');

const host = process.env.MONGO_HOST || 'localhost';
const user = process.env.MONGO_USER || 'mongoadmin';
const password = process.env.MONGO_PASSWORD || 'secret';

const Person = mongo.model('Person', new mongo.Schema({
  name: 'string', age: 'number'
}));

module.exports = async () => {
  const url = `mongodb://${host}:27017/test`;
  console.log('Connecting to mongodb', url);
  await mongo.connect(url, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
    auth: {
      authSource: "admin",
      user, password
    }
  });
  return { Person }
};
