import { createStore } from 'vuex';

export interface Anomaly {
  id: number;
  window: string;
  severity: string;
}

export const store = createStore({
  state: {
    anomalies: [] as Anomaly[],
  },
  mutations: {
    setAnomalies(state, anomalies: Anomaly[]) {
      state.anomalies = anomalies;
    },
  },
  actions: {
    async fetchAnomalies({ commit }) {
      const response = await fetch('http://localhost:8080/api/anomalies');
      const data = await response.json();
      commit('setAnomalies', data);
    },
  },
});
