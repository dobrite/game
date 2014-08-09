var ID, WORLD_X, WORLD_Y, CHUNK_X, CHUNK_Y;

var initConfig = function (data) {
  window.config.ID = data.id;
  window.config.WORLD_X = data.world_x;
  window.config.WORLD_Y = data.world_y;
  window.config.CHUNK_X = data.chunk_x;
  window.config.CHUNK_Y = data.chunk_y;
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
