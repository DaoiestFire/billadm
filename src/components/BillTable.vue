<template>
    <div class="bill-table-outer">
        <el-table :data="tableData" style="width: 100%" cell-class-name="bill-table-cell" :max-height="tableHeight"
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
                    <el-button size="small" @click="handleEdit(scope.row)">
                        编辑
                    </el-button>
                    <el-button size="small" type="danger" @click="handleDelete(scope.row.index)">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'

// 变量
const multipleSelection = ref([])
const windowHeight = ref(window.innerHeight)
const tableHeight = computed(() => {
    return windowHeight.value - 190
})
const tableData = ref()

// 窗口函数
window.addEventListener('resize', () => {
    windowHeight.value = window.innerHeight
})

// 单记录操作函数
const handleEdit = (info) => {
    console.log(info.index)
}

const handleDelete = (index) => {
    console.log(index)
}

// 表格函数
// 根据当前账本与时间刷新表格数据
const refreshTableDate = () => {
    tableData.value = [
        {
            index: "index1",
            money: "37.78",
            description: "游戏",
            type: "娱乐",
            time: "2024-06-05",
            tags: ["刘vdf", "吃饭", "消费"],
            income: true
        },
        {
            index: "index2",
            money: "3834",
            description: "晚饭",
            type: "饮食",
            time: "2024-06-05",
            tags: ["刘敬威", "吃饭", "正常消费"],
            income: false
        },
        {
            index: "index3",
            money: "384",
            description: "晚饭",
            type: "饮食",
            time: "2024-06-05",
            tags: ["学习", "吃饭", "消费"],
            income: false
        },
    ]
}

const handleSelectionChange = (val) => {
    multipleSelection.value = []
    val.forEach((info) => {
        multipleSelection.value.push(info.index)
    })
}

const deleteBillsByList = (indexList) => {
    console.log(indexList)
}

const deleteSelectedBills = () => {
    deleteBillsByList(multipleSelection.value)
    multipleSelection.value = []
}

// 组件函数
onMounted(() => {
    refreshTableDate()
})

// 导出成员
defineExpose({
    refreshTableDate,
    deleteSelectedBills,
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