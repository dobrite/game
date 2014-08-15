var world = require('../world'),
    config = require('../config'),
    ui = require('../ui'),
    player = require('../player');

var z, x, cz, cx;

var itemMessage = function (message) {
  z = message.worldCoords.coords[0];
  x = message.worldCoords.coords[1];
  cz = message.worldCoords.chunkCoords[0];
  cx = message.worldCoords.chunkCoords[1];

  if(player.getId() === message.id) {
    ui.updateUi(z, x, cz, cx);
    player.updatePosition(z, x, cz, cx);
  }

  world.renderItem(message.id, z, x, cz, cx, message.materialType);
};

module.exports = itemMessage;
