<template>
  <el-container style="height: 100%;">
    <el-aside v-if="showAside" width="200px">
      <el-container>
        <el-header height="40px">
          <el-container>
            <div class="aside-name">
              <el-text size="large" tag="b">账本</el-text>
            </div>
            <div class="aside-op-button">
              <el-tooltip effect="dark" placement="bottom-start" content="隐藏" hide-after=0>
                <BillButton height="40px" width="40px" radius="8px" offset="10px"
                            @click="stateStore.toggleShowBillDisplayAside">
                  <el-text>
                    <SvgIcon name="eye-off" size="15"/>
                  </el-text>
                </BillButton>
              </el-tooltip>
              <el-tooltip effect="dark" placement="bottom-start" content="添加账本" hide-after=0>
                <BillButton height="40px" width="40px" radius="8px" offset="10px">
                  <el-text>
                    <SvgIcon name="plus" size="15"/>
                  </el-text>
                </BillButton>
              </el-tooltip>
            </div>
          </el-container>
        </el-header>
        <el-container>
          <el-header height="10px"/>
          <el-main>
            <div class="billadm-vertical-center">
              <BillButton height="40px" width="180px" radius="8px" offset="10px" v-for="item in billbooks"
                          :is-active="selectedBillBook === item.value" :index="item.value"
                          @click="clickBillbookItem(item.value)">
                <el-text size="large">{{ item.label }}</el-text>
              </BillButton>
            </div>
          </el-main>
        </el-container>
      </el-container>
    </el-aside>
    <el-main>
      <el-container>
        <el-header height="40px">
          <div class="menu-header">
            <div class="date-picker-container">
              <el-date-picker v-model="timerange" type="daterange" unlink-panels range-separator="至"
                              start-placeholder="开始时间" end-placeholder="结束时间" :shortcuts="shortcuts"/>
            </div>
            <div class="button-container">
              <el-tooltip effect="dark" placement="bottom-start" content="新增记录" hide-after=0>
                <BillButton height="40px" width="40px" radius="8px" offset="10px" @click="addBillInfo">
                  <el-text>
                    <SvgIcon name="plus" size="15"/>
                  </el-text>
                </BillButton>
              </el-tooltip>
              <el-popconfirm confirm-button-text="是" cancel-button-text="否" title="确认删除吗?"
                             @confirm="handleBatchDelete">
                <template #reference>
                  <el-tooltip effect="dark" placement="bottom-start" content="批量删除" hide-after=0>
                    <BillButton height="40px" width="40px" radius="8px" offset="10px">
                      <el-text>
                        <SvgIcon name="trash" size="15"/>
                      </el-text>
                    </BillButton>
                  </el-tooltip>
                </template>
              </el-popconfirm>
            </div>
          </div>
        </el-header>
        <el-container>
          <el-header height="80px">
            <div class="statistic-dispaly">
              <el-row>
                <el-col :span="8">
                  <el-statistic title="记录总数/支出记录/收入记录" :value="lengthCost + lengthIncome">
                    <template #suffix>/{{ lengthCost }}/{{ lengthIncome }}</template>
                  </el-statistic>
                </el-col>
                <el-col :span="8">
                  <el-statistic title="总支出" :value="totalCost"/>
                </el-col>
                <el-col :span="8">
                  <el-statistic title="总收入" :value="totalIncome"/>
                </el-col>
              </el-row>
            </div>
          </el-header>
          <el-main>
            <BillTable ref="billTableInstance" @update-one-bill="handleBillEdit"
                       @upate-statistic-dispaly="handleUpdateStatistic"/>
            <BillForm ref="billFormInstance" @submit-bill="handleSubmitBill"/>
          </el-main>
        </el-container>
      </el-container>
    </el-main>
  </el-container>
</template>

<script setup>
import {computed, onMounted, ref} from 'vue';
import {ElMessage} from 'element-plus';
import BillTable from '@/components/BillTable.vue';
import BillForm from '@/components/BillForm.vue';
import SvgIcon from '@/components/base/SvgIcon.vue';
import BillButton from '@/components/base/BillButton.vue';
import {useBillbookStore} from '@/stores/billbook';
import {useStateStore} from '@/stores/state';
import {shortcuts} from '@/config/time_shortcuts';
// store
const stateStore = useStateStore();
const billbookStore = useBillbookStore();
// dom
const billFormInstance = ref(null);
const billTableInstance = ref(null);
// variable
const showAside = computed(() => stateStore.showBillDisplayAside);
const billbooks = ref([]);
const selectedBillBook = ref('');
const timerange = ref([new Date(), new Date()]);
const lengthIncome = ref(0);
const lengthCost = ref(0);
const totalIncome = ref(0);
const totalCost = ref(0);

// function
const clickBillbookItem = (bookname) => {
  selectedBillBook.value = bookname;
  billbookStore.setCurrentBook(bookname);
};

const addBillInfo = () => {
  billFormInstance.value.reset()
  billFormInstance.value.showForm()
};

const handleSubmitBill = (billFormData) => {
  // TODO
  console.log(billFormData)
  ElMessage({
    message: '操作成功',
    type: 'success',
    plain: true,
  })
};

const handleBillEdit = (info) => {
  billFormInstance.value.setBillForm(info)
  billFormInstance.value.showForm()
};

const handleUpdateStatistic = (info) => {
  lengthCost.value = info.lengthCost
  lengthIncome.value = info.lengthIncome
  totalCost.value = info.totalCost
  totalIncome.value = info.totalIncome
};

const handleBatchDelete = () => {
  billTableInstance.value.deleteSelectedBills()
};

// 组件函数
onMounted(() => {
  billbookStore.refreshBillbooks()
  billbookStore.getAllBillbooks.forEach((oneBillbook) => {
    billbooks.value.push(oneBillbook)
  })
  selectedBillBook.value = billbookStore.getCurrentBook
});
</script>

<style scoped>
.el-aside {
  height: 100vh;
  border-right-width: 1px;
  border-right-color: var(--el-color-info-light-7);
  border-right-style: solid;
}

.aside-name {
  display: flex;
  height: 40px;
  line-height: 40px;
  font-size: 18px;
  padding-left: 10px;
  justify-content: center;
  align-items: center;
}

.aside-op-button {
  margin-left: auto;
  display: flex;
}

.menu-header {
  height: 100%;
  display: flex;
  justify-content: end;
  align-items: center;
}

.date-picker-container {
  display: inline-block;
  margin-right: 50px;
}

.button-container {
  margin-right: 10px;
  display: flex;
}

.el-col {
  text-align: center;
}

.statistic-dispaly {
  padding-top: 16px;
}
</style>
