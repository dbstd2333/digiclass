import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite';
import { ArcoResolver } from 'unplugin-vue-components/resolvers';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(
  ), AutoImport({
    resolvers: [ArcoResolver()],
  }),
  Components({
    resolvers: [
      ArcoResolver({
        sideEffect: true
      })
    ]
  })],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: { //主要是加上这段代码
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:10000',	//实际请求地址
        changeOrigin: true
      },
    }
  }

})
