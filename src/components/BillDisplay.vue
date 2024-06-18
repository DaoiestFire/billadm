<template>
    <el-container>
        <el-header height="40px">
            <div class="header-container">
                <div class="date-picker-container">
                    <el-date-picker v-model="timerange" type="daterange" unlink-panels range-separator="至"
                        start-placeholder="开始时间" end-placeholder="结束时间" :shortcuts="shortcuts" :size="default" />
                </div>
                <el-button type="primary" @click="addBillInfo">
                    <el-icon>
                        <Plus />
                    </el-icon>
                    <span> 新增记录</span>
                </el-button>
                <el-button type="danger">
                    <el-icon>
                        <Delete />
                    </el-icon>
                    <span>批量删除</span>
                </el-button>
            </div>
        </el-header>
        <el-main>
            <BillTable />
            <BillForm v-if="isShowBillForm" />
        </el-main>
    </el-container>
</template>

<script setup>
import { ref } from 'vue';
import BillTable from './BillTable.vue';
import BillForm from './BillForm.vue'

const timerange = ref('')
const isShowBillForm = ref(false)
const shortcuts = [
    {
        text: '上周',
        value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
            return [start, end]
        },
    },
    {
        text: '上个月',
        value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
            return [start, end]
        },
    }
]

function addBillInfo() {
    isShowBillForm.value = true
}
</script>

<style scoped>
.header-container {
    height: 100%;
    float: right;
}

.date-picker-container {
    display: inline-block;
    margin-right: 50px;
}
</style>
