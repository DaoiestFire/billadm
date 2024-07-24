<template>
    <div class="control-button" :style="getStyle">
        <div class="inner-button" :style="getInnerStyle" :class="{ active: props.isActive }">
            <slot></slot>
        </div>
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
    const { height, width } = props;
    return {
        height: height,
        width: width,
    };
});

const getInnerStyle = computed(() => {
    const { height, width, radius } = props;
    const value = String(Math.min(Number(height.replace('px', '')), Number(width.replace('px', ''))) - 10) + 'px';
    console.log(value);

    return {
        borderRadius: radius,
        height: value,
        width: value,
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

.inner-button {
    display: flex;
    justify-content: center;
    align-items: center;
}

.inner-button:hover {
    background-color: var(--el-color-primary-light-7);
}

.inner-button.active {
    background-color: var(--el-color-primary-light-7);
}
</style>
