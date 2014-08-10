// grass tiles
// #21313E
// #20575F
// #268073
// #53A976
// #98CF6F
// #EFEE69

var itemGeo = new THREE.BoxGeometry(
  config.TILE_WIDTH / 2, // x
  config.TILE_HEIGHT / 2, // y
  config.TILE_DEPTH / 2 // z
);

var cubeGeo = new THREE.BoxGeometry(
  config.TILE_WIDTH, // x
  config.TILE_HEIGHT, // y
  config.TILE_DEPTH // z
);

var buildMesh = function (color) {
  return new THREE.MeshLambertMaterial({
    color: color,
    shading: THREE.FlatShading,
  });
};

var cubeFactory = function (cubeGeo, mesh) {
  return function(x, y) {
    var cube = new THREE.Mesh(cubeGeo, mesh);
    cube.position.x = x + config.TILE_WIDTH / 2;
    cube.position.z = y + config.TILE_DEPTH / 2;
    return cube;
  };
};

var itemFactory = function (itemGeo, mesh) {
  return function(x, y) {
    var item = new THREE.Mesh(itemGeo, mesh);
    item.position.x = x + config.TILE_WIDTH / 2;
    item.position.y = config.TILE_HEIGHT;
    item.position.z = y + config.TILE_DEPTH / 2;
    return item;
  };
};

var nothing = function(){};
var air = function(){};
var dirt = cubeFactory(cubeGeo, buildMesh(0x96712F));
var grass = cubeFactory(cubeGeo, buildMesh(0x80CF5A));
var water = cubeFactory(cubeGeo, buildMesh(0x85b9bb));

var player = itemFactory(itemGeo, buildMesh(0x5a6acf));
var cow = itemFactory(itemGeo, buildMesh(0x614126));
var pig = itemFactory(itemGeo, buildMesh(0xFCD7DE));

var tileFunctions = [nothing, air, dirt, grass, water];
var itemFunctions = [nothing, nothing, nothing, nothing, nothing, player, cow, pig];

module.exports = {
  tileFunctions: tileFunctions,
  itemFunctions: itemFunctions,
};
