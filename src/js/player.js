var player = {};

var initPlayer = function (data) {
  player.id = data.id;
  player.z = data.z;
  player.x = data.x;
  player.y = data.y;
  player.cz = data.cz;
  player.cx = data.cx;
  player.cy = data.cy;
};

var getId = function () {
  return player.id;
};

var getPlayer = function () {
  return player;
};

var updatePosition = function (z, x, y, cz, cx, cy) {
  player.z = z;
  player.x = x;
  player.y = y;
  player.cz = cz;
  player.cx = cx;
  player.cy = cy;
};

module.exports = {
  initPlayer: initPlayer,
  getId: getId,
  getPlayer: getPlayer,
  updatePosition: updatePosition
};
