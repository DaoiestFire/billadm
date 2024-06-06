<template>
    <el-card shadow="hover" body-style="padding: 0px">
        <div class="billinfo-outer">
            <div class="billinfo-inner">
                <div class="billinfo-money" :style="props.billInfo.income ? 'color: #67c23a' : 'color: #f56c6c'">
                    <span>{{ props.billInfo.income ? "+" : "-" }}{{ props.billInfo.money }}</span>
                </div>
                <div class="billinfo-inner2">
                    <div class="billinfo-description">
                        <span class="billinfo-title">备注:</span>
                        <span>{{ props.billInfo.description }}</span>
                    </div>
                    <div class="billinfo-inner3">
                        <div class="billinfo-type">
                            <span class="billinfo-title">类别:</span>
                            <span>{{ props.billInfo.type }}</span>
                        </div>
                        <div class="billinfo-time">
                            <span class="billinfo-title">时间:</span>
                            <span>{{ props.billInfo.time }}</span>
                        </div>
                    </div>
                </div>
            </div>
            <div class="billinfo-tag">
                <span class="billinfo-title">标签:</span>
                <el-space warp>
                    <el-tag v-for="tag in props.billInfo.tags">{{ tag }}</el-tag>
                </el-space>
            </div>
        </div>
    </el-card>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps(["billInfo"])

</script>

<style scoped>
.billinfo-outer {
    display: flex;
    flex-direction: column;
    border-width: 0px;
    border-color: var(--el-color-info-light-8);
    border-style: solid;
}

.billinfo-inner {
    height: 60px;
    display: flex;
    border-width: 0px 0px 1px 0px;
    border-color: var(--el-color-info-light-8);
    border-style: solid;
}

.billinfo-tag {
    height: 40px;
    display: flex;
    align-items: center;
}

.billinfo-money {
    position: relative;
    width: 30%;

}

.billinfo-money span {
    margin: 0;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-size: 40px;
    line-height: 100%;
}

.billinfo-inner2 {
    width: 70%;
    display: flex;
    flex-direction: column;
    border-width: 0px 0px 0px 1px;
    border-color: var(--el-color-info-light-8);
    border-style: solid;
}

.billinfo-description {
    height: 30px;
    display: flex;
    align-items: center;
}

.billinfo-inner3 {
    height: 30px;
    display: flex;
    border-width: 1px 0px 0px 0px;
    border-color: var(--el-color-info-light-8);
    border-style: solid;
}

.billinfo-type {
    width: 40%;
    display: flex;
    align-items: center;
}

.billinfo-time {
    width: 60%;
    border-width: 0px 0px 0px 1px;
    border-color: var(--el-color-info-light-8);
    border-style: solid;
    display: flex;
    align-items: center;
}

.billinfo-title {
    margin-left: 15px;
    margin-right: 15px;
    font-weight: 700;
}
</style>