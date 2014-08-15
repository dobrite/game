var config = require('../config');
    world = require('../world');

var configMessage = function (message) {
  config.initConfig(message);
  world.initLos(message);
};

module.exports = configMessage;
