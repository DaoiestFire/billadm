<template>
    <div class="menu-container">
        <BillButton height="30px" width="180px" radius="8px" :is-active="activeIndex === item.index"
            v-for="item in menuItems" :index="item.index" @click="clickMenuItem(item.index)">
            <SvgIcon :name="item.icon" size="15" />
            <span style="display: inline-block; margin-left: 10px;">{{ item.label }}</span>
        </BillButton>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import SvgIcon from './SvgIcon.vue'
import BillButton from './BillButton.vue';

// 变量
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

const clickMenuItem = (index) => {
    activeIndex.value = index
}
</script>

<style scoped>
.billbook-select {
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
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
</style>
