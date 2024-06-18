<template>
    <div class="bill-form-outer">
        <el-card body-style="display: flex;justify-content: center">
            <div class="bill-form-inner">
                <el-form :model="billForm" label-width="auto" style="max-width: 600px" label-position="left">
                    <el-form-item label="金额">
                        <el-input v-model.numer="billForm.money" />
                    </el-form-item>
                    <el-form-item label="收入/支出">
                        <el-radio-group v-model="billForm.income">
                            <el-radio value="false">支出</el-radio>
                            <el-radio value="true">收入</el-radio>
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item label="类型">
                        <el-select v-model="billForm.type" placeholder="选择类型">
                            <el-option label="三餐" value="food" />
                            <el-option label="游戏" value="game" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="标签">
                        <el-select v-model="billForm.tags" multiple placeholder="选择多个标签">
                            <el-option label="Zone one" value="shanghai" />
                            <el-option label="Zone two" value="beijing" />
                            <template #footer>
                                <el-button v-if="!isAdding" text bg size="small" @click="onAddOption">
                                    添加自定义标签
                                </el-button>
                                <template v-else>
                                    <el-input v-model="optionName" class="option-input" placeholder="输入标签"
                                        size="small" />
                                    <el-button type="primary" size="small" @click="onConfirm">
                                        confirm
                                    </el-button>
                                    <el-button size="small" @click="clear">cancel</el-button>
                                </template>
                            </template>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="时间">
                        <el-date-picker v-model="billForm.time" type="date" placeholder="选择时间" style="width: 100%" />
                    </el-form-item>

                    <el-form-item label="备注">
                        <el-input v-model="billForm.description" type="textarea" />
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">Create</el-button>
                        <el-button>Cancel</el-button>
                    </el-form-item>
                </el-form>
            </div>

        </el-card>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'

const billForm = reactive({
    money: "37.78",
    income: "false",
    type: "娱乐",
    time: "2024-06-05",
    description: "游戏",
    tags: ["刘vdf", "吃饭", "消费"],
})

const isAdding = ref(false)
const onAddOption = () => {
    isAdding.value = true
}
const optionName = ref('')
const onConfirm = () => {
    if (optionName.value) {
        clear()
    }
}

const clear = () => {
    isAdding.value = false
}

const onSubmit = () => {
    console.log('submit!')
}
</script>
<style scoped>
.bill-form-outer {
    width: 650px;
}

.bill-form-inner {
    width: 600px;
}

.option-input {
    width: 100%;
    margin-bottom: 8px;
}
</style>