<template>
  <div class="kbglobal">
    <a-spin :loading="loading" tip="查询中，请稍等...">
      <div class="topic">
        {{ gbv.weekly }}
      </div>
      <div class="homeglobal">
        <a-timeline mode="left" labelPosition="relative">
          <a-timeline-item label="上午 早读">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src1"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject1 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher1}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="上午 第1节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src2"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject2 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher2}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="上午 第2节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src3"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject3 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher3}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="上午 第3节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src4"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject4 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher4}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="上午 第4节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src5"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject5 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher5}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="上午 第5节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src6"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject6 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher6}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
        </a-timeline>
        <a-timeline mode="left" labelPosition="relative">
          <a-timeline-item label="下午 第1节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src7"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject7 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher7}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="下午 第2节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src8"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject8 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher8}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="下午 第3节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src9"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject9 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher9}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="下午 第4节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src10"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject10 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher10}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
        </a-timeline>

        <a-timeline mode="left" labelPosition="relative">
          <a-timeline-item label="晚上 第1节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src11"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject11 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher11}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
          <a-timeline-item label="晚上 2-4节">
            <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
              <img
                width="40"
                :style="{ marginRight: '16px', marginBottom: '12px' }"
                :src="gbv.src12"
              />
              <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                {{ gbv.subject12 }}
                <div :style="{ fontSize: '12px', color: '#666666' }">
                  {{gbv.teacher12}}
                </div>
              </div>
            </a-row>
          </a-timeline-item>
        </a-timeline>
      </div>
    </a-spin>
    <div class="butt">
      <a-button type="outline" @click="tomorrow" class="left"
        >查看明天课表</a-button
      >
      <a-button type="outline" @click="flush" style="margin-left: 10px"
        >返回</a-button
      >
    </div>
    <div class="inf">
      Version：1.0.0
      <br />
      Design by <a-link href="">db2333</a-link>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import "@arco-design/web-vue/dist/arco.css";
