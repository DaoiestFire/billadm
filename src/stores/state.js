/*存储各种前端状态*/
import {defineStore} from "pinia";

export const useStateStore = defineStore("state", {
    state: () => {
        return {
            showBillDisplayAside: true,
        }
    },
    actions: {
        toggleShowBillDisplayAside() {
            this.showBillDisplayAside = !this.showBillDisplayAside;
        },
    }
})