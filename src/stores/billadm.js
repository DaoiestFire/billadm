/*存储账本信息*/
import {defineStore} from "pinia";

export const useBilladmStore = defineStore("billbooks", {
    state: () => {
        return {
            billbooks: [],
            currentBook: '',
            showBillDisplayAside: true,
        }
    },
    actions: {
        setCurrentBook(bookValue) {
            this.currentBook = bookValue;
        },
        refreshBillbooks() {
            window.appObject.getAllBillbooks().then((response) => {
                this.billbooks = [];
                response.forEach(billbook => {
                    this.billbooks.push(billbook);
                });
            });
        },
        toggleShowBillDisplayAside() {
            this.showBillDisplayAside = !this.showBillDisplayAside;
        },
    }
})