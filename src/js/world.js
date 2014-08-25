var scene = require('./scene'),
    models = require('./models'),
    player = require('./player');

var los;
var items = {};

var initLos = function () {
  los = new Array(config.LOS_Z);
  for (var z = 0; z < config.LOS_Z; z++) {
    los[z] = new Array(config.LOS_X);
    for (var x = 0; x < config.LOS_X; x++) {
      los[z][x] = new Array(config.CHUNK_Z);
      for (var cz = 0; cz < config.CHUNK_Z; cz++) {
        los[z][x][cz] = new Array(config.CHUNK_X);
      }
    }
  }
};

var renderChunk = function (coords, chunk) {
  var cz = coords[0];
  var cx = coords[1];

  var p = player.getPlayer();

  var playerRelativeCoordsZ = cz - p.cz;
  var playerRelativeCoordsX = cx - p.cx;

  var halfLosZ = Math.floor(config.LOS_Z / 2);
  var halfLosX = Math.floor(config.LOS_X / 2);
  
  var losArrayCoordsZ = playerRelativeCoordsZ + halfLosZ;
  var losArrayCoordsX = playerRelativeCoordsX + halfLosX;

  var offsetChunkZ = cz - halfLosZ;
  var offsetChunkX = cx - halfLosX;

  var offsetZ = offsetChunkZ * config.CHUNK_Z * config.TILE_HEIGHT;
  var offsetX = offsetChunkX * config.CHUNK_X * config.TILE_WIDTH;

  for (var i = 0; i < config.CHUNK_Z; i++) {
    for (var j = 0; j < config.CHUNK_X; j++) {
      var cubeZ = i * config.TILE_HEIGHT;
      var cubeX = j * config.TILE_WIDTH;

      var sceneZ = cubeZ + offsetZ;
      var sceneX = cubeX + offsetX;

      var cube = los[losArrayCoordsZ][losArrayCoordsX][i][j];

      if (cube === undefined) {
        var tileType = chunk[i][j];
        var drawFunction = models.meshFunctions[tileType];
        cube = drawFunction();
        los[losArrayCoordsZ][losArrayCoordsX][i][j] = cube;
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
  renderChunk: renderChunk,
  renderItem: renderItem,
};
