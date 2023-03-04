<template>
  <div class="gb">
    <div style="-webkit-app-region: drag; -webkit-user-select: none; height: 30px"></div>
    <div class="toptext">新津成外云平台</div>
    <div style="width: 300px; margin-top: 40px; margin-left: 30px">
      <a-form :model="gbv" layout="vertical">
        <a-form-item field="username" label="用户名">
          <a-input v-model="gbv.username" placeholder="班级识别码 最大20位" />
        </a-form-item>
        <a-form-item field="passwd" label="密码">
          <a-input v-model="gbv.passwd" placeholder="密码默认12345678 遗失请联系开发者" />
        </a-form-item>
      </a-form>
    </div>
    <div style="margin-left: 100px; margin-right: 100px; margin-top: 100px">
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
import { ref, watch } from 'vue'
import gbv from './gbv'
import axios from 'axios'
export default {
  name: 'IndexLoginPage',
  data() {
    return {
      value: '',
      data: [
        {
          label: '高一 1班',
          value: 'g1c1'
        },
        {
          label: '高一 2班',
          value: 'g1c2'
        },
        {
          label: '高一 3班',
          value: 'g1c3'
        },
        {
          label: '高一 4班',
          value: 'g1c4'
        },
        {
          label: '高一 5班',
          value: 'g1c5'
        },
        {
          label: '高一 6班',
          value: 'g1c6'
        }
      ],
      gbv
    }
  },
  methods: {
    async checkcookie() {
      if (gbv.classcode != 'null') {
        console.log('close')
        win = await ipcRenderer.invoke('index-quit')
        setTimeout(win.close(), 1000)
      } else {
        this.$message.error('获取cookie失败')
      }
    },
    newwindow() {
      const { BrowserWindow } = require('electron')
      const win = new BrowserWindow({
        width: 300,
        height: 900,
        show: false,
        frame: false,
        x: 1300,
        y: 10
      })
      win.once('ready-to-show', () => {
        win.show()
        win.loadURL('/home')
      })
      console.log('msgwin end')
    },
    getcookie() {
      const { ipcRenderer } = require('electron')
      var status,cookie,test = ipcRenderer.invoke('index-login', gbv.username, gbv.passwd)
      console.log("status:",status,"cookie:", cookie,test)
      if (status == 'error') {
        this.$message.error('登录失败')
      } else {
        if (status == 'ok') {
          gbv.classcode = cookie
          this.$message.success('登录成功')
          this.newwindow()
          this.$message.success('成功打开课表')
          setTimeout(this.checkcookie, 2000)
        }
      }
    }
  }
}
</script>

<style>
.toptext {
  color: black;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu,
    Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  font-weight: 900;
  font-size: x-large;
  margin-top: 20px;
  text-align: center;
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