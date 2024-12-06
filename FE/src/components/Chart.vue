<template>
    <div>
      <canvas ref="chartCanvas"></canvas>
    </div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { Chart, registerables } from 'chart.js';
  
  Chart.register(...registerables);
  
  const chartCanvas = ref<HTMLCanvasElement | null>(null);
  
  const data = {
    labels: ['High', 'Medium', 'Low'],
    datasets: [
      {
        label: 'Anomalies Severity',
        data: [10, 5, 3],
        backgroundColor: ['red', 'orange', 'green'],
      },
    ],
  };
  
  onMounted(() => {
    if (chartCanvas.value) {
      new Chart(chartCanvas.value, {
        type: 'bar',
        data,
        options: {
          responsive: true,
          plugins: {
            legend: {
              position: 'top',
            },
          },
        },
      });
    }
  });
  </script>
  
  <style scoped>
  canvas {
    width: 100%;
    height: 400px;
  }
  </style>
  