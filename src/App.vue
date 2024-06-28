<template>
    <el-container class="outer">
        <el-header height="50px">
            <Header />
        </el-header>
        <el-main>
            <el-card class="max">
                <BillDisplay />
            </el-card>
        </el-main>
    </el-container>
</template>

<script setup>
import Header from './components/Header.vue';
import BillDisplay from './components/BillDisplay.vue'
</script>


<style scoped>
.outer {
    height: 100vh;
    width: 100vw;
}

.max {
    height: 100%;
    width: 100%;
}

::-webkit-scrollbar {
    width: 0 !important;
}

::-webkit-scrollbar {
    width: 0 !important;
    height: 0;
}
</style>
