import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: parseInt(process.env.VITE_WEB_UI_PORT || '3000'),
    host: true,
    strictPort: true,
    watch: {
      usePolling: true,
    },
  },
})
