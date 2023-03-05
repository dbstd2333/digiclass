<template>
  <div style="margin-left: 50px; margin-top: 0px">
    <div :style="{ marginTop: '20px' }">
      <a-radio-group
        size="large"
        type="button"
        v-model="kv.now_classcode"
        :model-value="kv.now_classcode"
        @change="changeclass(kv.now_classcode)"
      >
        <a-radio value="g1c1">高一 1班</a-radio>
        <a-radio value="g1c2">高一 2班</a-radio>
        <a-radio value="g1c3">高一 3班</a-radio>
        <a-radio value="g1c4">高一 4班</a-radio>
        <a-radio value="g1c5">高一 5班</a-radio>
        <a-radio value="g1c6">高一 6班</a-radio>
        <a-radio value="g1z1">高一 走班1</a-radio>
        <a-radio value="g1z2">高一 走班2</a-radio>
        <a-radio value="g1z3">高一 走班3</a-radio>
      </a-radio-group>
    </div>
    <div :style="{ marginTop: '20px' }">
      <a-radio-group
        size="large"
        type="button"
        v-model="kv.tweekly"
        :model-value="kv.tweekly"
        @change="changeweek(kv.tweekly)"
      >
        <a-radio value="1">星期一</a-radio>
        <a-radio value="2">星期二</a-radio>
        <a-radio value="3">星期三</a-radio>
        <a-radio value="4">星期四</a-radio>
        <a-radio value="5">星期五</a-radio>
        <a-radio value="6">星期六(上午)</a-radio>
        <a-radio value="0">星期日(下午)</a-radio>
      </a-radio-group>
    </div>
    <div style="margin-top: 80px">
      <a-timeline direction="horizontal" mode="alternate">
        <a-timeline-item label="早读">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession1"
                @change="xuankechange(kv.lession1, 'g1c1', 'subj', 1)"
              >
                <a-optgroup label="早读">
                  <a-option>语文</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  v-model="kv.teacher1"
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  @change="xuankechange(kv.teacher1, 'g1c1', 'teach', 1)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="上午第1节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession2"
                @change="xuankechange(kv.lession2, 'g1c1', 'subj', 2)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher2"
                  @change="xuankechange(kv.teacher2, 'g1c1', 'teach', 2)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src2"
              :value="kv.src2"
              @change="
                if (kv.src2 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 2);
                  kv.src2 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 2);
                  kv.src2 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="上午第2节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession3"
                @change="xuankechange(kv.lession3, 'g1c1', 'subj', 3)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher3"
                  @change="xuankechange(kv.teacher3, 'g1c1', 'teach', 3)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src3"
              :value="kv.src3"
              @change="
                if (kv.src3 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 3);
                  kv.src3 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 3);
                  kv.src3 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="上午第3节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession4"
                @change="xuankechange(kv.lession4, 'g1c1', 'subj', 4)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher4"
                  @change="xuankechange(kv.teacher4, 'g1c1', 'teach', 4)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src4"
              :value="kv.src4"
              @change="
                if (kv.src4 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 4);
                  kv.src4 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 4);
                  kv.src4 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="上午第4节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession5"
                @change="xuankechange(kv.lession5, 'g1c1', 'subj', 5)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher5"
                  @change="xuankechange(kv.teacher5, 'g1c1', 'teach', 5)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src5"
              :value="kv.src5"
              @change="
                if (kv.src5 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 5);
                  kv.src5 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 5);
                  kv.src5 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="上午第5节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession6"
                @change="xuankechange(kv.lession6, 'g1c1', 'subj', 6)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher6"
                  @change="xuankechange(kv.teacher6, 'g1c1', 'teach', 6)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src6"
              :value="kv.src6"
              @change="
                if (kv.src6 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 6);
                  kv.src6 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 6);
                  kv.src6 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="下午第1节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession7"
                @change="xuankechange(kv.lession7, 'g1c1', 'subj', 7)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher7"
                  @change="xuankechange(kv.teacher7, 'g1c1', 'teach', 7)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src7"
              :value="kv.src7"
              @change="
                if (kv.src7 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 7);
                  kv.src7 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 7);
                  kv.src7 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="下午第2节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession8"
                @change="xuankechange(kv.lession8, 'g1c1', 'subj', 8)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher8"
                  @change="xuankechange(kv.teacher8, 'g1c1', 'teach', 8)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src8"
              :value="kv.src8"
              @change="
                if (kv.src8 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 8);
                  kv.src8 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 8);
                  kv.src8 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="下午第3节">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession9"
                @change="xuankechange(kv.lession9, 'g1c1', 'subj', 9)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher9"
                  @change="xuankechange(kv.teacher9, 'g1c1', 'teach', 9)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src9"
              :value="kv.src9"
              @change="
                if (kv.src9 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 9);
                  kv.src9 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 9);
                  kv.src9 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
      </a-timeline>
    </div>
    <div style="margin-top: 80px">
      <a-timeline direction="horizontal">
        <a-timeline-item label="下午第4节(大练习1)">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession10"
                @change="xuankechange(kv.lession10, 'g1c1', 'subj', 10)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher10"
                  @change="xuankechange(kv.teacher10, 'g1c1', 'teach', 10)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src10"
              :value="kv.src10"
              @change="
                if (kv.src10 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 10);
                  kv.src10 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 10);
                  kv.src10 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="下午第5节(大练习2)">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession11"
                @change="xuankechange(kv.lession11, 'g1c1', 'subj', 11)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher11"
                  @change="xuankechange(kv.teacher11, 'g1c1', 'teach', 11)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src11"
              :value="kv.src11"
              @change="
                if (kv.src11 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 11);
                  kv.src11 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 11);
                  kv.src11 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="晚自习第1节(学科)">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession12"
                @change="xuankechange(kv.lession12, 'g1c1', 'subj', 12)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher12"
                  @change="xuankechange(kv.teacher12, 'g1c1', 'teach', 12)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src12"
              :value="kv.src12"
              @change="
                if (kv.src12 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 12);
                  kv.src12 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 12);
                  kv.src12 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="晚自习第2节(学科)">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession13"
                @change="xuankechange(kv.lession13, 'g1c1', 'subj', 13)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher13"
                  @change="xuankechange(kv.teacher13, 'g1c1', 'teach', 13)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src13"
              :value="kv.src13"
              @change="
                if (kv.src13 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 13);
                  kv.src13 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 13);
                  kv.src13 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
        <a-timeline-item label="晚自习第3-4节(公共)">
          <a-row :style="{ display: 'inline-flex', alignItems: 'center' }">
            <div :style="{ marginBottom: '12px' }">
              <a-select
                allow-search
                :style="{ width: '150px' }"
                placeholder="选择科目"
                v-model="kv.lession14"
                @change="xuankechange(kv.lession14, 'g1c1', 'subj', 14)"
              >
                <a-optgroup label="必选">
                  <a-option>语文</a-option>
                  <a-option>数学</a-option>
                  <a-option>英语</a-option><a-option>休息</a-option>
                </a-optgroup>
                <a-optgroup label="走班-理">
                  <a-option>物理</a-option>
                  <a-option>化学</a-option>
                  <a-option>生物</a-option>
                </a-optgroup>
                <a-optgroup label="走班-文">
                  <a-option>历史</a-option>
                  <a-option>政治</a-option>
                  <a-option>地理</a-option>
                </a-optgroup>
              </a-select>
              <div :style="{ fontSize: '12px', color: '#4E5969' }">
                <a-select
                  :style="{ width: '150px' }"
                  placeholder="选择老师"
                  style="margin-top: 10px"
                  allow-search
                  v-model="kv.teacher14"
                  @change="xuankechange(kv.teacher14, 'g1c1', 'teach', 14)"
                >
                  <a-optgroup label="语文老师">
                    <a-option>梁正平</a-option>
                    <a-option>裴玉东</a-option>
                    <a-option>王道军</a-option><a-option>休息</a-option>
                  </a-optgroup>
                  <a-optgroup label="数学老师">
                    <a-option>黄立秋</a-option>
                    <a-option>黄志平</a-option>
                    <a-option>蔡良君</a-option>
                  </a-optgroup>
                  <a-optgroup label="英语老师">
                    <a-option>缪雪剑</a-option>
                    <a-option>何心悦</a-option>
                    <a-option>张婷婷</a-option>
                    <a-option>孙敏</a-option>
                  </a-optgroup>
                  <a-optgroup label="理科老师">
                    <a-option>吴定才</a-option>
                    <a-option>曹体</a-option>
                    <a-option>何朝龙</a-option>
                    <a-option>陈冲</a-option>
                    <a-option>曾多</a-option>
                    <a-option>邓杰</a-option>
                  </a-optgroup>
                  <a-optgroup label="文科老师">
                    <a-option>陈琳琳</a-option>
                    <a-option>林艳茹</a-option>
                    <a-option>王艳洁</a-option>
                    <a-option>佘智兴</a-option>
                    <a-option>何仁春</a-option>
                  </a-optgroup>
                </a-select>
              </div>
            </div>
            <a-checkbox
              :model-value="kv.src14"
              :value="kv.src14"
              @change="
                if (kv.src14 == false) {
                  xuankechange(true, 'g1c1', 'zouban', 14);
                  kv.src14 = true;
                } else {
                  xuankechange(false, 'g1c1', 'zouban', 14);
                  kv.src14 = false;
                }
              "
              style="margin-left: 10px"
              >走班</a-checkbox
            >
          </a-row>
        </a-timeline-item>
      </a-timeline>
    </div>
    <div style="margin-left: 35%; margin-top: 50px">
      <a-button
        type="primary"
        size="large"
        long
        style="width: 400px"
        @click="submit()"
        >提交更改</a-button
      >
    </div>
  </div>
