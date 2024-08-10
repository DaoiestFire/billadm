<template>
  <el-dialog v-model="billadmStore.showBillbookForm" :show-close="false" :close-on-press-escape="false"
             style="width: 632px;"
             title="新建账本">
    <el-form label-width="auto" style="max-width: 600px" label-position="left">
      <el-form-item label="名称">
        <el-input v-model.trim="billadmStore.billbookForm.name"/>
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="billadmStore.billbookForm.description" type="text" maxlength="30" show-word-limit/>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="onCancel">退出</el-button>
        <el-button type="primary" @click="onSubmit">
          创建
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {useBilladmStore} from '@/stores/billadm';

// store
const billadmStore = useBilladmStore();
// function: 表单操作
const onCancel = () => {
  billadmStore.toggleShowBillbookForm();
}

const onSubmit = async () => {
  await billadmStore.addOneBillbook();
  await billadmStore.refreshBillbooks();
  billadmStore.resetBillbookForm();
  billadmStore.toggleShowBillbookForm();
}
</script>
<style scoped>
</style>