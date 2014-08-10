var scene = require('./scene');

var world;
var items = {};

// grass tiles
// #21313E
// #20575F
// #268073
// #53A976
// #98CF6F
// #EFEE69

var itemGeo = new THREE.BoxGeometry(config.TILE_WIDTH/2, 16, config.TILE_HEIGHT/2);
var cubeGeo = new THREE.BoxGeometry(config.TILE_WIDTH, 32, config.TILE_HEIGHT);

var buildMesh = function (color) {
  return new THREE.MeshLambertMaterial({
    color: color,
    shading: THREE.FlatShading,
  });
};

var cubeFactory = function (mesh) {
  return function(x, y) {
    var cube = new THREE.Mesh(cubeGeo, mesh);
    cube.position.x = x + 16;
    cube.position.z = y + 16;
    return cube;
  };
};

var itemFactory = function (mesh) {
  return function(x, y) {
    var item = new THREE.Mesh(itemGeo, mesh);
    item.position.x = x + 16;
    item.position.y = 32;
    item.position.z = y + 16;
    return item;
  };
};

var nothing = function(){};
var air = function(){};
var dirt = cubeFactory(buildMesh(0x96712F));
var grass = cubeFactory(buildMesh(0x80CF5A));
var water = cubeFactory(buildMesh(0x85b9bb));

var player = itemFactory(buildMesh(0x5a6acf));
var cow = itemFactory(buildMesh(0x614126));

var tileMethods = [nothing, air, dirt, grass, water];
var itemMethods = [nothing, nothing, nothing, nothing, nothing, player, cow];

var initWorld = function (data) {
  world = new Array(data.world_y);
  for (var i = 0; i < data.world_y; i ++) {
    world[i] = new Array(data.world_x);
  }
  for (var j = 0; j < data.world_y; j ++) {
    for (var k = 0; k < data.world_x; k++) {
      world[j][k] = initTiles(data.chunk_y, data.chunk_x);
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
  for (var y = 0; y < config.WORLD_Y; y++) {
    for (var x = 0; x < config.WORLD_X; x++) {
      renderChunk(y, x, chunks[y][x]);
    }
  }
};

var renderChunk = function (y, x, chunk) {
  var offset_y = chunk.coords[0] * config.CHUNK_Y * config.TILE_HEIGHT;
  var offset_x = chunk.coords[1] * config.CHUNK_X * config.TILE_WIDTH;

  for (var i = 0; i < config.CHUNK_Y; i++) {
    for (var j = 0; j < config.CHUNK_X; j++) {
      var cube_w = j * config.TILE_WIDTH;
      var cube_h = i * config.TILE_HEIGHT;

      var tileType = chunk.m[i][j];
      var drawTile = tileMethods[tileType];
      var cube = drawTile(cube_w + offset_x, cube_h + offset_y);

      if (world[y][x][i][j] === undefined) {
        world[y][x][i][j] = cube;
        scene.add(cube);
      } else {
        //nothing for now
      }
    }
  }
};

var renderItem = function (itemMsg) {
  // TODO this shouldn't know about itemMsg but w/e
  var y = itemMsg.world_coords.coords[0];
  var x = itemMsg.world_coords.coords[1];
  var cy = itemMsg.world_coords.chunk_coords[0];
  var cx = itemMsg.world_coords.chunk_coords[1];
  var offset_y = cy * config.CHUNK_Y * config.TILE_HEIGHT;
  var offset_x = cx * config.CHUNK_X * config.TILE_WIDTH;
  var itemType = itemMsg.material_type;

  //var cube_w = y * config.CHUNK_Y;
  //var cube_h = x * config.CHUNK_X;

  x = (x * config.TILE_WIDTH) + config.TILE_WIDTH/4 + offset_x;
  y = (y * config.TILE_HEIGHT) + config.TILE_HEIGHT/4 + offset_y;

  var drawItem = itemMethods[itemType];
  var item = items[itemMsg.id];
  if (item === undefined) {
    var itemMesh = drawItem(x, y);
    items[itemMsg.id] = itemMesh;
    scene.add(itemMesh);
  } else {
    item.position.x = x + 16;
    item.position.y = 32;
    item.position.z = y + 16;
  }
};

module.exports = {
  world: world,
  initWorld: initWorld,
  render: render,
  renderItem: renderItem,
};
