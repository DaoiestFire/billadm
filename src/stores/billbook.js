/*存储账本信息*/
import {defineStore} from "pinia";

export const useBillbookStore = defineStore("billbooks", {
    state: () => {
        return {
            billbooks: [],
            currentBook: '',
        }
    },
    actions: {
        setCurrentBook(bookValue) {
            this.currentBook = bookValue;
        },
        refreshBillbooks() {
            window.appObject.getAllBillbooks().then((response) => {
                response.forEach(billbook => {
                    this.billbooks.push(billbook);
                })
            });
        },
    }
})