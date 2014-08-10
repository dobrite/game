var world = require('../world');

var itemMessage = function (message) {
  world.renderItem(
    message.id,
    message.world_coords.coords[0],
    message.world_coords.coords[1],
    message.world_coords.chunk_coords[0],
    message.world_coords.chunk_coords[1],
    message.material_type
  );
};

module.exports = itemMessage;
