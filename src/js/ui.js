var s;
var ui = document.querySelectorAll('#coords')[0];
var sun = document.querySelectorAll('#sun')[0];

var updateUi = function (z, x, y, cz, cx, cy, sun) {
  ui.textContent = cz + ", " + cx + ", " + cz + ", " + z + ", " + x + ", " + y;
};

var updateSun = function (lighting) {
  sun.textContent = lighting.position.z + ", " + lighting.position.x + ", " + lighting.position.y;
};

module.exports = {
  updateUi: updateUi,
  updateSun: updateSun,
};
