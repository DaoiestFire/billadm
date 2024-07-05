<template>
    <div class="bill-table-outer">
        <el-table :data="tableData" style="width: 100%" cell-class-name="bill-table-cell" :max-height="tableHeight"
            @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="50px" align="center" />
            <el-table-column prop="money" label="金额" width="100px" align="center">
                <template #default="scope">
                    <span style="margin-left: 10px;"
                        :style="scope.row.income === 'true' ? 'color: #67c23a' : 'color: #f56c6c'">
                        {{ scope.row.income === 'true' ? "+" : "-" }}{{ scope.row.money }}</span>
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
                    <el-button-group>
                        <el-button size="small" @click="handleEdit(scope.row)">
                            <el-icon>
                                <SvgIcon name="pencil" size="15" />
                            </el-icon>
                        </el-button>
                        <el-popconfirm confirm-button-text="是" cancel-button-text="否" title="确认删除吗?"
                            @confirm="handleDelete(scope.row.index)">
                            <template #reference>
                                <el-button size="small" type="danger">
                                    <el-icon>
                                        <SvgIcon name="trash2" size="15" />
                                    </el-icon>
                                </el-button>
                            </template>
                        </el-popconfirm>
                    </el-button-group>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup>
import { ElMessage } from 'element-plus';
import { computed, onMounted, ref, watchEffect } from 'vue'
import SvgIcon from './SvgIcon.vue';

// 变量
const emit = defineEmits(['updateOneBill', 'upateStatisticDispaly'])
const multipleSelection = ref([])
const windowHeight = ref(window.innerHeight)
const tableHeight = computed(() => {
    return windowHeight.value - 90
})
const tableData = ref([])

// 窗口函数
window.addEventListener('resize', () => {
    windowHeight.value = window.innerHeight
})

// 单记录操作函数
const handleEdit = (info) => {
    tableData.value.push(
        {
            index: "index1",
            money: "1.74",
            description: "游戏",
            type: "娱乐",
            time: "2024-06-05",
            tags: ["吃饭", "消费", "cxbskxs"],
            income: 'true',
        }
    )
    // emit('updateOneBill', info)
    // refreshTableDate()
}

const handleDelete = (index) => {
    deleteBillsByList([index])
}

// 表格函数
const handleSelectionChange = (val) => {
    multipleSelection.value = []
    val.forEach((info) => {
        multipleSelection.value.push(info.index)
    })
}

const deleteSelectedBills = () => {
    deleteBillsByList(multipleSelection.value)
    multipleSelection.value = []
}

// 组件函数
// 根据当前账本与时间刷新表格数据
const refreshTableDate = () => {
    // TODO
    tableData.value = [
        {
            index: "index1",
            money: "37.78",
            description: "游戏",
            type: "娱乐",
            time: "2024-06-05",
            tags: ["刘vdf", "吃饭", "消费"],
            income: 'true',
        },
        {
            index: "index2",
            money: "3834",
            description: "晚饭",
            type: "饮食",
            time: "2024-06-05",
            tags: ["刘敬威", "吃饭", "正常消费"],
            income: 'false',
        },
        {
            index: "index3",
            money: "384",
            description: "晚饭",
            type: "饮食",
            time: "2024-06-05",
            tags: ["学习", "吃饭", "消费"],
            income: 'true',
        },
    ]
}

const deleteBillsByList = (indexList) => {
    // TODO
    let len = indexList.length
    console.log(indexList)
    ElMessage({
        message: `${len}条记录删除成功`,
        type: 'success',
    })
}

watchEffect(() => {
    let lengthIncome = 0
    let lengthCost = 0
    let totalIncome = 0
    let totalCost = 0
    tableData.value.forEach((info) => {
        if (info.income == 'true') {
            totalIncome += Number(info.money)
            lengthIncome++
        } else {
            totalCost += Number(info.money)
            lengthCost++
        }
    })

    emit('upateStatisticDispaly', {
        lengthIncome: lengthIncome,
        lengthCost: lengthCost,
        totalIncome: totalIncome,
        totalCost: totalCost,
    })
})

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
.bill-table-cell {
    padding: 3px !important;
}
</style>