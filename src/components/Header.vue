<template>
    <el-menu :default-active="activeIndex" :ellipsis="false">
        <div class="header-menu">
            <div class="control-button left-button">
                <el-icon>
                    <Menu />
                </el-icon>
                <icon-menu />
            </div>
            <div class="control-button">
                <el-icon>
                    <Plus />
                </el-icon>
            </div>
        </div>
        <div class="billbook-select">
            <el-select v-model="selectedBillBook" size="default" style="width: 160px;" @change="onSelectChange">
                <el-option v-for="item in billbooks" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
        </div>
        <el-menu-item index="bill">
            <el-icon>
                <Edit />
            </el-icon>
            <span>账单</span>
        </el-menu-item>
        <el-menu-item index="analysis">
            <el-icon>
                <PieChart />
            </el-icon>
            <span>统计</span>
        </el-menu-item>
        <el-menu-item index="property">
            <el-icon>
                <Coin />
            </el-icon>
            <span>资产</span>
        </el-menu-item>
        <el-menu-item index="setting">
            <el-icon>
                <Setting />
            </el-icon>
            <span>设置</span>
        </el-menu-item>
    </el-menu>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useBillbookStore } from '../stores/billbook';

// 变量
const activeIndex = ref('bill')
const billbooks = ref([])
const selectedBillBook = ref('')
const billbookStore = useBillbookStore()

// 账本选择器函数
const onSelectChange = () => {
    billbookStore.setCurrentBook(selectedBillBook.value)
}

// 组件函数
onMounted(() => {
    billbookStore.refreshBillbooks()
    console.log(billbookStore.getAllBillbooks)
    billbookStore.getAllBillbooks.forEach((oneBillbook) => {
        billbooks.value.push(oneBillbook)
    })
    selectedBillBook.value = billbookStore.getCurrentBook
})
</script>

<style scoped>
.billbook-select {
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.header-menu{
    display: flex;
    justify-content: end;
}

.left-button{
    margin-right: auto;
}

.control-button {
    height: 30px;
    width: 40px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.control-button:hover {
    background-color: var(--el-color-primary-light-9);
}
</style>
