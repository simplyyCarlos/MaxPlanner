import { defineConfig } from 'vite'
import deno from '@deno/vite-plugin'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [deno(), react()],
  server : {
    host: "0.0.0.0",
    port: 5173,
  },
});
