<template>
  <div class="bill-table-outer">
    <el-table :data="billadmStore.bills" style="width: 100%" cell-class-name="bill-table-cell" :max-height="tableHeight"
              @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="50px" align="center"/>
      <el-table-column prop="money" label="金额" width="100px" align="center">
        <template #default="scope">
                    <span style="margin-left: 10px;"
                          :style="scope.row.income === 'true' ? 'color: #67c23a' : 'color: #f56c6c'">
                        {{ scope.row.income === 'true' ? "+" : "-" }}{{ scope.row.money }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="type" label="类型" width="100px" align="center">
      </el-table-column>
      <el-table-column label="时间" width="120px" align="center">
        <template #default="scope">{{ timestampToLocalTimeString(scope.row.creationTime) }}</template>
      </el-table-column>
      <el-table-column prop="description" label="描述" width="150px" align="center">
      </el-table-column>
      <el-table-column prop="tags" label="标签">
        <template #default="scope">
          <el-space warp>
            <el-tag v-for="tag in scope.row.tags">{{ tag }}</el-tag>
          </el-space>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="100px" align="center">
        <template #default="scope">
          <div class="billadm-flex">
            <el-tooltip effect="dark" placement="bottom-start" content="编辑" v-bind="{ 'hide-after' : 0 }">
              <BillButton height="30px" width="30px" radius="6px" offset="6px"
                          @click="handleEdit(scope.row)">
                <SvgIcon name="pencil" size="15"/>
              </BillButton>
            </el-tooltip>
            <el-tooltip effect="dark" placement="bottom-start" content="删除" v-bind="{ 'hide-after' : 0 }">
              <BillButton height="30px" width="30px" radius="6px" offset="6px"
                          @click="handleDelete(scope.row.id)">
                <SvgIcon name="trash2" size="15"/>
              </BillButton>
            </el-tooltip>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
import {ElMessage} from 'element-plus';
import {computed, ref, watch} from 'vue'
import SvgIcon from '@/components/base/SvgIcon.vue';
import BillButton from '@/components/base/BillButton.vue';
import {useBilladmStore} from '@/stores/billadm';
import {timestampToLocalTimeString} from "@/utils/timeutils";

// store
const billadmStore = useBilladmStore();
// 变量
const emit = defineEmits(['updateOneBill', 'updateStatisticDisplay'])
const multipleSelection = ref([]);
// 窗口
const windowHeight = ref(window.innerHeight);
const tableHeight = computed(() => {
  return windowHeight.value - 90
});
window.addEventListener('resize', () => {
  windowHeight.value = window.innerHeight
});
// 单记录操作函数
const handleEdit = (info) => {
  emit('updateOneBill', info)
};

const handleDelete = (id) => {
  deleteBillsByList([id])
};
// 表格函数
const handleSelectionChange = (val) => {
  multipleSelection.value = []
  val.forEach((info) => {
    multipleSelection.value.push(info.id)
  })
};
// 组件函数
const deleteSelectedBills = () => {
  deleteBillsByList(multipleSelection.value)
  multipleSelection.value = []
};

const deleteBillsByList = (idList) => {

  ElMessage({
    message: `${idList.length}条记录删除成功`,
    type: 'success',
    plain: true,
  });
};

watch(computed(() => billadmStore.currentBook), () => {
  billadmStore.refreshBills();
});

// 导出成员
defineExpose({
  deleteSelectedBills,
});
</script>

<style>
.bill-table-cell {
  padding: 3px !important;
}

.billadm-flex {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>