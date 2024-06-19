<template>
    <el-dialog v-model="show" :show-close="false" :close-on-press-escape="false" title="新建记录">
        <el-form :model="billForm" label-width="auto" style="max-width: 600px" label-position="left" ref="bill-form">
            <el-form-item label="金额">
                <el-input v-model.trim="billForm.money" />
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
                            <el-input v-model.trim="optionName" class="option-input" placeholder="输入标签" size="small" />
                            <el-button type="primary" size="small" @click="onConfirm">
                                确认
                            </el-button>
                            <el-button size="small" @click="clear">取消</el-button>
                        </template>
                    </template>
                </el-select>
            </el-form-item>
            <el-form-item label="时间">
                <el-date-picker v-model="billForm.time" type="date" placeholder="选择时间" style="width: 100%" />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model.trim="billForm.description" type="textarea" />
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="onCancel">退出</el-button>
                <el-button type="primary" @click="onSubmit">
                    提交
                </el-button>
            </div>
        </template>
    </el-dialog>
</template>

<script setup>
import { reactive, ref, defineProps } from 'vue'

const show = ref(true)
const billForm = reactive({
    money: "",
    income: "false",
    type: "",
    time: new Date(),
    description: "",
    tags: [],
})

// variable
const isAdding = ref(false)
const optionName = ref('')

// function: 标签操作
const onAddOption = () => {
    isAdding.value = true
}

const onConfirm = () => {
    if (optionName.value) {
        billForm.tags.push(optionName.value)
        clear()
    }
}

const clear = () => {
    isAdding.value = false
    optionName.value = ""
}

// function: 表单操作
const onCancel = () => {
}

const onSubmit = () => {
    
    reset()
}

const reset = () => {
    billForm.money = "";
    billForm.income = "false";
    billForm.type = "";
    billForm.time = new Date();
    billForm.description = "";
    billForm.tags = [];
}
</script>
<style scoped>
.option-input {
    width: 100%;
    margin-bottom: 8px;
}
</style>