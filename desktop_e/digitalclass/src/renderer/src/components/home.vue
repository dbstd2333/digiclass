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
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src1" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject1 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    7:20-7:45
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="上午 第1节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src2" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject2 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    7:50-8:30
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="上午 第2节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src3" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject3 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    8:40-9:20
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="上午 第3节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src4" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject4 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    9:30-10:10
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="上午 第4节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src5" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject5 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    10:45-11:25
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="上午 第5节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src6" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject6 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    11:35-12:15
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                </a-timeline>
                <a-timeline mode="left" labelPosition="relative">
                    <a-timeline-item label="下午 第1节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src7" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject7 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    2:00-2:40
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="下午 第2节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src8" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject8 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    2:50-3:30
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="下午 第3节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src9" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject9 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    3:40-4:20
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="下午 第4节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src10" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject10 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    4:30-5:05
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="下午 第5节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src11" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject11 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    5:15-5:45
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                </a-timeline>

                <a-timeline mode="left" labelPosition="relative">
                    <a-timeline-item label="晚上 第1节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src12" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject12 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    6:50-7:30
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                    <a-timeline-item label="晚上 2-4节">
                        <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
                            <img width="40" :style="{ marginRight: '16px', marginBottom: '12px' }" :src="gbv.src13" />
                            <div :style="{ fontSize: '18px', marginBottom: '12px' }">
                                {{ gbv.subject13 }}
                                <div :style="{ fontSize: '12px', color: '#4E5969' }">
                                    7:35-10:00
                                </div>
                            </div>
                        </a-row>
                    </a-timeline-item>
                </a-timeline>
            </div>
        </a-spin>
        <div class="butt">
            <a-button @click="tomorrow" class="left">查看明天课表</a-button>
            <a-button @click="flush">返回</a-button>
        </div>
    </div>

</template>

<script>
import { ref } from 'vue';
import '@arco-design/web-vue/dist/arco.css';
//import axios from 'axios'
import fetch from 'electron-fetch'
import gbv from './gbv'
export default {
    name: 'HomePage',
    data() {
        return {
            gbv,
            loading: false
        }
    },
    methods: {
        tomorrow() {
            console.log("flush")
            this.loading = true
            fetch('http://127.0.0.1:10000/api/class/getsubject', {
                headers: { 'Content-Type': 'application/json' },
                method: 'POST',
                body: http.Body.json({
                    class: 2,
                    weekly: 1
                })
            }).then(res => {
                this.gbv.weekly = res.data.weekly
                this.gbv.subject1 = res.data.lession1
                this.gbv.subject2 = res.data.lession2
                this.gbv.subject3 = res.data.lession3
                this.gbv.subject4 = res.data.lession4
                this.gbv.subject5 = res.data.lession5
                this.gbv.subject6 = res.data.lession6
                this.gbv.subject7 = res.data.lession7
                this.gbv.subject8 = res.data.lession8
                this.gbv.subject9 = res.data.lession9
                this.gbv.subject10 = res.data.lession10
                this.gbv.subject11 = res.data.lession11
                this.gbv.subject12 = res.data.lession12
                this.gbv.subject13 = res.data.lession13
                this.gbv.src1 = res.data.src1
                this.gbv.src2 = res.data.src2
                this.gbv.src3 = res.data.src3
                this.gbv.src4 = res.data.src4
                this.gbv.src5 = res.data.src5
                this.gbv.src6 = res.data.src6
                this.gbv.src7 = res.data.src7
                this.gbv.src8 = res.data.src8
                this.gbv.src9 = res.data.src9
                this.gbv.src10 = res.data.src10
                this.gbv.src11 = res.data.src11
                this.gbv.src12 = res.data.src12
                this.gbv.src13 = res.data.src13
                this.gbv.zhiri = res.data.zhiri
                this.gbv.cleang = res.data.cleang
                this.$forceUpdate()
                this.loading = false
                return {
                    gbv,
                    loading
                }
            });
            console.log("flush end")
        },
        fls() {
            this.$forceUpdate();
            console.log("forceUpdate")
        },
        flush() {
            console.log("flush")
            fetch('http://127.0.0.1:10000//api/class/getsubject', {
                headers: { 'Content-Type': 'application/json' },
                method: 'POST',
                body: http.Body.json({
                    class: 2,
                    weekly: this.gbv.tomorrow
                })
            }).then(res => {
                this.gbv.weekly = res.data.weekly
                this.gbv.subject1 = res.data.lession1
                this.gbv.subject2 = res.data.lession2
                this.gbv.subject3 = res.data.lession3
                this.gbv.subject4 = res.data.lession4
                this.gbv.subject5 = res.data.lession5
                this.gbv.subject6 = res.data.lession6
                this.gbv.subject7 = res.data.lession7
                this.gbv.subject8 = res.data.lession8
                this.gbv.subject9 = res.data.lession9
                this.gbv.subject10 = res.data.lession10
                this.gbv.subject11 = res.data.lession11
                this.gbv.subject12 = res.data.lession12
                this.gbv.subject13 = res.data.lession13
                this.gbv.src1 = res.data.src1
                this.gbv.src2 = res.data.src2
                this.gbv.src3 = res.data.src3
                this.gbv.src4 = res.data.src4
                this.gbv.src5 = res.data.src5
                this.gbv.src6 = res.data.src6
                this.gbv.src7 = res.data.src7
                this.gbv.src8 = res.data.src8
                this.gbv.src9 = res.data.src9
                this.gbv.src10 = res.data.src10
                this.gbv.src11 = res.data.src11
                this.gbv.src12 = res.data.src12
                this.gbv.src13 = res.data.src13
                this.gbv.zhiri = res.data.zhiri
                this.gbv.cleang = res.data.cleang
                this.$forceUpdate()
                return {
                    gbv
                }
            });
            console.log("flush end")
        }
    },
    beforeCreate() {
        this.flush()
        return {
            mode,
            gbv,
        }

    },

}
</script>

<style>
.left{
    margin-left: 10px;
}
.zhiri {
    margin-left: 20px;
}

.butt {
    margin: 0 auto;
    text-align: center;
    margin-bottom: 10px;
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
    font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}
</style>