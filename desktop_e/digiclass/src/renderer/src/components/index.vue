<template>
  <div class="gb">
    <div class="toptext">新成新高考在线排课</div>
    <div>
      <a-space class="selec" direction="vertical" size="large">
        <a-select
          v-model="value"
          :style="{ width: '320px' }"
          placeholder="请选择班级..."
          ><a-option v-for="item of data" :value="item" :label="item.label" />
        </a-select>
      </a-space>
    </div>
    <div></div>
    <div style="margin-left: 150px; margin-right: 150px; margin-top: 100px">
      <a-button type="primary" long @click="getcookie">登录</a-button>
    </div>
    <div class="inf">
      Version：1.0.0
      <br />
      Design by <a-link href="">db2333</a-link>
    </div>
  </div>
</template>

<script>
import { ref, watch } from "vue";
import { http } from "@tauri-apps/api";
import { appWindow, WebviewWindow } from "@tauri-apps/api/window";
import gbv from "./gbv";
import { readTextFile, writeTextFile, BaseDirectory } from "@tauri-apps/api/fs";
import { checkUpdate, installUpdate } from "@tauri-apps/api/updater";
export default {
  name: "IndexLoginPage",
  data() {
    return {
      value: "",
      data: [
        {
          label: "高一 1班",
          value: "g1c1",
        },
        {
          label: "高一 2班",
          value: "g1c2",
        },
        {
          label: "高一 3班",
          value: "g1c3",
        },
        {
          label: "高一 4班",
          value: "g1c4",
        },
        {
          label: "高一 5班",
          value: "g1c5",
        },
        {
          label: "高一 6班",
          value: "g1c6",
        },
      ],
    };
  },
  methods: {
    checkcookie() {
      if (gbv.classcode != "null") {
        console.log("close");
        setTimeout(appWindow.close(), 1000);
      } else {
        this.$message.error("获取cookie失败");
      }
    },
    newwindow() {
      const webview = new WebviewWindow("kebiaomain", {
        url: "/home",
        title: "Digital Class",
        height: 900,
        width: 300,
        x: 1300,
        y: 10,
        decorations: false,
        resizable: true,
      });
      webview.once("tauri://created", function () {
        // webview window successfully created
        console.log("success");
      });
      webview.once("tauri://error", function (e) {
        // an error happened creating the webview window
        console.log("err");
      });
      console.log("msgwin end");
    },
    getcookie() {
      http
        .fetch("http://127.0.0.1:10000/api/auth", {
          headers: { "Content-Type": "application/json" },
          method: "POST",
          body: http.Body.json({
            authcode: this.value.value,
            passwd: gbv.publicpwd,
          }),
        })
        .then((res) => {
          console.log(this.value.value);
          if (res.data.status == "error") {
            this.$message.error("登录失败");
          } else {
            gbv.classcode = res.data.cookie;
            this.$message.success("登录成功");
            this.newwindow();
            setTimeout(this.checkcookie, 3000);
          }
        });
    },
  },
};
</script>

<style>
.toptext {
  color: black;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
  font-weight: 900;
  font-size: x-large;
  top: 10px;
}
.inf {
  color: rgb(91, 91, 91);
  text-align: center;
  font-size: small;
  margin-top: 40px;
}
.selec {
  margin-top: 50px;
}
.gb {
}
</style>