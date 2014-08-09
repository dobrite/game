var configMessage = require('./configMessage'),
    worldMessage = require('./worldMessage'),
    itemMessage = require('./itemMessage');

var messageToHandler = {
  "game:config": configMessage,
  "game:world": worldMessage,
  "game:item": itemMessage,
};

module.exports = messageToHandler;
