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
var pig = itemFactory(buildMesh(0xFCD7DE));

var tileFunctions = [nothing, air, dirt, grass, water];
var itemFunctions = [nothing, nothing, nothing, nothing, nothing, player, cow, pig];

module.exports = {
  tileFunctions: tileFunctions,
  itemFunctions: itemFunctions,
};
