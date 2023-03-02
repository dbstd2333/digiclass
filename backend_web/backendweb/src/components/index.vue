<template>
  <a-layout class="layout-demo">
    <a-layout>
      <a-layout-header
        ><a-space style="margin-left: 15px; margin-top: 20px">
          <img
            alt="logo"
            src="//p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image"
          />
          <a-typography-title
            :style="{ margin: 0, fontSize: '18px' }"
            :heading="6"
            class="title"
          >
            新津成外新高考走班排课管理后台
          </a-typography-title>
          <icon-menu-fold
            style="font-size: 22px; cursor: pointer"
            @click="toggleDrawerMenu"
          /> </a-space
      ></a-layout-header>
      <a-layout>
        <a-layout-sider
          ><a-menu
            :default-open-keys="['1']"
            :default-selected-keys="['01']"
            :style="{ width: '100%' }"
            @menu-item-click="onClickMenuItem"
          >
            <a-menu-item key="01">
              <IconHome></IconHome>
              修改课表
            </a-menu-item>
            <a-menu-item key="02">
              <IconCalendar></IconCalendar>
              发布消息
            </a-menu-item>
            <a-menu-item key="03">
              <IconCalendar></IconCalendar>
              走班统计
            </a-menu-item>
          </a-menu></a-layout-sider
        >
        <a-layout-content
          ><keep-alive><component :is="comName" /></keep-alive
        ></a-layout-content>
      </a-layout>
      <a-layout-footer>Design by db2333</a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import { defineComponent } from "vue";
import { Message } from "@arco-design/web-vue";
import {
  IconCaretRight,
  IconCaretLeft,
  IconHome,
  IconCalendar,
} from "@arco-design/web-vue/es/icon";
import SetSubj from "./kbmng.vue";
import StuMng from "./stumng.vue";
import SetMsg from "./setmsg.vue";

export default defineComponent({
  data() {
    return {
      comName: SetSubj,
    };
  },
  components: {
    IconCaretRight,
    IconCaretLeft,
    IconHome,
    IconCalendar,
    SetSubj,
    StuMng,
    SetMsg,
  },
  methods: {
    onClickMenuItem(key) {
      if (key == "01") {
        this.comName = SetSubj;
        Message.info({ content: `切换到修改课表`, showIcon: true });
        this.$forceUpdate()
      }
      if (key == "02") {
        this.comName = SetMsg;
        Message.info({ content: `切换到发布消息`, showIcon: true });
        this.$forceUpdate()
      }
      if (key == "03") {
        this.comName = StuMng;
        Message.info({ content: `切换到走班统计`, showIcon: true });
        this.$forceUpdate()
      }
    },
  },
});
</script>
<style scoped>
.layout-demo {
  height: 900px;
  background: var(--color-fill-2);
  border: 1px solid var(--color-border);
}
.layout-demo :deep(.arco-layout-sider) .logo {
  height: 32px;
  margin: 12px 8px;
  background: rgba(255, 255, 255, 0.2);
}
.layout-demo :deep(.arco-layout-sider-light) .logo {
  background: var(--color-fill-2);
}
.layout-demo :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
}
.layout-demo :deep(.arco-layout-footer) {
  height: 48px;
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  line-height: 48px;
}
.layout-demo :deep(.arco-layout-content) {
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-3);
}
.layout-demo :deep(.arco-layout-footer),
.layout-demo :deep(.arco-layout-content) {
  display: flex;
  flex-direction: column;
  color: var(--color-white);
  font-size: 16px;
  font-stretch: condensed;
  text-align: center;
  margin-top: 0px;
}
.navbar {
  display: flex;
  justify-content: space-between;
  height: 100%;
  background-color: var(--color-bg-2);
  border-bottom: 1px solid var(--color-border);
}

.left-side {
  display: flex;
  align-items: center;
  padding-left: 20px;
}
.title {
  margin-top: 20px;
}
</style>