var configMessage = require('./configMessage'),
worldMessage = require('./worldMessage');

var messageToHandler = {
  "game:config": configMessage,
  "game:world": worldMessage,
};

module.exports = messageToHandler;
