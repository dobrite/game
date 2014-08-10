var scene = require('./scene'),
    models = require('./models');

var los;
var items = {};

var initLos = function (data) {
  los = new Array(config.LOS_Y);
  for (var i = 0; i < config.LOS_Y; i ++) {
    los[i] = new Array(config.LOS_X);
  }
  for (var j = 0; j < config.LOS_Y; j ++) {
    for (var k = 0; k < config.LOS_X; k++) {
      los[j][k] = initTiles(config.CHUNK_Y, config.CHUNK_X);
    }
  }
};

var initTiles = function (y, x) {
  var tiles = new Array(y);
  for (var i = 0; i < y; i++) {
    tiles[i] = new Array(x);
  }
  return tiles;
};

var render = function (chunks) {
  for (var y = 0; y < config.LOS_Y; y++) {
    for (var x = 0; x < config.LOS_X; x++) {
      renderChunk(y, x, chunks[y][x]);
    }
  }
};

var offset = function (y, x) {
};

var renderChunk = function (y, x, chunk) {
  //y, x are los, i.e. los 3,3 [[0,0],[0,1]...[2,2]]
  var offset_y = (y - Math.floor(config.LOS_Y / 2)) * config.CHUNK_Y * config.TILE_HEIGHT;
  var offset_x = (x - Math.floor(config.LOS_X / 2)) * config.CHUNK_X * config.TILE_WIDTH;

  for (var i = 0; i < config.CHUNK_Y; i++) {
    for (var j = 0; j < config.CHUNK_X; j++) {
      var cube_w = j * config.TILE_WIDTH;
      var cube_h = i * config.TILE_HEIGHT;

      var cube = los[y][x][i][j];

      if (cube === undefined) {
        var tileType = chunk.m[i][j];
        var drawFunction = models.meshFunctions[tileType];
        cube = drawFunction();
        los[y][x][i][j] = cube;
        scene.add(cube);
      }

      cube.position.x = cube_w + offset_x + config.TILE_WIDTH / 2;
      cube.position.z = cube_h + offset_y + config.TILE_DEPTH / 2;
    }
  }
};

var renderItem = function (id, y, x, cy, cx, materialType) {
  //cy, cx are world coords
  var offset_y = (cy - Math.floor(config.LOS_Y / 2)) * config.CHUNK_Y * config.TILE_HEIGHT;
  var offset_x = (cx - Math.floor(config.LOS_X / 2)) * config.CHUNK_X * config.TILE_WIDTH;

  var cube_w = y * config.TILE_WIDTH;
  var cube_h = x * config.TILE_HEIGHT;

  y = cube_w + config.TILE_HEIGHT/4 + offset_y;
  x = cube_h + config.TILE_WIDTH/4 + offset_x;

  var item = items[id];

  if (item === undefined) {
    var drawFunction = models.meshFunctions[materialType];
    item = drawFunction();
    items[id] = item;
    scene.add(item);
  }

  item.position.x = x + config.TILE_WIDTH / 2;
  item.position.y = config.TILE_HEIGHT;
  item.position.z = y + config.TILE_DEPTH / 2;
};

module.exports = {
  initLos: initLos,
  render: render,
  renderItem: renderItem,
};
