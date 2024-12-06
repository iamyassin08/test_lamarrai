<template>
    <div ref="threeContainer" class="three-container"></div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import * as THREE from 'three';
  import { OBJLoader } from 'three/examples/jsm/loaders/OBJLoader';
  import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
  
  const threeContainer = ref<HTMLElement | null>(null);
  
  onMounted(() => {
    const scene = new THREE.Scene();
    const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
    const renderer = new THREE.WebGLRenderer();
    const raycaster = new THREE.Raycaster();
    const mouse = new THREE.Vector2();
  
    renderer.setSize(window.innerWidth, window.innerHeight);
    if (threeContainer.value) threeContainer.value.appendChild(renderer.domElement);
  
    const loader = new OBJLoader();
    loader.load('/assets/building.obj', (obj) => {
      scene.add(obj);
      animate();
    });
  
    camera.position.z = 5;
  
    const animate = () => {
      requestAnimationFrame(animate);
      renderer.render(scene, camera);
    };
  
    // Handle mouse hover and interactions
    window.addEventListener('mousemove', (event) => {
      mouse.x = (event.clientX / window.innerWidth) * 2 - 1;
      mouse.y = -(event.clientY / window.innerHeight) * 2 + 1;
      raycaster.update(camera, renderer.domElement);
    });
  });
  </script>
  
  <style scoped>
  .three-container {
    width: 100%;
    height: 100vh;
  }
  </style>
  