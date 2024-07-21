<template>
    <div class="control-button" :style="getStyle" :class="{ active: props.isActive }">
        <slot></slot>
    </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
    height: {
        type: String,
        required: true
    },
    width: {
        type: String,
        required: true,
    },
    radius: {
        type: String,
        default: "0px",
    },
    isActive: {
        type: Boolean,
        default: false
    }
});

const getStyle = computed(() => {
    const { height, width, radius } = props;
    return {
        height: height,
        width: width,
        borderRadius: radius,
    };
});

</script>

<style scoped>
.control-button {
    display: flex;
    justify-content: center;
    align-items: center;
    -webkit-app-region: no-drag;
}

.control-button:hover {
    background-color: var(--el-color-primary-light-9);
}

.active {
    background-color: var(--el-color-primary-light-9);
    color: var(--el-color-primary);
}
</style>
