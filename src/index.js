
require('dotenv').config();
const serve = require('./server');

(async () => {
  try {
    await serve();
  } catch (err) {
    console.error(err);
  }
})();
