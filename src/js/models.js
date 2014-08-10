// grass tiles
// #21313E
// #20575F
// #268073
// #53A976
// #98CF6F
// #EFEE69

var nothing = function(){};
var air = function(){};

var meshFunctions = [nothing, air];

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

var meshFactory = function (geo, mesh) {
  var meshFunction = function() {
    return new THREE.Mesh(geo, mesh);
  };
  meshFunctions.push(meshFunction);
  return meshFunction;
};

var dirt = meshFactory(cubeGeo, buildMesh(0x96712F));
var grass = meshFactory(cubeGeo, buildMesh(0x80CF5A));
var water = meshFactory(cubeGeo, buildMesh(0x85b9bb));
var stone = meshFactory(itemGeo, buildMesh(0xCCCCCC));
var player = meshFactory(itemGeo, buildMesh(0x5a6acf));
var cow = meshFactory(itemGeo, buildMesh(0x614126));
var pig = meshFactory(itemGeo, buildMesh(0xFCD7DE));

module.exports = {
  meshFunctions: meshFunctions,
};
