<template>
  <div>
    <div class="toptex">新消息</div>
    <div style="{ display: flex }">
    <a-card :style="{ width: '450px' }" :title="gbv.msgtitle">
      <template #extra>
        <a-link>{{gbv.lastmsgtime}}</a-link>
      </template>
      {{gbv.msg}}
    </a-card>
  </div>
    <div style="margin-top: 30px;">
    <a-button @click="close" type="primary" size="large">关闭此窗口</a-button></div>
    <div class="inf">
      Version：1.0.0 <a-button type="text" size="mini">检查更新</a-button>
      <br />
      Design by <a-link href="">db2333</a-link>
    </div>
  </div>
</template>

<script>
import { Modal } from "@arco-design/web-vue";
import gbv from "./gbv";
export default {
  name: "MsgPage",
  data() {
    return {
        gbv
    };
  },
  created() {
    this.handleClick();
  },
  methods: {
    close() {
      appWindow.close();
    },
    handleClick() {
      http
        .fetch("http://127.0.0.1:10000/api/class/getmsg", {
          headers: { "Content-Type": "application/json" },
          method: "POST",
          body: http.Body.json({
            cookie: gbv.classcode,
          }),
        })
        .then((res) => {
          gbv.msg = res.data.msg;
          gbv.msgtitle = res.data.msgtitle;
          gbv.lastmsgtime = res.data.time.slice(0,-9)
        });
    },
  },
};
</script>

<style>
.inf {
  color: rgb(91, 91, 91);
  text-align: center;
  font-size: small;
  margin-top: 30px;
}
.toptex{
    color: black;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  font-weight: 900;
  font-size: x-large;
  top: 10px;
  margin-bottom: 30px;
}
</style>