var connection = new WebSocket('ws://localhost:3000/sock/');

var CHUNK_Y;
var CHUNK_X;
var ID;

var STAGE_WIDTH = 1920;
var STAGE_HEIGHT = 1024;

var TILE_HEIGHT =  32;
var TILE_WIDTH = 32;

var CHUNK;

var stage = new PIXI.Stage(0xEEFFFF);
var renderer = PIXI.autoDetectRenderer(STAGE_WIDTH, STAGE_HEIGHT);

document.body.appendChild(renderer.view);

var graphics = new PIXI.Graphics();
stage.addChild(graphics);

// An iso tile is twice as wide as it is tall (2w x h)
function isoTile(backgroundColor, borderColor, w, h) {
  var h_2 = h/2;

  return function(x, y) {
    graphics.beginFill(backgroundColor);
    graphics.lineStyle(1, borderColor, 1);
    graphics.moveTo(x, y);
    graphics.lineTo(x + w, y + h_2);
    graphics.lineTo(x, y + h);
    graphics.lineTo(x - w, y + h_2);
    graphics.lineTo(x , y);
    graphics.endFill();
  };
}

function isoItem(backgroundColor, borderColor, w, h) {
  var h_2 = h/2;

  return function(x, y) {
    graphics.beginFill(backgroundColor);
    graphics.lineStyle(1, borderColor, 1);
    graphics.moveTo(x, y);
    graphics.lineTo(x + w, y + h_2);
    graphics.lineTo(x, y + h);
    graphics.lineTo(x - w, y + h_2);
    graphics.lineTo(x , y);
    graphics.endFill();
  };
}

// tiles
var grass = isoTile(0x80CF5A, 0x339900, TILE_WIDTH, TILE_HEIGHT);
var dirt = isoTile(0x96712F, 0x403014, TILE_WIDTH, TILE_HEIGHT);
var water = isoTile(0x85b9bb, 0x476263, TILE_WIDTH, TILE_HEIGHT);
var player = isoItem(0x5a6acf, 0x2b40cc, TILE_WIDTH/2, TILE_HEIGHT/2);
var empty = function(){};
var tileMethods = [grass, dirt, water, empty];
var itemMethods = [empty, player];

function drawMap(terrain, xOffset) {
    var tileType, x, y, isoX, isoY, idx;

    for (var i = 0, iL = terrain.length; i < iL; i++) {
        for (var j = 0, jL = terrain[i].length; j < jL; j++) {
            // cartesian 2D coordinate
            x = j * TILE_WIDTH;
            y = i * TILE_HEIGHT;

            // iso coordinate
            isoX = x - y;
            isoY = (x + y) / 2;

            tileType = terrain[i][j];
            drawTile = tileMethods[tileType];
            drawTile(xOffset + isoX, isoY);
        }
    }
}

var items = [
  [[4, 8], 1],
];

function drawItems(items, xOffset) {
    var itemType, x, y, isoX, isoY, idx, coords, item;

    for (var i =0, iL = items.length; i < iL; i++) {
      item = items[i];
      coords = item[0];
      y = coords[0];
      x = coords[1];
      itemType = item[1];

      // cartesian 2D coordinate
      x = (x * TILE_WIDTH) + TILE_WIDTH/4;
      y = (y * TILE_HEIGHT) + TILE_HEIGHT/4;

      // iso coordinate
      isoX = x - y;
      isoY = (x + y) / 2;

      drawItem = itemMethods[itemType];
      drawItem(xOffset + isoX, isoY);
    }
}

// When the connection is open, send some data to the server
connection.onopen = function () {
  //connection.send('Ping'); // Send the message 'Ping' to the server
};

// Log errors
connection.onerror = function (error) {
  console.log('WebSocket Error ' + error);
};

function handleGameConfigMessage(message) {
  CHUNK_Y = message.chunk_y;
  CHUNK_X = message.chunk_x;
  ID = message.id;
}

function dtodd(arr) {
  var newArr = [];
  while(arr.length) newArr.push(arr.splice(0,CHUNK_X));
  drawMap(newArr, STAGE_WIDTH / 2);
}

function handleGameWorldMessage(message) {
  dtodd(message.data.m);
}

function handleGameItemsMessage(message) {
  drawItems(message.items, STAGE_WIDTH / 2);
}

// Log messages from the server
connection.onmessage = function (e) {
  var message = JSON.parse(e.data);
  console.log(message);
  if(message.event == "game:config") {
    handleGameConfigMessage(message);
  }
  if(message.event == "game:world") {
    handleGameWorldMessage(message);
  }
  if(message.event == "game:items") {
    handleGameItemsMessage(message);
  }
};

function buildMove(y, x) {
  return {id: ID, event: "game:move", data: {y: y, x: x}};
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
    requestAnimFrame(animate);
    renderer.render(stage);
  }
  requestAnimFrame(animate);
}

start();
