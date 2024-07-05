<template>
    <div class="header-menu">
        <div class="control-button left-button">
            <SvgIcon name="menu" size="15" />
        </div>
        <div class="control-button">
            <SvgIcon name="plus" size="15" />
        </div>
    </div>
    <div class="billbook-select">
        <el-select v-model="selectedBillBook" size="default" style="width: 160px;" @change="onSelectChange">
            <el-option v-for="item in billbooks" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
    </div>
    <div class="menu-container">
        <div class="menu-button-outer" :class="{ 'menu-active': activeIndex == item.index }" v-for="item in menuItems"
            :index="item.index" @click="clickMenuItem(item.index)">
            <SvgIcon :name="item.icon" size="15" />
            <span style="display: inline-block;margin-left: 5px;">{{ item.label }}</span>
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useBillbookStore } from '../stores/billbook'
import SvgIcon from './SvgIcon.vue'

// 变量
const billbooks = ref([])
const selectedBillBook = ref('')
const billbookStore = useBillbookStore()
const activeIndex = ref('bill')
const menuItems = [
    {
        index: 'bill',
        icon: 'edit',
        label: '账单',
    },
    {
        index: 'analysis',
        icon: 'pie-chart',
        label: '统计',
    },
    {
        index: 'property',
        icon: 'coins',
        label: '资产',
    },
    {
        index: 'setting',
        icon: 'settings',
        label: '设置',
    }
]

// 账本选择器函数
const onSelectChange = () => {
    billbookStore.setCurrentBook(selectedBillBook.value)
}

const clickMenuItem = (index) => {
    activeIndex.value = index
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

.header-menu {
    display: flex;
    justify-content: end;
}

.menu-container {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.menu-button-outer {
    padding-left: 30px;
    height: 30px;
    width: 170px;
    display: flex;
    align-items: center;
    justify-content: start;
    color: var(--el-text-color-regular);
    border-radius: 8px;
    font-size: 15px;
}

.menu-button-outer:hover {
    background-color: var(--el-color-primary-light-9);
}

.menu-active {
    color: var(--el-color-primary);
    background-color: var(--aside-bg-color-light-2);
}

.left-button {
    margin-right: auto;
}

.control-button {
    height: 30px;
    width: 45px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.control-button:hover {
    background-color: var(--el-color-primary-light-9);
}
</style>
