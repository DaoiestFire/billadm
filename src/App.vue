<script setup>
import Header from './components/Header.vue';
</script>

<template>
    <div class="common-layout">
        <el-container>
            <el-header>
                <Header />
            </el-header>
            <el-main>
                <el-empty description="没有数据" />
            </el-main>
        </el-container>
    </div>
</template>

<style scoped></style>
