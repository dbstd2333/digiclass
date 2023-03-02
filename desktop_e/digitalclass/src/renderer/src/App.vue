<template>
  <div class="global">
    <HomePage />
    <div class="inf">
      Version：1.0.0<br>
      Design by <a-link href="https://www.douyin.com/user/MS4wLjABAAAAr8k1ALRoXaorW7jJ8yEM-rYguQeN7YXpbPSnEqWC4LE">db2333</a-link>
    </div>
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
.inf {
  color: rgb(91, 91, 91);
  text-align: center;
  font-size: small;
}

.global {
  background-color: rgba(255, 255, 255, 1);
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  overflow: hidden;
}
</style>
