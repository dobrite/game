var world = require('../world');

var worldMessage = function (message) {
  world.render(message.data);
};

module.exports = worldMessage;
