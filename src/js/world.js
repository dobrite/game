var scene = require('./scene'),
    models = require('./models'),
    player = require('./player');

var los;
var items = {};

// lol
var initLos = function () {
  los = new Array(config.LOS_Z);
  for (var z = 0; z < config.LOS_Z; z++) {
    los[z] = new Array(config.LOS_X);
    for (var x = 0; x < config.LOS_X; x++) {
      los[z][x] = new Array(config.LOS_Y);
      for (var y = 0; y < config.LOS_Y; y++) {
        los[z][x][y] = new Array(config.CHUNK_Z);
        for (var cz = 0; cz < config.CHUNK_Z; cz++) {
          los[z][x][y][cz] = new Array(config.CHUNK_X);
          for (var cx = 0; cx < config.CHUNK_X; cx++) {
            los[z][x][y][cz][cx] = new Array(config.CHUNK_Y);
          }
        }
      }
    }
  }
};

var renderChunk = function (coords, chunk) {
  var cz = coords[0];
  var cx = coords[1];
  var cy = coords[2];

  var p = player.getPlayer();

  var playerRelativeCoordsZ = cz - p.cz;
  var playerRelativeCoordsX = cx - p.cx;

  var halfLosZ = Math.floor(config.LOS_Z / 2);
  var halfLosX = Math.floor(config.LOS_X / 2);

  var losArrayCoordsZ = playerRelativeCoordsZ + halfLosZ;
  var losArrayCoordsX = playerRelativeCoordsX + halfLosX;

  var offsetChunkZ = cz - halfLosZ;
  var offsetChunkX = cx - halfLosX;

  var offsetZ = offsetChunkZ * config.CHUNK_Z * config.TILE_DEPTH;
  var offsetX = offsetChunkX * config.CHUNK_X * config.TILE_WIDTH;
  var offsetY = cy * config.CHUNK_Y * config.TILE_HEIGHT;

  for (var k = 0; k < config.CHUNK_Y; k++) {
    for (var i = 0; i < config.CHUNK_Z; i++) {
      for (var j = 0; j < config.CHUNK_X; j++) {
        var cubeZ = i * config.TILE_DEPTH;
        var cubeX = j * config.TILE_WIDTH;
        var cubeY = k * config.TILE_HEIGHT;

        var sceneZ = cubeZ + offsetZ;
        var sceneX = cubeX + offsetX;
        var sceneY = cubeY + offsetY;

        // really needs to be worldHeight / chunkY
        var cube = los[losArrayCoordsZ][losArrayCoordsX][cy][k][i][j];

        if (cube === undefined) {
          var tileType = chunk[k][i][j];
          var drawFunction = models.meshFunctions[tileType];
          cube = drawFunction();
          if (cube !== undefined) {
            los[losArrayCoordsZ][losArrayCoordsX][cy][k][i][j] = cube;
            scene.add(cube);
          }
        }

        if (cube !== undefined) {
          cube.position.z = sceneZ + config.TILE_DEPTH / 2;
          cube.position.x = sceneX + config.TILE_WIDTH / 2;
          cube.position.y = sceneY;
        }
      }
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
