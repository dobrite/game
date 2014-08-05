var connection = new WebSocket('ws://localhost:3000/sock/');

var ID;

var SCENE_WIDTH = 1920;
var SCENE_HEIGHT = 1024;

var TILE_HEIGHT =  32;
var TILE_WIDTH = 32;

var CHUNK;

var aspect = window.innerWidth / window.innerHeight;
//var camera = new THREE.OrthographicCamera( - d * aspect, d * aspect, d, - d, 1, 1000 );
camera = new THREE.OrthographicCamera(window.innerWidth / -2, window.innerWidth / 2, window.innerHeight / 2, window.innerHeight / -2, -5000, 5000);

var scene = new THREE.Scene();

//var axis = new THREE.AxisHelper(128);
//axis.position.set(0, 0, 0);
//scene.add(axis);
//
//var gridHelper = new THREE.GridHelper(1024, 32);
//scene.add(gridHelper);

//var ch = new THREE.CameraHelper(camera);
//scene.add(ch);

var renderer = new THREE.WebGLRenderer();

renderer.setSize(SCENE_WIDTH, SCENE_HEIGHT);
document.body.appendChild(renderer.domElement);

camera.position.set(200, 200, 200); // all components equal
camera.up = new THREE.Vector3(0, 1, 0);
camera.lookAt(new THREE.Vector3(0, 0, 0));

var ambientLight = new THREE.AmbientLight(0x10);
scene.add(ambientLight);

var directionalLight = new THREE.DirectionalLight(0xffffff);
directionalLight.position.x = 500;
directionalLight.position.y = 1000;
directionalLight.position.z = 500;
directionalLight.position.normalize();
scene.add( directionalLight );

var cubeGeo = new THREE.BoxGeometry(TILE_WIDTH, 32, TILE_HEIGHT);

var grassMesh = new THREE.MeshLambertMaterial({color: 0x80CF5A, shading: THREE.FlatShading});
var dirtMesh = new THREE.MeshLambertMaterial({color: 0x96712F, shading: THREE.FlatShading});
var waterMesh = new THREE.MeshLambertMaterial({color: 0x85b9bb, shading: THREE.FlatShading});

var itemGeo = new THREE.BoxGeometry(TILE_WIDTH/2, 16, TILE_HEIGHT/2);

var itemMesh = new THREE.MeshLambertMaterial({color: 0x5a6acf, shading: THREE.FlatShading});

// tiles
var grass = isoTile(grassMesh, TILE_WIDTH, TILE_HEIGHT);
var dirt = isoTile(dirtMesh, TILE_WIDTH, TILE_HEIGHT);
var water = isoTile(waterMesh, TILE_WIDTH, TILE_HEIGHT);

var player = isoItem(itemMesh, TILE_WIDTH/2, TILE_HEIGHT/2);

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

var empty = function(){};
var tileMethods = [grass, dirt, water, empty];
var itemMethods = [empty, empty, player];

function drawMap(terrain) {
  for (var i = 0, iL = terrain.length; i < iL; i++) {
   for (var j = 0, jL = terrain[i].length; j < jL; j++) {
      var x = j * TILE_WIDTH;
      var y = i * TILE_HEIGHT;

      var tileType = terrain[i][j];
      var drawTile = tileMethods[tileType];
      var cube = drawTile(x, y);
      if (TILES[i][j] === undefined) {
        TILES[i][j] = cube;
        scene.add(cube);
      } else {
        //nothing for now
      }
    }
  }
}

// TODO store by UUID
var PLAYER;

function drawItems(items) {
  for (var i =0, iL = items.length; i < iL; i++) {
    var item = items[i];
    var y = item.coords[0];
    var x = item.coords[1];
    var itemType = item.mt;

    x = (x * TILE_WIDTH) + TILE_WIDTH/4;
    y = (y * TILE_HEIGHT) + TILE_HEIGHT/4;

    var drawItem = itemMethods[itemType];
    if (PLAYER === undefined) {
      PLAYER = drawItem(x, y);
      scene.add(PLAYER);
    } else {
      // TODO this sucks
      PLAYER.position.x = x + 16;
      PLAYER.position.y = 32;
      PLAYER.position.z = y + 16;
    }
  }
}

connection.onopen = function () {
  console.log("connected");
};

connection.onerror = function (error) {
  console.log('WebSocket Error ' + error);
};

var TILES;

function initTiles(y, x) {
  var tiles = new Array(y);
  for (var i = 0; i < y; i++) {
    tiles[i] = new Array(x);
  }
  return tiles;
}

function handleGameConfigMessage(message) {
  TILES = initTiles(message.chunk_y, message.chunk_x);
  ID = message.id;
}

function handleGameWorldMessage(message) {
  drawMap(message.data.m);
  drawItems(message.data.i);
}

var messageToHandler = {
  "game:config": handleGameConfigMessage,
  "game:world": handleGameWorldMessage,
};

// Log messages from the server
connection.onmessage = function (e) {
  var message = JSON.parse(e.data);
  messageToHandler[message.event](message);
};

function buildMove(y, x) {
  return {
    id: ID,
    event: "game:move",
    data: {
      y: y,
      x: x
    },
  };
}

function moveAvatar(y, x) {
  connection.send(JSON.stringify(buildMove(y, x)));
}

function moveUp(e) {
  moveAvatar(-1, 0);
}
function moveDown(e) {
  moveAvatar(1, 0);
}
function moveLeft(e) {
  moveAvatar(0, -1);
}
function moveRight(e) {
  moveAvatar(0, 1);
}

// game loop optimized keyboard handling
kd.UP.down(moveUp);
kd.DOWN.down(moveDown);
kd.LEFT.down(moveLeft);
kd.RIGHT.down(moveRight);

function start () {
  function animate() {
    kd.tick();
    requestAnimationFrame(animate);
    //var timer = Date.now() * 0.0001;

    //camera.position.x = Math.cos(timer) * 200;
    //camera.position.z = Math.sin(timer) * 200;
    //camera.position.y = Math.tan(timer) * 200;
    //camera.lookAt(scene.position);
    renderer.render(scene, camera);
  }
  requestAnimationFrame(animate);
}

start();
