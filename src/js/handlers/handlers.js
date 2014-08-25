var configMessage = require('./configMessage'),
    chunkMessage = require('./chunkMessage'),
    itemMessage = require('./itemMessage'),
    spawnMessage = require('./spawnMessage');

// TODO create a decorator or register on import
var messageToHandler = {
  "game:config": configMessage,
  "game:chunk": chunkMessage,
  "game:item": itemMessage,
  "game:spawn": spawnMessage,
};

module.exports = messageToHandler;
