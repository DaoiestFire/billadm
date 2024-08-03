<template>
  <div class="menu-container">
    <el-tooltip effect="dark" placement="right-start" v-bind="{ 'hide-after' : 0 }" v-for="item in menuItems"
                :content="item.label">
      <BillButton height="40px" width="40px" offset="10px" radius="8px" :is-active="activeName === item.name"
                  @click="clickMenuItem(item.name)">
        <el-text>
          <SvgIcon :name="item.icon" size="18"/>
        </el-text>
      </BillButton>
    </el-tooltip>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue';
import SvgIcon from '@/components/base/SvgIcon.vue';
import BillButton from '@/components/base/BillButton.vue';
import {menuItems} from '@/config/menu';
import {useBilladmStore} from '@/stores/billadm';
import {useRouter} from 'vue-router';

const billadmStore = useBilladmStore();
const router = useRouter();

// 变量
const activeName = ref('bill');

const clickMenuItem = (name) => {
  activeName.value = name;
  router.push({name: name});

  if (name === 'bill') {
    if (!billadmStore.showBillDisplayAside) {
      billadmStore.toggleShowBillDisplayAside();
    }
  }
};

onMounted(() => {
  router.push({name: 'bill'});
});
</script>

<style scoped>
.menu-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
