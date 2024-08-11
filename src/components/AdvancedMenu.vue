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

  <div v-if="showWorkspaceSecondMenu" class="advanced-menu billadm-vertical-all-center" style="left: 220px">
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="openOrCreateWorkspace">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          新建/打开
        </div>
      </el-text>
    </BillButton>
    <el-divider/>
    <BillButton v-for="(workspace,index) of billadmStore.workspaceState.workspaces" height="40px" width="200px"
                radius="8px"
                offset="10px" :is-active="billadmStore.workspaceState.current === workspace[0]"
                @click="onClickOneWorkspace(workspace[1],index)">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          {{ workspace[0] }}
          <SvgIcon name="chevron-right" size="15" style="margin-left: auto"/>
        </div>
      </el-text>
    </BillButton>
  </div>
  <div v-if="showWorkspaceThirdMenu" class="advanced-menu billadm-vertical-all-center" style="left: 430px"
       :style="getWorkspaceMenuStyle">
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="switchWorkspace">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          <SvgIcon name="link-external" size="15" style="margin-right: 10px"/>
          切换到
        </div>
      </el-text>
    </BillButton>
    <el-divider/>
    <BillButton height="40px" width="200px" radius="8px" offset="10px" @click="removeWorkspace">
      <el-text>
        <div class="billadm-horizontal-left menu-item">
          <SvgIcon name="trash" size="15" style="margin-right: 10px"/>
          从列表中移除
        </div>
      </el-text>
    </BillButton>
  </div>
</template>

<script setup>
import SvgIcon from "@/components/base/SvgIcon.vue";
import BillButton from "@/components/base/BillButton.vue";
import {computed, onMounted, ref} from "vue";
import {useBilladmStore} from "@/stores/billadm";
// store
const billadmStore = useBilladmStore();
// 工作空间
const showWorkspaceSecondMenu = ref(false);
const showWorkspaceThirdMenu = ref(false);
const checkedWorkspaceDir = ref('');
const workspaceIndex = ref(0);
const onclickWorkspace = () => {
  showWorkspaceSecondMenu.value = !showWorkspaceSecondMenu.value;
  if (!showWorkspaceSecondMenu.value) {
    showWorkspaceThirdMenu.value = false;
  }
}
const openOrCreateWorkspace = () => {
  billadmStore.showInitWorkspaceFormCloseButton = true;
  billadmStore.showInitWorkspaceForm = true;
}
const onClickOneWorkspace = (workspaceDir, index) => {
  checkedWorkspaceDir.value = workspaceDir;
  workspaceIndex.value = index;
  showWorkspaceThirdMenu.value = true;
}
const switchWorkspace = async () => {
  const flag = await billadmStore.initWorkspace(checkedWorkspaceDir.value);
  if (flag) {
    await billadmStore.refreshWorkspace();
    await billadmStore.refreshWorkspaceState();
    billadmStore.showAdvancedMenu = false;
  }
}
const removeWorkspace = async () => {
  console.log('start to remove workspace', checkedWorkspaceDir.value);
  const flag = await billadmStore.removeWorkspace(checkedWorkspaceDir.value);
  if (flag) {
    billadmStore.showAdvancedMenu = false;
    await billadmStore.refreshWorkspaceState();
    if (billadmStore.workspaceState.workspaces.length === 0) {
      billadmStore.showInitWorkspaceFormCloseButton = false;
      billadmStore.showInitWorkspaceForm = true;
    } else {
      await billadmStore.refreshWorkspace();
    }
  }
}
const getWorkspaceMenuStyle = computed(() => {
  const topValue = 80 + workspaceIndex.value * 40
  return {
    top: topValue + 'px',
  };
});
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
  background-color: #ffffff;
  width: 200px;
  position: absolute;
  top: 40px;
  left: 10px;
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