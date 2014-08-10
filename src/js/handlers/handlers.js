var configMessage = require('./configMessage'),
    worldMessage = require('./worldMessage'),
    itemMessage = require('./itemMessage');

var messageToHandler = {
  "game:config": configMessage,
  "game:los": worldMessage,
  "game:item": itemMessage,
};

module.exports = messageToHandler;
