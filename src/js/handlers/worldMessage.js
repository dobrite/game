var world = require('../world');

var worldMessage = function (message) {
  world.renderAll(message.data);
};

module.exports = worldMessage;
