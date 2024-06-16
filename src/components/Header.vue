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
        <div class="right-container">
            <el-select v-model="selectedBillBook" size="large" style="width: 240px;">
                <el-option v-for="item in billbooks" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
            <div class="window-control-container">
                <el-button type="primary" circle @click="windowMinimize">
                    <el-icon>
                        <Minus />
                    </el-icon>
                </el-button>
                <el-button type="primary" circle @click="windowMaximize">
                    <el-icon>
                        <FullScreen />
                    </el-icon>
                </el-button>
                <el-button type="primary" circle @click="winClose">
                    <el-icon>
                        <Close />
                    </el-icon>
                </el-button>
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

.right-container {
    margin-left: auto;
    display: flex;
    align-items: center;
}

.window-control-container {
    width: 200px;
    display: flex;
    justify-content: flex-end;
}
</style>
