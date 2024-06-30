import { defineStore } from "pinia";

export const useBillStore = defineStore("bills", {
    state: () = ({
        items: new Map(),
    }),
    getters: {
        allBills: (state) => state.items.values(),
        length: (state) => state.items.size
    },
    actions: {
        addOneBill(billInfo) {
            this.state.items.set(billInfo.id, billInfo);
        },
        
    }
})