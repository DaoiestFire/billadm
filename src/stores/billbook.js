/*存储账本信息*/
import { ElMessage } from "element-plus";
import { defineStore } from "pinia";

export const useBillbookStore = defineStore("billbooks", {
    state: () => {
        return {
            billbooks: new Map([
                ['default', {
                    value: "default",
                    label: "默认账本"
                }],
                ['salary', {
                    value: "salary",
                    label: "工资账本"
                }],
                ['property', {
                    value: "property",
                    label: "资产账本"
                }]
            ]),
            currentBook: '',
        }
    },
    getters: {
        getAllBillbooks: (state) => Array.from(state.billbooks.values()),
        getCurrentBook: (state) => state.currentBook,
    },
    actions: {
        setCurrentBook(bookvalue) {
            this.currentBook = bookvalue;
            ElMessage({
                message: `账本[${this.billbooks.get(bookvalue).label}]切换成功`,
                type: 'success'
            })
        },
        refreshBillbooks() {
            //TODO
            if (this.currentBook === '') {
                this.currentBook = 'default'
            }
        },
    }
})