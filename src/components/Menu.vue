<template>
    <div class="menu-container">
        <el-tooltip effect="dark" placement="right-start" hide-after=0 v-for="item in menuItems" :content="item.label">
            <BillButton height="40px" width="40px" offset="10px" radius="8px" :is-active="activeIndex === item.index"
                :index="item.index" @click="clickMenuItem(item.index)">
                <SvgIcon :name="item.icon" size="18" />
            </BillButton>
        </el-tooltip>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import SvgIcon from './base/SvgIcon.vue'
import BillButton from './base/BillButton.vue';
import { menuItems } from '../config/menu';

// 变量
const activeIndex = ref('bill')

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
