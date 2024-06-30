<template>
    <el-container>
        <el-header height="40px">
            <div class="header-container">
                <div class="date-picker-container">
                    <el-date-picker v-model="timerange" type="daterange" unlink-panels range-separator="至"
                        start-placeholder="开始时间" end-placeholder="结束时间" :shortcuts="shortcuts" size="default" />
                </div>
                <el-button type="primary" @click="addBillInfo">
                    <el-icon>
                        <Plus />
                    </el-icon>
                    <span>新增记录</span>
                </el-button>
                <el-button type="danger" @click="handleBatchDelete">
                    <el-icon>
                        <Delete />
                    </el-icon>
                    <span>批量删除</span>
                </el-button>
            </div>
        </el-header>
        <el-main>
            <BillTable ref="billTableInstance" />
            <BillForm ref="billFormInstance" @submit-bill="handleSubmitBill" />
        </el-main>
    </el-container>
</template>

<script setup>
import { ref } from 'vue';
import BillTable from './BillTable.vue';
import BillForm from './BillForm.vue'
import {
    getLastMonthDate,
    getLastWeekData,
    getLastYearDate,
    getThisMonthData,
    getThisWeekData,
    getThisYearDate
} from '../utils/timeutils';

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

// function
const addBillInfo = () => {
    billFormInstance.value.showForm()
}

const handleSubmitBill = (billFormData) => {
    console.log(billFormData)
}

const handleBatchDelete = () => {
    billTableInstance.value.deleteSelectedBills()
}
</script>

<style scoped>
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
</style>
