var scene = require('./scene'),
    models = require('./models');

var los;
var items = {};

var initLos = function (data) {
  los = new Array(config.LOS_Z);
  for (var i = 0; i < config.LOS_Z; i ++) {
    los[i] = new Array(config.LOS_X);
  }
  for (var j = 0; j < config.LOS_Z; j ++) {
    for (var k = 0; k < config.LOS_X; k++) {
      los[j][k] = initTiles(config.CHUNK_Z, config.CHUNK_X);
    }
  }
};

var initTiles = function (z, x) {
  var tiles = new Array(z);
  for (var i = 0; i < z; i++) {
    tiles[i] = new Array(x);
  }
  return tiles;
};

var renderAll = function (chunks) {
  for (var z = 0; z < config.LOS_Z; z++) {
    for (var x = 0; x < config.LOS_X; x++) {
      renderChunk(z, x, chunks[z][x]);
    }
  }
};

var offset = function (z, x) {
};

var renderChunk = function (z, x, chunk) {
  //y, x are los, i.e. los 3,3 [[0,0],[0,1]...[2,2]]
  var offsetZ = (z - Math.floor(config.LOS_Z / 2)) * config.CHUNK_Z * config.TILE_HEIGHT;
  var offsetX = (x - Math.floor(config.LOS_X / 2)) * config.CHUNK_X * config.TILE_WIDTH;

  for (var i = 0; i < config.CHUNK_Z; i++) {
    for (var j = 0; j < config.CHUNK_X; j++) {
      var cubeZ = i * config.TILE_HEIGHT;
      var cubeX = j * config.TILE_WIDTH;

      var sceneZ = cubeZ + offsetZ;
      var sceneX = cubeX + offsetX;

      var cube = los[z][x][i][j];

      if (cube === undefined) {
        var tileType = chunk.m[i][j];
        var drawFunction = models.meshFunctions[tileType];
        cube = drawFunction();
        los[z][x][i][j] = cube;
        scene.add(cube);
      }

      cube.position.z = sceneZ + config.TILE_DEPTH / 2;
      cube.position.x = sceneX + config.TILE_WIDTH / 2;
    }
  }
};

var renderItem = function (id, z, x, cz, cx, materialType) {
  //cz, cx are world coords
  var offsetZ = (cz - Math.floor(config.LOS_Z / 2)) * config.CHUNK_Z * config.TILE_HEIGHT;
  var offsetX = (cx - Math.floor(config.LOS_X / 2)) * config.CHUNK_X * config.TILE_WIDTH;

  var cubeZ = z * config.TILE_WIDTH;
  var cubeX = x * config.TILE_HEIGHT;

  var sceneZ = cubeZ + offsetZ + config.TILE_HEIGHT / 4;
  var sceneX = cubeX + offsetX + config.TILE_WIDTH / 4;

  var item = items[id];

  if (item === undefined) {
    var drawFunction = models.meshFunctions[materialType];
    item = drawFunction();
    items[id] = item;
    scene.add(item);
  }

  item.position.z = sceneZ + config.TILE_DEPTH / 2;
  item.position.x = sceneX + config.TILE_WIDTH / 2;
  item.position.y = config.TILE_HEIGHT;
};

module.exports = {
  initLos: initLos,
  renderAll: renderAll,
  renderItem: renderItem,
};
