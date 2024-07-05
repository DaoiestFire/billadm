<template>
    <el-container class="outer">
        <div class="window-control">
            <div class="control-button" @click="windowMinimize">
                <el-icon>
                    <SvgIcon name="minus" size="15" />
                </el-icon>
            </div>
            <div class="control-button" @click="windowMaximize">
                <el-icon>
                    <SvgIcon name="square" size="15" />
                </el-icon>
            </div>
            <div class="control-button" @click="winClose">
                <el-icon>
                    <SvgIcon name="close" size="15" />
                </el-icon>
            </div>
        </div>
        <el-container>
            <div class="app-aside">
                <el-aside width="200px">
                    <Menu />
                </el-aside>
            </div>
            <el-main>
                <BillDisplay />
            </el-main>
        </el-container>
    </el-container>
</template>

<script setup>
import Menu from './components/Menu.vue';
import BillDisplay from './components/BillDisplay.vue'
import SvgIcon from './components/SvgIcon.vue';

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
.outer {
    height: 100vh;
    width: 100vw;
}

.window-control {
    z-index: 100;
    top: 0px;
    right: 0px;
    position: absolute;
    display: flex;

}

::-webkit-scrollbar {
    width: 0 !important;
}

::-webkit-scrollbar {
    width: 0 !important;
    height: 0;
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

.app-aside {
    border-right-width: 1px;
    border-right-color: var(--el-color-info-light-7);
    border-right-style: solid;
}
</style>
