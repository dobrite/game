var scene = require('./scene');

//var ambientLight = new THREE.AmbientLight(0x10);
//scene.add(ambientLight);

var directionalLight = new THREE.DirectionalLight(0xffffff);
directionalLight.position.x = 500;
directionalLight.position.y = 1000;
directionalLight.position.z = 500;
directionalLight.position.normalize();
scene.add(directionalLight);
