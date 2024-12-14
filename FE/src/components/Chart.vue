<script setup lang="ts">
import { Bar } from 'vue-chartjs';
import { ref, onMounted } from 'vue';
import { Chart as ChartJS, CategoryScale, BarElement, Title, Tooltip, Legend, LinearScale } from 'chart.js';

// Register necessary components for Chart.js
ChartJS.register(
  CategoryScale, // For categorical x-axis
  BarElement,    // For bar chart rendering
  Title,         // For adding chart titles
  Tooltip,       // For tooltips on hover
  Legend,        // For chart legend
  LinearScale    // For linear y-axis
);

const chartData = ref({
  labels: ['High', 'Medium', 'Low'],
  datasets: [
    {
      label: 'Windows by Severity',
      data: [0, 0, 0],  // Initial data
      backgroundColor: ['#ff0000', '#ffcc00', '#00ff00'],
    },
  ],
});

const chartOptions = ref({
  responsive: true,
  plugins: {
    legend: {
      position: 'top',
    },
  },
});

// Fetch data from the API and update the chart
const fetchAnomalies = async () => {
  const response = await fetch('/api/anomalies');
  const anomalies = await response.json();
  const severityCount = { High: 0, Medium: 0, Low: 0 };

  anomalies.forEach((anomaly) => {
    severityCount[anomaly.severity]++;
  });

  chartData.value.datasets[0].data = [
    severityCount['High'],
    severityCount['Medium'],
    severityCount['Low'],
  ];
};

onMounted(() => {
  fetchAnomalies();
});
</script>

<template>
  <div class="chart-container">
    <Bar :data="chartData" :options="chartOptions" />
  </div>
</template>
