<template>
    <div class="bill-table-outer">
        <el-table :data="tableData" style="width: 100%" cell-class-name="bill-table-cell" :height="tableHeight"
            @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="50px" align="center" />
            <el-table-column prop="money" label="金额" width="100px" align="center">
                <template #default="scope">
                    <span style="margin-left: 10px;" :style="scope.row.income ? 'color: #67c23a' : 'color: #f56c6c'">
                        {{ scope.row.income ? "+" : "-" }}{{ scope.row.money }}</span>
                </template>
            </el-table-column>
            <el-table-column prop="type" label="类型" width="100px" align="center">
            </el-table-column>
            <el-table-column prop="time" label="时间" width="120px" align="center">
            </el-table-column>
            <el-table-column prop="description" label="描述" width="150px" align="center">
            </el-table-column>
            <el-table-column prop="tags" label="标签">
                <template #default="scope">
                    <el-space warp>
                        <el-tag v-for="tag in scope.row.tags">{{ tag }}</el-tag>
                    </el-space>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="150px" align="center">
                <template #default="scope">
                    <el-button size="small" @click="handleEdit(scope.$index, scope.row)">
                        编辑
                    </el-button>
                    <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup>
import { computed, ref, watchEffect } from 'vue'

const multipleSelection = ref([])
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

const tableData = [
    {
        money: "37.78",
        description: "游戏",
        type: "娱乐",
        time: "2024-06-05",
        tags: ["刘vdf", "吃饭", "消费"],
        income: true
    },
    {
        money: "3834",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: false
    },
    {
        money: "37.78",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: true
    },
    {
        money: "37.78",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: true
    },
    {
        money: "37.78",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: true
    },
    {
        money: "37.78",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: true
    },
    {
        money: "3237.8",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: false
    },
    {
        money: "37.78",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: true
    },
    {
        money: "37.78",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: true
    },
    {
        money: "37.78",
        description: "晚饭",
        type: "饮食",
        time: "2024-06-05",
        tags: ["刘敬威", "吃饭", "正常消费"],
        income: false
    },
]

// variable
const windowHeight = ref(window.innerHeight)
const tableHeight = computed(() => {
    return windowHeight.value - 190
})

// window function
window.addEventListener('resize', () => {
    windowHeight.value = window.innerHeight
})
</script>

<style>
.bill-table-outer {
    min-height: 0px;
}

.bill-table-cell {
    padding: 3px !important;
}
</style>