<template>
  <el-dialog v-model="showSoftwareInfo" title="关于软件" width="500" :before-close="handleSoftwareClose" top="40vh">
    <el-row>
      <el-text size="large">Billadm是一款支持数据分析的桌面端记账工具</el-text>
    </el-row>
    <el-row>
      <el-text>作者：DaoiestFire</el-text>
    </el-row>
    <el-row>
      <el-text>软件版本：{{ appVersion }}</el-text>
    </el-row>
  </el-dialog>
  <div class="advanced-menu billadm-vertical-all-center">
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="onclickSoftwareInfo">
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
import {onMounted, ref} from "vue";
// store
const billadmStore = useBilladmStore();
// 关于软件
const appVersion = ref('');
const showSoftwareInfo = ref(false);
const handleSoftwareClose = () => {
  billadmStore.toggleShowHelpMenu();
}
const onclickSoftwareInfo = () => {
  showSoftwareInfo.value = !showSoftwareInfo.value;
}

onMounted(async () => {
  const billadmInfo = await window.appObject.getBilladmInfo();
  appVersion.value = billadmInfo.version;
})
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