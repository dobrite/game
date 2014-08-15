var player = require('./player');

var buildMove = function (z, x) {
  return {
    id: player.getId(),
    event: "game:move",
    data: {
      z: z,
      x: x
    },
  };
};

module.exports = {
  buildMove: buildMove,
};
