<template>
    <el-table :data="tableData" style="width: 100%" max-height="500px" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="50px" />
        <el-table-column label="账单信息">
            <template #default="scope">
                <BillInfo :billInfo="scope.row.data">
                </BillInfo>
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
</template>

<script setup>
import { ref } from 'vue'
import BillInfo from './BillInfo.vue'

const multipleSelection = ref([])
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

const tableData = [
    {
        data: {
            money: "37.78",
            description: "晚饭",
            type: "饮食",
            time: "2024-06-05 08:23:56",
            tags: ["刘敬威", "吃饭", "正常消费"],
            income: true
        },
    },
]
</script>