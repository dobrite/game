var configMessage = require('./configMessage'),
    worldMessage = require('./worldMessage'),
    itemMessage = require('./itemMessage'),
    spawnMessage = require('./spawnMessage');

// TODO create a decorator or register on import
var messageToHandler = {
  "game:config": configMessage,
  "game:los": worldMessage,
  "game:item": itemMessage,
  "game:spawn": spawnMessage,
};

module.exports = messageToHandler;
