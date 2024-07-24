<template>
    <el-container style="height: 100%;">
        <el-aside v-if="showAside" width="200px">
            <el-container>
                <el-header height="40px">
                    <el-container>
                        <div class="aside-name">
                            <span>账本</span>
                        </div>
                        <div class="aside-op-button">
                            <el-tooltip effect="dark" placement="bottom-start" content="隐藏" hide-after=0>
                                <BillButton height="40px" width="40px" radius="8px" offset="10px"
                                    @click="toggleShowAside">
                                    <SvgIcon name="eye-off" size="15" />
                                </BillButton>
                            </el-tooltip>
                            <el-tooltip effect="dark" placement="bottom-start" content="添加账本" hide-after=0>
                                <BillButton height="40px" width="40px" radius="8px" offset="10px">
                                    <SvgIcon name="plus" size="15" />
                                </BillButton>
                            </el-tooltip>
                        </div>
                    </el-container>
                </el-header>
            </el-container>
        </el-aside>
        <el-main>
            <el-container>
                <el-header height="50px">
                    <div class="menu-header">
                        <div class="billbook-select-container">
                            <el-select v-model="selectedBillBook" style="width: 160px;" @change="onSelectChange">
                                <el-option v-for="item in billbooks" :key="item.value" :label="item.label"
                                    :value="item.value" />
                            </el-select>
                        </div>
                        <div class="date-picker-container">
                            <el-date-picker v-model="timerange" type="daterange" unlink-panels range-separator="至"
                                start-placeholder="开始时间" end-placeholder="结束时间" :shortcuts="shortcuts" />
                        </div>
                        <div class="button-container">
                            <el-button-group>
                                <el-button type="primary" @click="addBillInfo">
                                    <el-icon>
                                        <SvgIcon name="plus" size="15" />
                                    </el-icon>
                                    <span style="">新增记录</span>
                                </el-button>
                                <el-popconfirm confirm-button-text="是" cancel-button-text="否" title="确认删除吗?"
                                    @confirm="handleBatchDelete">
                                    <template #reference>
                                        <el-button type="danger">
                                            <el-icon>
                                                <SvgIcon name="trash" size="15" />
                                            </el-icon>
                                            <span>批量删除</span>
                                        </el-button>
                                    </template>
                                </el-popconfirm>
                            </el-button-group>
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
                                    <el-statistic title="总支出" :value="totalCost" />
                                </el-col>
                                <el-col :span="8">
                                    <el-statistic title="总收入" :value="totalIncome" />
                                </el-col>
                            </el-row>
                        </div>
                    </el-header>
                    <el-main>
                        <BillTable ref="billTableInstance" @update-one-bill="handleBillEdit"
                            @upate-statistic-dispaly="handleUpdateStatistic" />
                        <BillForm ref="billFormInstance" @submit-bill="handleSubmitBill" />
                    </el-main>
                </el-container>
            </el-container>
        </el-main>
    </el-container>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import BillTable from './BillTable.vue';
import BillForm from './BillForm.vue';
import SvgIcon from './base/SvgIcon.vue';
import BillButton from './base/BillButton.vue';
import {
    getLastMonthDate,
    getLastWeekData,
    getLastYearDate,
    getThisMonthData,
    getThisWeekData,
    getThisYearDate
} from '../utils/timeutils';
import { ElMessage } from 'element-plus';
import { useBillbookStore } from '../stores/billbook'

// variable
const showAside = ref(true);
const hideAfter = 0
const billbooks = ref([]);
const selectedBillBook = ref('');
const billbookStore = useBillbookStore();
const billFormInstance = ref(null);
const billTableInstance = ref(null);
const timerange = ref([new Date(), new Date()]);
const shortcuts = [
    {
        text: '上周',
        value: getLastWeekData(),
    },
    {
        text: '上个月',
        value: getLastMonthDate(),
    },
    {
        text: '去年',
        value: getLastYearDate(),
    },
    {
        text: '本周',
        value: getThisWeekData(),
    },
    {
        text: '本月',
        value: getThisMonthData(),
    },
    {
        text: '今年',
        value: getThisYearDate(),
    },
    {
        text: '今天',
        value: [],
    },
    {
        text: '全部',
        value: [null, null],
    },
];
const lengthIncome = ref(0);
const lengthCost = ref(0);
const totalIncome = ref(0);
const totalCost = ref(0);

// 账本选择器函数
const onSelectChange = () => {
    billbookStore.setCurrentBook(selectedBillBook.value)
}

// function
const toggleShowAside = () => {
    showAside.value = !showAside.value;
}

const addBillInfo = () => {
    billFormInstance.value.reset()
    billFormInstance.value.showForm()
}

const handleSubmitBill = (billFormData) => {
    // TODO
    console.log(billFormData)
    ElMessage({
        message: '操作成功',
        type: 'success',
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
    console.log(billbookStore.getAllBillbooks)
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

.billbook-select-container {
    margin-left: 20px;
    margin-right: auto;
}

.date-picker-container {
    display: inline-block;
    margin-right: 50px;
}

.button-container {
    margin-right: 10px;
}

.el-col {
    text-align: center;
}

.statistic-dispaly {
    padding-top: 16px;
}
</style>
