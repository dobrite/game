var scene = require('./scene');

var los;
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

var renderChunk = function (y, x, chunk) {
  //y, x are los, i.e. los 3,3 [[0,0],[0,1]...[2,2]]
  var offset_y = (y - Math.floor(config.LOS_Y / 2)) * config.CHUNK_Y * config.TILE_HEIGHT;
  var offset_x = (x - Math.floor(config.LOS_X / 2)) * config.CHUNK_X * config.TILE_WIDTH;

  for (var i = 0; i < config.CHUNK_Y; i++) {
    for (var j = 0; j < config.CHUNK_X; j++) {
      var cube_w = j * config.TILE_WIDTH;
      var cube_h = i * config.TILE_HEIGHT;

      var tileType = chunk.m[i][j];
      var drawTile = tileMethods[tileType];
      var cube = drawTile(cube_w + offset_x, cube_h + offset_y);

      if (los[y][x][i][j] === undefined) {
        los[y][x][i][j] = cube;
        scene.add(cube);
      } else {
        //nothing for now
      }
    }
  }
};

var renderItem = function (id, y, x, cy, cx, material_type) {
  //var offset_y = chunk.coords[0] * config.CHUNK_Y * config.TILE_HEIGHT;
  //var offset_y = (y - Math.floor(config.LOS_Y / 2)) * config.CHUNK_Y * config.TILE_HEIGHT;

  var offset_y = (cy - Math.floor(config.LOS_Y / 2)) * config.CHUNK_Y * config.TILE_HEIGHT;
  var offset_x = (cx - Math.floor(config.LOS_X / 2)) * config.CHUNK_X * config.TILE_WIDTH;
  var itemType = material_type;

  var cube_w = y * config.TILE_WIDTH;
  var cube_h = x * config.TILE_HEIGHT;

  y = cube_w + config.TILE_HEIGHT/4 + offset_y;
  x = cube_h + config.TILE_WIDTH/4 + offset_x;

  var drawItem = itemMethods[itemType];
  var item = items[id];
  if (item === undefined) {
    var itemMesh = drawItem(x, y);
    items[id] = itemMesh;
    scene.add(itemMesh);
  } else {
    item.position.x = x + 16;
    item.position.y = 32;
    item.position.z = y + 16;
  }
};

module.exports = {
  initLos: initLos,
  render: render,
  renderItem: renderItem,
};
