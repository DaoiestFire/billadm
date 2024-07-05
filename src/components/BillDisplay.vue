<template>
    <el-container style="height: 100%;">
        <el-container>
            <el-header height="30px">
                <div class="menu-header">
                    <div class="header-container">
                        <div class="date-picker-container">
                            <el-date-picker v-model="timerange" type="daterange" unlink-panels range-separator="至"
                                start-placeholder="开始时间" end-placeholder="结束时间" :shortcuts="shortcuts" size="small" />
                        </div>
                    </div>
                    <el-button-group>
                        <el-button type="primary" size="small" text @click="addBillInfo">
                            <el-icon>
                                <SvgIcon name="plus" size="15" />
                            </el-icon>
                            <span style="">新增记录</span>
                        </el-button>
                        <el-popconfirm confirm-button-text="是" cancel-button-text="否" title="确认删除吗?"
                            @confirm="handleBatchDelete">
                            <template #reference>
                                <el-button type="danger" size="small" text>
                                    <el-icon>
                                        <SvgIcon name="trash" size="15" />
                                    </el-icon>
                                    <span>批量删除</span>
                                </el-button>
                            </template>
                        </el-popconfirm>
                    </el-button-group>
                </div>
            </el-header>
            <el-container>
                <el-header height="60px">
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
        <div class="billdispaly-aside">
            <el-aside width="200px">
            </el-aside>
        </div>
    </el-container>
</template>

<script setup>
import { ref } from 'vue'
import BillTable from './BillTable.vue'
import BillForm from './BillForm.vue'
import {
    getLastMonthDate,
    getLastWeekData,
    getLastYearDate,
    getThisMonthData,
    getThisWeekData,
    getThisYearDate
} from '../utils/timeutils'
import { ElMessage } from 'element-plus';
import SvgIcon from './SvgIcon.vue';

// variable
const billFormInstance = ref(null)
const billTableInstance = ref(null)
const timerange = ref([new Date(), new Date()])
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
]
const lengthIncome = ref(0)
const lengthCost = ref(0)
const totalIncome = ref(0)
const totalCost = ref(0)

// function
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
}

const handleBillEdit = (info) => {
    billFormInstance.value.setBillForm(info)
    billFormInstance.value.showForm()
}

const handleUpdateStatistic = (info) => {
    lengthCost.value = info.lengthCost
    lengthIncome.value = info.lengthIncome
    totalCost.value = info.totalCost
    totalIncome.value = info.totalIncome
}

const handleBatchDelete = () => {
    billTableInstance.value.deleteSelectedBills()
}
</script>

<style scoped>
.menu-header {
    height: 30px;
    display: flex;
    justify-content: end;
    align-items: center;
}

.header-container {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: right;
}

.date-picker-container {
    display: inline-block;
    margin-right: 50px;
}

.billdispaly-aside {
    border-left-width: 1px;
    border-left-color: var(--el-color-info-light-7);
    border-left-style: solid;
    background-color: var(--aside-bg-color-light);
}

.el-col {
    text-align: center;
}

.statistic-dispaly {
    padding-top: 10px;
}
</style>
