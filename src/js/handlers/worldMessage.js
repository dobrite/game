var world = require('../world');

var worldMessage = function (message) {
  world.render(message.data);
  //this should be a separate message
  //world.renderItems(message.data.i);
};

module.exports = worldMessage;
