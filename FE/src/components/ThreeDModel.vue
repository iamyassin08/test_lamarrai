<template>
  <div ref="modelContainer" class="w-full h-full"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import * as THREE from 'three';
import { OBJLoader } from 'three/examples/jsm/loaders/OBJLoader';
// import { OBJLoader } from "../assets/3D/scene_mesh_textured_material_18_map_Kd.jpg"

// References for 3D scene, camera, renderer, and models
const modelContainer = ref<HTMLElement | null>(null);
const scene = ref<THREE.Scene | null>(null);
const camera = ref<THREE.PerspectiveCamera | null>(null);
const renderer = ref<THREE.WebGLRenderer | null>(null);

// Function to initialize the 3D model viewer
const init3DModel = () => {
  scene.value = new THREE.Scene();
  camera.value = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
  renderer.value = new THREE.WebGLRenderer();
  
  // Set renderer size and append to container
  if (modelContainer.value) {
    renderer.value.setSize(window.innerWidth, window.innerHeight);
    modelContainer.value.appendChild(renderer.value.domElement);
  }

  const light = new THREE.AmbientLight(0xffffff, 1); // Add ambient light to the scene
  scene.value.add(light);

  // Load multiple 3D OBJ files dynamically (from _00_ to _18_)
  const loader = new OBJLoader();
  for (let i = 0; i <= 18; i++) {
    const filePath = `/src/assets/3D/scene_mesh_textured_material_18_map_Kd_${String(i).padStart(2, '0')}.obj`;
    loader.load(
      filePath,
      (object) => {
        // When the object is loaded, add it to the scene
        scene.value?.add(object);
      },
      undefined,
      (error) => {
        console.error(`Error loading model ${filePath}:`, error);
      }
    );
  }

  // Set the camera position
  if (camera.value) {
    camera.value.position.z = 5;
  }
};

// Animation loop to continuously render the scene
const animate = () => {
  if (renderer.value && scene.value && camera.value) {
    requestAnimationFrame(animate);
    renderer.value.render(scene.value, camera.value);
  }
};

// Add hotspot to the 3D model at a specific position
const addHotspot = (x: number, y: number, z: number, label: string) => {
  if (scene.value) {
    const geometry = new THREE.SphereGeometry(0.1, 32, 32);
    const material = new THREE.MeshBasicMaterial({ color: 0xff0000 });
    const sphere = new THREE.Mesh(geometry, material);
    sphere.position.set(x, y, z);
    scene.value.add(sphere);

    sphere.userData = { label: label };

    // Add hover effect to display a tooltip
    sphere.on('mouseover', () => {
      alert(`Hotspot: ${label}`);
    });
  }
};

// Set up the model and start the animation when the component is mounted
onMounted(() => {
  init3DModel();
  animate();
});

// Clean up the scene and renderer when the component is unmounted
onBeforeUnmount(() => {
  if (renderer.value && modelContainer.value) {
    modelContainer.value.removeChild(renderer.value.domElement);
  }
});
</script>
