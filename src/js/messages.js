var config = require('./config');

var buildMove = function (y, x) {
  return {
    id: config.ID,
    event: "game:move",
    data: {
      y: y,
      x: x
    },
  };
};

module.exports = {
  buildMove: buildMove,
};
