var connection = new WebSocket('ws://localhost:3000/sock/');

// When the connection is open, send some data to the server
connection.onopen = function () {
  //connection.send('Ping'); // Send the message 'Ping' to the server
};

// Log errors
connection.onerror = function (error) {
  console.log('WebSocket Error ' + error);
};

// Log messages from the server
connection.onmessage = function (e) {
  console.log(JSON.parse(e.data));
};

// keyboard input, ARROW to move, SHIFT + ARROW to move on axis
function moveUp(e) {
  moveAvatar(0, -2);
}
function moveDown(e) {
  moveAvatar(0, -2);
}
function moveLeft(e) {
  moveAvatar(0, -2);
}
function moveRight(e) {
  moveAvatar(0, -2);
}

// game loop optimized keyboard handling
kd.UP.down(moveUp);
kd.DOWN.down(moveDown);
kd.LEFT.down(moveLeft);
kd.RIGHT.down(moveRight);
