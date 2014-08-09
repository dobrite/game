var config = require('../config');
world = require('../world');

var configMessage = function (message) {
  config.initConfig(message);
  world.initWorld(message);
};

module.exports = configMessage;
