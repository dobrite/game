var renderer = new THREE.WebGLRenderer();

renderer.setSize(config.SCENE_WIDTH, config.SCENE_HEIGHT);
document.body.appendChild(renderer.domElement);

module.exports = renderer;
