<template>
  <el-container style="height: 100%;">
    <el-aside v-if="billadmStore.showBillDisplayAside" width="200px">
      <el-container>
        <el-header height="40px">
          <el-container>
            <div class="aside-name">
              <el-text size="large" tag="b">账本</el-text>
            </div>
            <div class="aside-op-button">
              <el-tooltip effect="dark" placement="bottom-start" content="隐藏" v-bind="{ 'hide-after' : 0 }">
                <BillButton height="40px" width="40px" radius="8px" offset="10px"
                            @click="billadmStore.toggleShowBillDisplayAside">
                  <el-text>
                    <SvgIcon name="eye-off" size="15"/>
                  </el-text>
                </BillButton>
              </el-tooltip>
              <el-tooltip effect="dark" placement="bottom-start" content="添加账本" v-bind="{ 'hide-after' : 0 }">
                <BillButton height="40px" width="40px" radius="8px" offset="10px"
                            @click="addBillbook">
                  <el-text>
                    <SvgIcon name="plus" size="15"/>
                  </el-text>
                </BillButton>
              </el-tooltip>
              <el-tooltip effect="dark" placement="bottom-start" content="删除账本" v-bind="{ 'hide-after' : 0 }">
                <BillButton height="40px" width="40px" radius="8px" offset="10px" @click="deleteBillbook">
                  <el-text>
                    <SvgIcon name="trash" size="15"/>
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
              <BillButton height="40px" width="180px" radius="8px" offset="10px" v-for="item in billadmStore.billbooks"
                          :is-active="billadmStore.currentBook === item.id" :key="item.id"
                          @click="chooseBillbook(item.id)">
                <el-text size="large">{{ item.name }}</el-text>
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
              <el-date-picker v-model="billadmStore.timeRange" type="daterange" unlink-panels range-separator="至"
                              start-placeholder="开始时间" end-placeholder="结束时间" :shortcuts="shortcuts"
                              @change="billadmStore.handleTimeRangeChange"/>
            </div>
            <div class="button-container">
              <el-tooltip effect="dark" placement="bottom-start" content="新增记录" v-bind="{ 'hide-after' : 0 }">
                <BillButton height="40px" width="40px" radius="8px" offset="10px" @click="addBillInfo">
                  <el-text>
                    <SvgIcon name="plus" size="15"/>
                  </el-text>
                </BillButton>
              </el-tooltip>
              <el-tooltip effect="dark" placement="bottom-start" content="批量删除" v-bind="{ 'hide-after' : 0 }">
                <BillButton height="40px" width="40px" radius="8px" offset="10px" @click="handleBatchDelete">
                  <el-text>
                    <SvgIcon name="trash" size="15"/>
                  </el-text>
                </BillButton>
              </el-tooltip>
            </div>
          </div>
        </el-header>
        <el-container>
          <el-main>
            <BillTable ref="billTableInstance"/>
            <BillForm/>
            <BillbookForm/>
          </el-main>
        </el-container>
      </el-container>
    </el-main>
  </el-container>
</template>

<script setup>
import {onMounted, ref, toRaw} from 'vue';
import BillTable from '@/components/BillTable.vue';
import BillForm from '@/components/BillForm.vue';
import SvgIcon from '@/components/base/SvgIcon.vue';
import BillButton from '@/components/base/BillButton.vue';
import {useBilladmStore} from '@/stores/billadm';
import {shortcuts} from '@/config/time_shortcuts';
import BillbookForm from "@/components/BillbookForm.vue";

// store
const billadmStore = useBilladmStore();
//dom
const billTableInstance = ref(null);
// function
const addBillInfo = () => {
  billadmStore.resetBillForm();
  billadmStore.toggleShowBillForm();
};
const handleBatchDelete = async () => {
  await billTableInstance.value.deleteSelectedBills();
};
const chooseBillbook = async (id) => {
  billadmStore.setCurrentBook(id);
  await billadmStore.refreshBills();
};
const addBillbook = () => {
  billadmStore.resetBillbookForm();
  billadmStore.toggleShowBillbookForm();
};
const deleteBillbook = async () => {
  await billadmStore.deleteOneBillbook();
  await billadmStore.refreshBillbooks();
  await billadmStore.refreshBills();
};
</script>

<style scoped>
.el-aside {
  height: 100%;
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
</style>
