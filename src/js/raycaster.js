var projector = require('./projector'),
    camera = require('./camera'),
    renderer = require('./renderer');

// didn't convert Y to Z
renderer.domElement.addEventListener('mousedown', function (event) {
  var vector = new THREE.Vector3(
    renderer.devicePixelRatio * (event.pageX - this.offsetLeft) / this.width * 2 - 1,
    -renderer.devicePixelRatio * (event.pageY - this.offsetTop) / this.height * 2 + 1,
    0
  );
  var raycaster = projector.pickingRay(vector, camera);
  //var intersects.length = raycaster.intersectObjects([arr of meshes]);
  if (intersects.length) {
    // intersects[0] is the clicked object
  }
}, false);
