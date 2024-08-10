<template>
  <div class="advanced-menu billadm-vertical-all-center">
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="onclickWorkspace">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          <SvgIcon name="folder" size="15" style="margin-right: 10px"/>
          工作空间
          <SvgIcon name="chevron-right" size="15" style="margin-left: auto"/>
        </div>
      </el-text>
    </BillButton>
    <el-divider/>
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="toggleDevTools">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          <SvgIcon name="bug" size="15" style="margin-right: 10px"/>
          开发者工具
        </div>
      </el-text>
    </BillButton>
  </div>

  <div v-if="showSecondMenu" class="advanced-menu billadm-vertical-all-center" style="left: 220px">
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="openOrCreateWorkspace">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          新建/打开
        </div>
      </el-text>
    </BillButton>
    <el-divider/>
    <BillButton v-for="workspace of billadmStore.workspaceState.workspaces" height="40px" width="200px" radius="8px"
                offset="10px" :is-active="billadmStore.workspaceState.current === workspace[0]"
                @click="switchWorkspace(workspace[1])">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          {{ workspace[0] }}
          <SvgIcon name="chevron-right" size="15" style="margin-left: auto"/>
        </div>
      </el-text>
    </BillButton>
  </div>
</template>

<script setup>
import SvgIcon from "@/components/base/SvgIcon.vue";
import BillButton from "@/components/base/BillButton.vue";
import {onMounted, ref} from "vue";
import {useBilladmStore} from "@/stores/billadm";
// store
const billadmStore = useBilladmStore();
// 工作空间
const showSecondMenu = ref(false);
const onclickWorkspace = () => {
  showSecondMenu.value = !showSecondMenu.value;
}
const openOrCreateWorkspace = () => {
  billadmStore.showInitWorkspaceFormCloseButton = true;
  billadmStore.showInitWorkspaceForm = true;
}

const switchWorkspace = async (workspaceDir) => {
  const flag = await billadmStore.initWorkspace(workspaceDir);
  if (flag) {
    await billadmStore.refreshWorkspace();
    await billadmStore.refreshWorkspaceState();
    billadmStore.showAdvancedMenu = false;
  }
}
// 开发者工具
const toggleDevTools = () => {
  window.appObject.send('devtools.toggle');
}

onMounted(async () => {
  await billadmStore.refreshWorkspaceState();
})
</script>

<style scoped>
.advanced-menu {
  width: 200px;
  position: absolute;
  top: 40px;
  left: 10px;
  background-color: #ffffff;
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