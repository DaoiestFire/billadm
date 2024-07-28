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
            <el-table-column label="操作" width="100px" align="center">
                <template #default="scope">
                    <div class="billadm-flex">
                        <el-tooltip effect="dark" placement="bottom-start" content="编辑" hide-after=0>
                            <BillButton height="30px" width="30px" radius="6px" offset="6px"
                                @click="handleEdit(scope.row)">
                                <SvgIcon name="pencil" size="15" />
                            </BillButton>
                        </el-tooltip>
                        <el-tooltip effect="dark" placement="bottom-start" content="删除" hide-after=0>
                            <BillButton height="30px" width="30px" radius="6px" offset="6px"
                                @click="handleDelete(scope.row.index)">
                                <SvgIcon name="trash2" size="15" />
                            </BillButton>
                        </el-tooltip>
                    </div>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script setup>
import { ElMessage } from 'element-plus';
import { computed, onMounted, ref, watchEffect } from 'vue'
import SvgIcon from '@/components/base/SvgIcon.vue';
import BillButton from '@/components/base/BillButton.vue';

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
    emit('updateOneBill', info)
    refreshTableDate()
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
        plain: true,
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

.billadm-flex {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>