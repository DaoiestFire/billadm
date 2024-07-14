<template>
    <el-container class="outer" @contextmenu="windowUp">
        <div class="drag-region" @dblclick="doubleClick" @mousedown.self="windowDown" @mouseup.self="windowUp" />
        <el-header height="30px">
            <el-container>
                <BillButton height="30px" width="45px">
                    <SvgIcon name="menu" size="15" />
                </BillButton>
                <div class="window-control">
                    <BillButton height="30px" width="45px" @click="windowMinimize">
                        <SvgIcon name="minus" size="15" />
                    </BillButton>
                    <BillButton height="30px" width="45px" @click="windowMaximize">
                        <SvgIcon name="square" size="15" />
                    </BillButton>
                    <BillButton height="30px" width="45px" @click="winClose">
                        <SvgIcon name="close" size="15" />
                    </BillButton>
                </div>
            </el-container>
        </el-header>
        <el-main>
            <el-container>
                <el-aside width="200px">
                    <Menu />
                </el-aside>
                <el-main>
                    <BillDisplay />
                </el-main>
            </el-container>
        </el-main>
    </el-container>
</template>

<script setup>
import Menu from './components/Menu.vue';
import BillDisplay from './components/BillDisplay.vue'
import SvgIcon from './components/SvgIcon.vue';
import BillButton from './components/BillButton.vue';

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

const doubleClick = () => {
    window.windowController.send("window-maximize")
}

const windowDown = () => {
    window.windowController.send("window-move", true);
}

const windowUp = () => {
    window.windowController.send("window-move", false);
}
</script>


<style scoped>
.outer {
    height: 100vh;
    width: 100vw;
}

.drag-region {
    width: 100%;
    height: 30px;
    position: absolute;
}

.window-control {
    margin-left: auto;
    display: flex;
}

.el-header {
    padding: 0px;
    border-bottom-width: 1px;
    border-bottom-color: var(--el-color-info-light-7);
    border-bottom-style: solid;
}

.el-aside {
    padding: 5px;
    height: 100vh;
    border-right-width: 1px;
    border-right-color: var(--el-color-info-light-7);
    border-right-style: solid;
}

::-webkit-scrollbar {
    width: 0 !important;
}

::-webkit-scrollbar {
    width: 0 !important;
    height: 0;
}
</style>
