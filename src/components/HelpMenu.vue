<template>
  <div class="advanced-menu billadm-vertical-all-center">
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="onclickWorkspace">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          <SvgIcon name="info-circle" size="15" style="margin-right: 10px"/>
          关于软件
        </div>
      </el-text>
    </BillButton>
  </div>
</template>

<script setup>
import SvgIcon from "@/components/base/SvgIcon.vue";
import BillButton from "@/components/base/BillButton.vue";
import {useBilladmStore} from "@/stores/billadm";
// store
const billadmStore = useBilladmStore();
// 工作空间
// 开发者工具
const toggleDevTools = () => {
  window.appObject.send('devtools.toggle');
}
</script>

<style scoped>
.advanced-menu {
  background-color: #ffffff;
  width: 200px;
  position: absolute;
  bottom: 40px;
  right: 10px;
  z-index: 100;
  box-shadow: var(--el-box-shadow);
  border-radius: 8px;
}

.menu-item {
  width: 160px;
}

.el-divider {
  margin: 0;
}

</style>