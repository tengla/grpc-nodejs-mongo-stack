
require('dotenv').config();
const serve = require('./server');

(async () => {
  await serve();
})();
