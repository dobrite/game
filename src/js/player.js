var player = {};

var initPlayer = function(data) {
  player.id = data.id;
  player.z = data.z;
  player.x = data.x;
  player.cz = data.cz;
  player.cx = data.cx;
};

var getId = function() {
  return player.id;
};

var updatePosition = function (z, x, cz, cx) {
  player.z = z;
  player.x = x;
  player.cz = cz;
  player.cx = cx;
};

module.exports = {
  initPlayer: initPlayer,
  getId: getId,
  updatePosition: updatePosition
};
