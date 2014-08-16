var scene = require('./scene'),
    models = require('./models');

var los;
var items = {};

var initLos = function (data) {
  console.log(config.LOS_Z * config.CHUNK_Z);
  los = new Array(config.LOS_Z * config.CHUNK_Z);
  for (var z = 0; z < config.LOS_Z * config.CHUNK_Z; z ++) {
    los[z] = new Array(config.LOS_X * config.CHUNK_X);
    for (var y = 0; y < config.LOS_Y * config.CHUNK_Y; y ++) {
      los[z][y] = new Array(config.LOS_Y * config.CHUNK_Y);
    }
  }
};

var renderAll = function (chunks) {
  for (var z = 0; z < config.LOS_Z; z++) {
    for (var x = 0; x < config.LOS_X; x++) {
      //var chunk = chunks[z][x];
      //var wcZ = chunk.coords[0];
      //var wcX = chunk.coords[1];
      renderChunk(z, x, chunks[z][x]);
    }
  }
};

var render = function () {
};

var renderChunk = function (z, x, chunk) {
  //z, x are los, i.e. los 3,3 [[0,0],[0,1]...[2,2]]
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