</template>

<script>
import kv from "../assets/kv.js";
import axios from "axios";
export default {
  name: "SetSubj",
  data() {
    return {
      kv,
    };
  },
  created() {},
  methods: {
    submit() {
      axios
        .post("/api/admin/changesubj", {
          classcode: kv.now_classcode,
          weekly1: kv.tweekly,
          lession1: kv.lession1,
          src1: kv.src1,
          teacher1: kv.teacher1,
          lession2: kv.lession2,
          src2: kv.src2,
          teacher2: kv.teacher2,
          lession3: kv.lession3,
          src3: kv.src3,
          teacher3: kv.teacher3,
          lession4: kv.lession4,
          src4: kv.src4,
          teacher4: kv.teacher4,
          lession5: kv.lession5,
          src5: kv.src5,
          teacher5: kv.teacher5,
          lession6: kv.lession6,
          src6: kv.src6,
          teacher6: kv.teacher6,
          lession7: kv.lession7,
          src7: kv.src7,
          teacher7: kv.teacher7,
          lession8: kv.lession8,
          src8: kv.src8,
          teacher8: kv.teacher8,
          lession9: kv.lession9,
          src9: kv.src9,
          teacher9: kv.teacher9,
          lession10: kv.lession10,
          src10: kv.src10,
          teacher10: kv.teacher10,
          lession11: kv.lession11,
          src11: kv.src11,
          teacher11: kv.teacher11,
          lession12: kv.lession2,
          src12: kv.src12,
          teacher12: kv.teacher12,
          lession13: kv.lession12,
          src13: kv.src13,
          teacher13: kv.teacher13,
          lession14: kv.lession14,
          src14: kv.src14,
          teacher14: kv.teacher14,
        })
        .then((res) => {
          console.log(res.data);
          if (res.data.status !== "ok" || res.status != 200) {
            this.$notification.error("提交失败");
          } else {
            this.$notification.success("提交成功");
          }
        });
    },
    changeclass(selectclass) {
      var msg = `切换到：` + selectclass;
      this.$message.info({ content: msg, showIcon: true });
    },
    changeweek() {
      axios
        .post("/api/admin/getsubj", {
          classcode: kv.now_classcode,
          weekly: kv.weekly,
        })
        .then((res) => {
          kv.weekly = res.data.weekly;
          kv.subject1 = res.data.lession1;
          kv.subject2 = res.data.lession2;
          kv.subject3 = res.data.lession3;
          kv.subject4 = res.data.lession4;
          kv.subject5 = res.data.lession5;
          kv.subject6 = res.data.lession6;
          kv.subject7 = res.data.lession7;
          kv.subject8 = res.data.lession8;
          kv.subject9 = res.data.lession9;
          kv.subject10 = res.data.lession10;
          kv.subject11 = res.data.lession11;
          kv.subject12 = res.data.lession12;
          kv.subject13 = res.data.lession13;
          kv.subject14 = res.data.lession14;
          kv.src1 = res.data.src1;
          kv.src2 = res.data.src2;
          kv.src3 = res.data.src3;
          kv.src4 = res.data.src4;
          kv.src5 = res.data.src5;
          kv.src6 = res.data.src6;
          kv.src7 = res.data.src7;
          kv.src8 = res.data.src8;
          kv.src9 = res.data.src9;
          kv.src10 = res.data.src10;
          kv.src11 = res.data.src11;
          kv.src12 = res.data.src12;
          kv.src13 = res.data.src13;
          kv.src14 = res.data.src14;
          kv.teacher1 = res.data.teacher1;
          kv.teacher2 = res.data.teacher2;
          kv.teacher3 = res.data.teacher3;
          kv.teacher4 = res.data.teacher4;
          kv.teacher5 = res.data.teacher5;
          kv.teacher6 = res.data.teacher6;
          kv.teacher7 = res.data.teacher7;
          kv.teacher8 = res.data.teacher8;
          kv.teacher9 = res.data.teacher9;
          kv.teacher10 = res.data.teacher10;
          kv.teacher11 = res.data.teacher11;
          kv.teacher12 = res.data.teacher12;
          kv.teacher13 = res.data.teacher13;
          kv.teacher14 = res.data.teacher14;
          console.log(res.data);
        });
      var msg = `切换到：星期` + kv.tweekly;
      this.$message.info({ content: msg, showIcon: true });
    },
    xuankechange(value, classcode, type, th) {
      console.log(value, classcode, type, th);
      console.log(kv);
    },
  },
};
</script>

<style>
</style>