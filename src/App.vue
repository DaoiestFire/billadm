<template>
    <el-container class="outer">
        <div class="drag-region" />
        <el-header height="40px">
            <el-container>
                <el-tooltip effect="dark" placement="right-start" content="菜单" hide-after=0>
                    <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius" :offset="buttonOffset">
                        <SvgIcon name="menu" size="15" />
                    </BillButton>
                </el-tooltip>
                <div class="window-control">
                    <el-tooltip effect="dark" placement="bottom-start" content="最小化" hide-after=0>
                        <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius"
                            :offset="buttonOffset" @click="windowMinimize">
                            <SvgIcon name="minus" size="15" />
                        </BillButton>
                    </el-tooltip>
                    <el-tooltip effect="dark" placement="bottom-start" content="最大化" hide-after=0>
                        <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius"
                            :offset="buttonOffset" @click="windowMaximize">
                            <SvgIcon name="square" size="15" />
                        </BillButton>
                    </el-tooltip>
                    <el-tooltip effect="dark" placement="bottom-start" content="关闭" hide-after=0>
                        <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius"
                            :offset="buttonOffset" @click="winClose">
                            <SvgIcon name="close" size="15" />
                        </BillButton>
                    </el-tooltip>
                </div>
            </el-container>
        </el-header>
        <el-main>
            <el-container>
                <el-aside width="40px">
                    <Menu />
                </el-aside>
                <el-main>
                    <BillDisplay />
                </el-main>
            </el-container>
        </el-main>
        <el-footer height="40px">
            <span>no data</span>
        </el-footer>
    </el-container>
</template>

<script setup>
import Menu from './components/Menu.vue';
import BillDisplay from './components/BillDisplay.vue'
import SvgIcon from './components/base/SvgIcon.vue';
import BillButton from './components/base/BillButton.vue';

const buttonSize = "40px";
const buttonRadius = "8px";
const buttonOffset = "10px";

// 窗口控制函数
const winClose = () => {
    window.appObject.send("window-close")
}

const windowMaximize = () => {
    window.appObject.send("window-maximize")
}

const windowMinimize = () => {
    window.appObject.send("window-minimize")
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
    -webkit-app-region: drag;
}

.window-control {
    margin-left: auto;
    display: flex;
}

.el-header {
    border-bottom-width: 1px;
    border-bottom-color: var(--el-color-info-light-7);
    border-bottom-style: solid;
}

.el-footer {
    border-top-width: 1px;
    border-top-color: var(--el-color-info-light-7);
    border-top-style: solid;
}

.el-aside {
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
