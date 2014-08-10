var ID, LOS_X, LOS_Y, CHUNK_X, CHUNK_Y;

var initConfig = function (data) {
  window.config.ID = data.id;
  window.config.LOS_X = data.los_x; // how many chunks around player to render
  window.config.LOS_Y = data.los_y; // how many chunks around player to render
  window.config.CHUNK_X = data.chunk_x; // tiles per chunk
  window.config.CHUNK_Y = data.chunk_y; // tiles per chunk
};

window.config = {
  SCENE_WIDTH: 1640,
  SCENE_HEIGHT: 1024,
  TILE_WIDTH: 32, // x
  TILE_HEIGHT: 32, // y
  TILE_DEPTH: 32, // z (right now y) north is -Z
};

module.exports = {
  initConfig: initConfig,
};
