<template>
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false">
        <div class="logo">
            <span>Billadm</span>
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
        <div class="right-container">
            <el-select v-model="selectedBillBook" size="default" style="width: 160px; margin-right: 80px;">
                <el-option v-for="item in billbooks" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
            <div class="control-button" @click="windowMinimize">
                <el-icon>
                    <Minus />
                </el-icon>
            </div>
            <div class="control-button" @click="windowMaximize">
                <el-icon>
                    <FullScreen />
                </el-icon>
            </div>
            <div class="control-button" @click="winClose">
                <el-icon>
                    <Close />
                </el-icon>
            </div>
        </div>
    </el-menu>
</template>

<script setup>
import { Close, FullScreen } from '@element-plus/icons-vue';
import { ref } from 'vue'

const activeIndex = ref('bill')
const billbooks = [
    {
        value: "default",
        label: "默认账本"
    },
    {
        value: "salary",
        label: "工资账本"
    }
]
const selectedBillBook = ref("default")

// 窗口控制函数
const winClose = () => {
    window.windowController.send("window-close")
}

const windowMaximize = () => {
    window.windowController.send("window-maximize")
}

const windowMinimize = () => {
    window.windowController.send("window-minimize")
}
</script>

<style scoped>
.logo {
    width: 150px;
    position: relative;
    -webkit-app-region: drag;
}

.logo span {
    margin: 0;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-size: 30px;
    font-weight: bold;
    color: var(--el-color-primary-dark-2);
}

.el-menu-demo {
    height: 50px;
}

.right-container {
    margin-left: auto;
    display: flex;
    align-items: center;
    justify-content: right;
}

.control-button {
    height: 50px;
    width: 50px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.control-button:hover {
    background-color: var(--el-color-primary-light-9);
}
</style>