import { http } from "@tauri-apps/api";
import { WebviewWindow } from "@tauri-apps/api/window";
import gbv from "./gbv";
export default {
  name: "HomePage",
  data() {
    return {
      gbv,
      loading: false,
      lastmsgid:0
    };
  },
  methods: {
    tomorrow() {
      console.log("flush");
      this.loading = true;
      http
        .fetch("http://127.0.0.1:10000/api/class/gettomrowsubj", {
          headers: { "Content-Type": "application/json" },
          method: "POST",
          body: http.Body.json({
            cookie: "g202201",
          }),
        })
        .then((res) => {
          this.gbv.weekly = res.data.weekly;
          this.gbv.subject1 = res.data.lession1;
          this.gbv.subject2 = res.data.lession2;
          this.gbv.subject3 = res.data.lession3;
          this.gbv.subject4 = res.data.lession4;
          this.gbv.subject5 = res.data.lession5;
          this.gbv.subject6 = res.data.lession6;
          this.gbv.subject7 = res.data.lession7;
          this.gbv.subject8 = res.data.lession8;
          this.gbv.subject9 = res.data.lession9;
          this.gbv.subject10 = res.data.lession10;
          this.gbv.subject11 = res.data.lession11;
          this.gbv.subject12 = res.data.lession12;
          this.gbv.subject13 = res.data.lession13;
          this.gbv.src1 = res.data.src1;
          this.gbv.src2 = res.data.src2;
          this.gbv.src3 = res.data.src3;
          this.gbv.src4 = res.data.src4;
          this.gbv.src5 = res.data.src5;
          this.gbv.src6 = res.data.src6;
          this.gbv.src7 = res.data.src7;
          this.gbv.src8 = res.data.src8;
          this.gbv.src9 = res.data.src9;
          this.gbv.src10 = res.data.src10;
          this.gbv.src11 = res.data.src11;
          this.gbv.src12 = res.data.src12;
          this.gbv.src13 = res.data.src13;
          this.gbv.zhiri = res.data.zhiri;
          this.gbv.cleang = res.data.cleang;
          this.$forceUpdate();
          this.loading = false;
          return {
            gbv,
            loading,
          };
        });
      console.log("flush end");
    },
    fls() {
      this.$forceUpdate();
      console.log("forceUpdate");
    },
    flush() {
      console.log("flush");
      http
        .fetch("http://127.0.0.1:10000/api/class/getsubject", {
          headers: { "Content-Type": "application/json" },
          method: "POST",
          body: http.Body.json({
            cookie: "g202201",
          }),
        })
        .then((res) => {
          this.gbv.weekly = res.data.weekly;
          this.gbv.subject1 = res.data.lession1;
          this.gbv.subject2 = res.data.lession2;
          this.gbv.subject3 = res.data.lession3;
          this.gbv.subject4 = res.data.lession4;
          this.gbv.subject5 = res.data.lession5;
          this.gbv.subject6 = res.data.lession6;
          this.gbv.subject7 = res.data.lession7;
          this.gbv.subject8 = res.data.lession8;
          this.gbv.subject9 = res.data.lession9;
          this.gbv.subject10 = res.data.lession10;
          this.gbv.subject11 = res.data.lession11;
          this.gbv.subject12 = res.data.lession12;
          this.gbv.subject13 = res.data.lession13;
          this.gbv.src1 = res.data.src1;
          this.gbv.src2 = res.data.src2;
          this.gbv.src3 = res.data.src3;
          this.gbv.src4 = res.data.src4;
          this.gbv.src5 = res.data.src5;
          this.gbv.src6 = res.data.src6;
          this.gbv.src7 = res.data.src7;
          this.gbv.src8 = res.data.src8;
          this.gbv.src9 = res.data.src9;
          this.gbv.src10 = res.data.src10;
          this.gbv.src11 = res.data.src11;
          this.gbv.src12 = res.data.src12;
          this.gbv.src13 = res.data.src13;
          this.gbv.zhiri = res.data.zhiri;
          this.gbv.cleang = res.data.cleang;
          this.$forceUpdate();
          return {
            gbv,
          };
        });
      console.log("flush end");
    },
    msgwindow() {
      const webview = new WebviewWindow("msg", {
        url: "/msg",
        title: "新消息",
        center: true,
        height: 250,
        width: 500,
        decorations: false,
        resizable: true,
      });
      webview.once("tauri://created", function () {
        // webview window successfully created
        console.log("success");
      });
      webview.once("tauri://error", function (e) {
        // an error happened creating the webview window
        console.log("err"+e);
      });
      console.log("msgwin end");
    },
    getMsg() {
      http
        .fetch("http://127.0.0.1:10000/api/class/getmsg", {
          headers: { "Content-Type": "application/json" },
          method: "POST",
          body: http.Body.json({
            cookie: gbv.classcode,
          }),
        })
        .then((res) => {
          console.log(res.data.msgid);
          console.log(res.data.classcode);
          console.log(gbv.msgid);
          if (res.data.msgid != gbv.msgid && res.data.time != gbv.lastmsgtime) {
            if (res.data.classcode == "all" || res.data.classcode == gbv.classcode) {
              this.msgwindow();
              gbv.msgid = res.data.msgid;
              gbv.lastmsgtime = res.data.time
              console.warn("different")
            }
            console.warn("!=")
          }
          setTimeout(this.getMsg, 10000);
        });
      console.log("get msg");
    },
  },
  created() {
    console.log("Created");
    const mode = ref("left");
    this.loading = true;
    this.flush();
    this.loading = false;
    this.getMsg();
    console.log("Created end");
    return {
      mode,
      gbv,
    };
  },
};
</script>

<style>
.left {
  margin-left: 10px;
}
.zhiri {
  margin-left: 20px;
}
.inf {
  color: rgb(91, 91, 91);
  text-align: center;
  font-size: small;
}
.butt {
  margin: 0 auto;
  text-align: center;
  margin-bottom: 30px;
}

.kbglobal {
  overflow: hidden;
}

.homeglobal {
  background-color: rgba(255, 255, 255, 1);
  margin-top: 30px;
  overflow: hidden;
}

.topic {
  background-color: rgba(255, 255, 255, 1);
  text-align: center;
  margin-top: 30px;
  font-size: xx-large;
  font-weight: 400;
  color: black;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
}
</style>