var ID, LOS_X, LOS_Y, CHUNK_X, CHUNK_Y;

var initConfig = function (data) {
  window.config.ID = data.id;
  window.config.LOS_X = data.los_x; // how many chunks around player to render
  window.config.LOS_Y = data.los_y; // how many chunks around player to render
  window.config.CHUNK_X = data.chunk_x; // tiles per chunk
  window.config.CHUNK_Y = data.chunk_y; // tiles per chunk
};

var SCENE_WIDTH = 1920;
var SCENE_HEIGHT = 1024;

var TILE_HEIGHT =  32;
var TILE_WIDTH = 32;

window.config = {
  SCENE_WIDTH: SCENE_WIDTH,
  SCENE_HEIGHT: SCENE_HEIGHT,
  TILE_WIDTH: TILE_WIDTH,
  TILE_HEIGHT: TILE_HEIGHT,
};

module.exports = {
  initConfig: initConfig,
};
