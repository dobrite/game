var world = require('../world');

var itemMessage = function (message) {
  world.renderItem(
    message.id,
    message.worldCoords.coords[0],
    message.worldCoords.coords[1],
    message.worldCoords.chunkCoords[0],
    message.worldCoords.chunkCoords[1],
    message.materialType
  );
};

module.exports = itemMessage;
