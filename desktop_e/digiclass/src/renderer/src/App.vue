<template>
  <div class="global">
    <div data-tauri-drag-region class="top">点我移动</div>
    <router-view></router-view>
  </div>
</template>

<script>
import HomePage from './components/home.vue'
export default {
  name: 'App',
  components: {
    HomePage
  }
}
</script>

<style>
.top{
  text-align: right;
}

.global {
  background-color: rgba(255, 255, 255, 1);
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  overflow: hidden;
  text-align: center;
}
</style>
