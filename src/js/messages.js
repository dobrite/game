var buildMove = function (z, x) {
  return {
    id: config.ID,
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
