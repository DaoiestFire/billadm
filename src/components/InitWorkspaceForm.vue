<template>
  <el-dialog v-model="billadmStore.showInitWorkspaceForm" :show-close="false" :close-on-press-escape="false" top="30vh"
             :close-on-click-modal="false"
             style="width: 632px;"
             title="选择工作空间目录">
    <el-text>首次打开软件，请选择一个目录作为工作空间或打开一个已存在的工作空间</el-text>
    <template #footer>
      <div class="dialog-footer">
        <el-input v-model="workspacePath" type="text" :clearable="true">
          <template #prepend>
            <el-button type="primary" @click="onChoose">
              选择目录
            </el-button>
          </template>
          <template #append>
            <el-button type="primary" @click="onSubmit">
              确认
            </el-button>
          </template>
        </el-input>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {useBilladmStore} from '@/stores/billadm';
import {ref} from "vue";

// store
const billadmStore = useBilladmStore();
// var
const workspacePath = ref('');
// function: 表单操作
const onChoose = async () => {
  workspacePath.value = await billadmStore.chooseWorkspaceDirectory()
}
const onSubmit = async () => {
  const flag = await billadmStore.initWorkspace(workspacePath.value);
  if (flag) {
    await billadmStore.refreshWorkspace();
    billadmStore.showInitWorkspaceForm = false;
  }
}
</script>
<style scoped>
</style>