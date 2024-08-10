<template>
  <el-container class="outer">
    <div class="drag-region"/>
    <el-header height="40px">
      <el-container>
        <AdvancedMenu v-if="billadmStore.showAdvancedMenu"/>
        <div class="left-control billadm-horizontal-left">
          <el-tooltip effect="dark" placement="right-start" content="菜单" v-bind="{ 'hide-after' : 0 }">
            <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius" :offset="buttonOffset"
                        @click="billadmStore.toggleShowAdvancedMenu">
              <el-text size="large">
                <SvgIcon name="menu" size="15"/>
              </el-text>
            </BillButton>
          </el-tooltip>
        </div>
        <div class="show-info billadm-vertical-all-center">
          <el-text>{{ billadmStore.workspaceState.current }}-{{ op }}</el-text>
        </div>
        <div class="window-control">
          <el-tooltip effect="dark" placement="bottom-start" content="最小化" v-bind="{ 'hide-after' : 0 }">
            <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius"
                        :offset="buttonOffset" @click="windowMinimize">
              <el-text size="large">
                <SvgIcon name="minus" size="15"/>
              </el-text>
            </BillButton>
          </el-tooltip>
          <el-tooltip effect="dark" placement="bottom-start" content="最大化" v-bind="{ 'hide-after' : 0 }">
            <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius"
                        :offset="buttonOffset" @click="windowMaximize">
              <el-text size="large">
                <SvgIcon name="square" size="15"/>
              </el-text>
            </BillButton>
          </el-tooltip>
          <el-tooltip effect="dark" placement="bottom-start" content="关闭" v-bind="{ 'hide-after' : 0 }">
            <BillButton :height="buttonSize" :width="buttonSize" :radius="buttonRadius"
                        :offset="buttonOffset" @click="winClose">
              <el-text size="large">
                <SvgIcon name="close" size="15"/>
              </el-text>
            </BillButton>
          </el-tooltip>
        </div>
      </el-container>
    </el-header>
    <el-main>
      <el-container>
        <el-aside v-if="billadmStore.showMenu" width="40px">
          <Menu/>
        </el-aside>
        <el-main>
          <InitWorkspaceForm/>
          <RouterView/>
        </el-main>
      </el-container>
    </el-main>
    <el-footer height="40px">
      <div class="show-statistic billadm-horizontal-right">
        <el-tooltip v-if="!billadmStore.showMenu" effect="dark" placement="right-start" v-bind="{ 'hide-after' : 0 }"
                    content="展开侧边菜单">
          <BillButton height="39px" :width="buttonSize" :offset="buttonOffset"
                      :radius="buttonRadius" @click="billadmStore.toggleShowMenu"
                      style="margin-right: auto">
            <el-text>
              <SvgIcon name="sidebar-expand" size="15"/>
            </el-text>
          </BillButton>
        </el-tooltip>
        <el-tooltip v-if="billadmStore.showMenu" effect="dark" placement="right-start" v-bind="{ 'hide-after' : 0 }"
                    content="折叠侧边菜单">
          <BillButton height="39px" :width="buttonSize" :offset="buttonOffset"
                      :radius="buttonRadius" @click="billadmStore.toggleShowMenu"
                      style="margin-right: auto">
            <el-text>
              <SvgIcon name="sidebar-collapse" size="15"/>
            </el-text>
          </BillButton>
        </el-tooltip>
        <el-text type="info">
          总记录数
          <el-text type="info" style="margin-right: 10px">{{ billsCnt }}</el-text>
          支出
          <el-text type="danger" style="margin-right: 10px">{{ billsSpend }}</el-text>
          收入
          <el-text type="success" style="margin-right: 10px">{{ billsIncome }}</el-text>
        </el-text>
        <el-tooltip effect="dark" placement="right-start" v-bind="{ 'hide-after' : 0 }" content="帮助">
          <BillButton height="39px" :width="buttonSize" :offset="buttonOffset" :radius="buttonRadius"
                      @click="billadmStore.toggleShowHelpMenu">
            <el-text>
              <SvgIcon name="help" size="15"/>
            </el-text>
          </BillButton>
        </el-tooltip>
      </div>
      <HelpMenu v-if="billadmStore.showHelpMenu"/>
    </el-footer>
  </el-container>
</template>

<script setup>
import {RouterView} from 'vue-router';
import Menu from '@/components/Menu.vue';
import SvgIcon from '@/components/base/SvgIcon.vue';
import BillButton from '@/components/base/BillButton.vue';
import InitWorkspaceForm from "@/components/InitWorkspaceForm.vue";
import {computed, onMounted, ref, watch} from "vue";
import {useBilladmStore} from '@/stores/billadm';
import AdvancedMenu from "@/components/AdvancedMenu.vue";
import router from "@/router";
import {menuItems} from "@/config/menu";
import HelpMenu from "@/components/HelpMenu.vue";

// store
const billadmStore = useBilladmStore();

const buttonSize = "40px";
const buttonRadius = "8px";
const buttonOffset = "10px";
const op = computed(() => {
  let name
  if (router.currentRoute.value.name) {
    name = router.currentRoute.value.name;
  } else {
    name = 'bill';
  }
  const item = menuItems.find((item) => item.name === name);
  return item.label;
});

const billsCnt = ref(0);
const billsIncome = ref(0);
const billsSpend = ref(0);

watch(() => billadmStore.bills, (bills) => {
  billsCnt.value = bills.length;
  billsIncome.value = 0;
  billsSpend.value = 0;
  for (let bill of bills) {
    if (bill.income === 'false') {
      billsSpend.value += Number(bill.money);
    } else {
      billsIncome.value += Number(bill.money);
    }
  }
})


// 窗口控制函数
const winClose = () => {
  window.appObject.send("window.close")
}

const windowMaximize = () => {
  window.appObject.send("window.maximize")
}

const windowMinimize = () => {
  window.appObject.send("window.minimize")
}

onMounted(async () => {
  const firstOpen = await billadmStore.isFirstOpen();
  billadmStore.showInitWorkspaceFormCloseButton = false;
  billadmStore.showInitWorkspaceForm = firstOpen;
  if (!firstOpen) {
    await billadmStore.refreshWorkspaceState();
    await billadmStore.refreshWorkspace();
  }
})
</script>


<style scoped>
.outer {
  height: 100vh;
  width: 100vw;
}

.drag-region {
  width: 100%;
  height: 40px;
  position: absolute;
  -webkit-app-region: drag;
}

.left-control {
  width: 120px;
  height: 40px;
}

.show-info {
  width: 100%;
  height: 40px;
}

.window-control {
  margin-left: auto;
  display: flex;
}

.show-statistic {
  width: 100%;
  height: 100%;
}

.el-header {
  border-bottom-width: 1px;
  border-bottom-color: var(--el-color-info-light-7);
  border-bottom-style: solid;
}

.el-footer {
  border-top-width: 1px;
  border-top-color: var(--el-color-info-light-7);
  border-top-style: solid;
}

.el-aside {
  height: 100%;
  border-right-width: 1px;
  border-right-color: var(--el-color-info-light-7);
  border-right-style: solid;
}

::-webkit-scrollbar {
  width: 0 !important;
}

::-webkit-scrollbar {
  width: 0 !important;
  height: 0;
}
</style>
