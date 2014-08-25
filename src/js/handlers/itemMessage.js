var world = require('../world'),
    config = require('../config'),
    ui = require('../ui'),
    player = require('../player');

var z, x, cz, cx;

var itemMessage = function (message) {
  z = message.wc.c[0];
  x = message.wc.c[1];
  y = message.wc.c[2];
  cz = message.wc.cc[0];
  cx = message.wc.cc[1];
  cy = message.wc.cc[2];

  if(player.getId() === message.id) {
    ui.updateUi(z, x, y, cz, cx, cy);
    player.updatePosition(z, x, y, cz, cx, cy);
  }

  world.renderItem(message.id, z, x, cz, cx, message.materialType);
};

module.exports = itemMessage;
