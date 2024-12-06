export const fetchAnomalies = async (): Promise<any[]> => {
    const response = await fetch('http://localhost:8080/api/anomalies');
    return await response.json();
  };
  