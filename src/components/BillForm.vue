<template>
  <el-dialog v-model="billadmStore.showBillForm" :show-close="false" :close-on-press-escape="false"
             style="width: 632px;"
             :title="billadmStore.billForm.id === '' ? '新建记录' : '更新记录'">
    <el-form :model="billForm" label-width="auto" style="max-width: 600px" label-position="left" ref="bill-form">
      <el-form-item label="金额">
        <el-input v-model.trim="billadmStore.billForm.money"/>
      </el-form-item>
      <el-form-item label="收入/支出">
        <el-radio-group v-model="billadmStore.billForm.income">
          <el-radio value="false">支出</el-radio>
          <el-radio value="true">收入</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="类型">
        <el-select v-model="billadmStore.billForm.type" placeholder="选择类型">
          <el-option v-for="key in billadmStore.billTypes.keys()" :label="billadmStore.billTypes.get(key)"
                     :value="key"/>
        </el-select>
      </el-form-item>
      <el-form-item label="标签">
        <el-select v-model="billadmStore.billForm.tags" multiple placeholder="选择多个标签">
          <!--新增候选标签-->
          <template #footer>
            <el-button v-if="!isAdding" text bg size="small" @click="onAddOption">
              添加自定义标签
            </el-button>
            <template v-else>
              <el-input v-model.trim="optionName" class="option-input" placeholder="输入标签" size="small"/>
              <el-button type="primary" size="small" @click="onConfirm">
                确认
              </el-button>
              <el-button size="small" @click="clear">取消</el-button>
            </template>
          </template>
        </el-select>
      </el-form-item>
      <el-form-item label="时间">
        <el-date-picker v-model="billadmStore.billForm.creationTime" type="date" placeholder="选择时间"
                        style="width: 100%"/>
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model.trim="billadmStore.billForm.description" type="textarea"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="onCancel">退出</el-button>
        <el-button v-if="billadmStore.billForm.id === ''" @click="billadmStore.resetBillForm">重置</el-button>
        <el-button type="primary" @click="onSubmit">
          提交
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import {ref} from 'vue';
import {useBilladmStore} from '@/stores/billadm';

// store
const billadmStore = useBilladmStore();
// variable
const isAdding = ref(false)
const optionName = ref('')
// function: 标签操作
const onAddOption = () => {
  isAdding.value = true
}

const onConfirm = () => {
  if (optionName.value) {
    billadmStore.billForm.tags.push(optionName.value)
    clear()
  }
}

const clear = () => {
  isAdding.value = false
  optionName.value = ''
}

// function: 表单操作
const onCancel = () => {
  billadmStore.toggleShowBillForm();
}

const onSubmit = async () => {
  await billadmStore.addOneBill();
  await billadmStore.refreshBills();
  billadmStore.resetBillForm();
  billadmStore.toggleShowBillForm();
}
</script>
<style scoped>
.option-input {
  width: 100%;
  margin-bottom: 8px;
}
</style>