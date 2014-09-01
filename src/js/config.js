var initConfig = function (data) {
  window.config.LOS_Z = data.los_z; // how many chunks around player to render
  window.config.LOS_X = data.los_x; // how many chunks around player to render
  window.config.LOS_Y = data.world_y / data.chunk_y; // how many chunks around player to render
  window.config.CHUNK_Z = data.chunk_z; // tiles per chunk
  window.config.CHUNK_X = data.chunk_x; // tiles per chunk
  window.config.CHUNK_Y = data.chunk_y; // tiles per chunk
  console.log(config.LOS_Y);
  console.log(data.world_y, data.chunk_y);
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
