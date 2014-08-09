var scene = require('./scene');

var world;

// tiles
var itemGeo = new THREE.BoxGeometry(config.TILE_WIDTH/2, 16, config.TILE_HEIGHT/2);
var cubeGeo = new THREE.BoxGeometry(config.TILE_WIDTH, 32, config.TILE_HEIGHT);
var itemMesh = new THREE.MeshLambertMaterial({color: 0x5a6acf, shading: THREE.FlatShading});
var player = isoItem(itemMesh, config.TILE_WIDTH/2, config.TILE_HEIGHT/2);

var grassMesh = new THREE.MeshLambertMaterial({color: 0x80CF5A, shading: THREE.FlatShading});
var dirtMesh = new THREE.MeshLambertMaterial({color: 0x96712F, shading: THREE.FlatShading});
var waterMesh = new THREE.MeshLambertMaterial({color: 0x85b9bb, shading: THREE.FlatShading});

function isoTile(mesh, w, h) {
  return function(x, y) {
    var cube = new THREE.Mesh(cubeGeo, mesh);
    cube.position.x = x + 16;
    cube.position.z = y + 16;
    return cube;
  };
}

function isoItem(mesh, w, h) {
  return function(x, y) {
    var item = new THREE.Mesh(itemGeo, mesh);
    item.position.x = x + 16;
    item.position.y = 32;
    item.position.z = y + 16;
    return item;
  };
}

var grass = isoTile(grassMesh, config.TILE_WIDTH, config.TILE_HEIGHT);
var dirt = isoTile(dirtMesh, config.TILE_WIDTH, config.TILE_HEIGHT);
var water = isoTile(waterMesh, config.TILE_WIDTH, config.TILE_HEIGHT);

var empty = function(){};
var tileMethods = [grass, dirt, water, empty];
var itemMethods = [empty, empty, player];

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

var renderChunk = function (y, x, chunk) {
  var chunk_coord_x = chunk.coords[0];
  var chunk_coord_y = chunk.coords[1];

  var offset_y = chunk_coord_y * config.CHUNK_Y * config.TILE_HEIGHT;
  var offset_x = chunk_coord_x * config.CHUNK_X * config.TILE_WIDTH;

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

var render = function (chunks) {
  for (var y = 0; y < config.WORLD_Y; y++) {
    for (var x = 0; x < config.WORLD_X; x++) {
      renderChunk(y, x, chunks[y][x]);
    }
  }
};

function renderItems(items) {
  for (var i = 0, iL = items.length; i < iL; i++) {
    var item = items[i];
    var y = item.coords[0];
    var x = item.coords[1];
    var itemType = item.mt;

    x = (x * TILE_WIDTH) + TILE_WIDTH/4;
    y = (y * TILE_HEIGHT) + TILE_HEIGHT/4;

    var drawItem = itemMethods[itemType];
    if (window.PLAYER === undefined) {
      window.PLAYER = drawItem(x, y);
      scene.add(window.PLAYER);
    } else {
      // TODO this sucks
      window.PLAYER.position.x = x + 16;
      window.PLAYER.position.y = 32;
      window.PLAYER.position.z = y + 16;
    }
  }
}


module.exports = {
  world: world,
  initWorld: initWorld,
  render: render,
  renderItems: renderItems,
};
