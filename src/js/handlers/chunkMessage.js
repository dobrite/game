var world = require('../world');

var chunkMessage = function (message) {
  world.renderChunk(message.coords, message.materials);
};

module.exports = chunkMessage;
