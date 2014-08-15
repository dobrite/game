var player = require('../player');

var spawnMessage = function (message) {
  player.initPlayer(message);
};

module.exports = spawnMessage;
