/*存储各种前端状态*/
import { defineStore } from "pinia";

export const useStateStore = defineStore("state", {
    state: () => {
        return {
            showBillDisplayAside: true,
        }
    },
    getters: {
        getShowBillDisplayAside: (state) => state.showBillDisplayAside,
    },
    actions: {
        toggleShowBillDisplayAside() {
            this.showBillDisplayAside = !this.showBillDisplayAside
        },
    }
})