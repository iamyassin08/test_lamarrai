import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(), 
    ],

  server: {
    cors: false,
    host: true,
    port: 5173,
    proxy: {
      '*': {
        target: 'https://prep.com',
        ws: true,
        changeOrigin: true
      }
    }
  },

  resolve: {

    alias: [
      {
        find: /@\/components\/((?!.*[.](ts|js|tsx|jsx|vue)$).*$)/,
        replacement: fileURLToPath(
            new URL("./src/components/$1/index.vue", import.meta.url)
        ),
    },
      {
        find: '@',
        replacement: fileURLToPath(new URL('./src', import.meta.url))
      }
  ]
  },

  build: {
    sourcemap: true
  }
})