/*存储账本信息*/
import {defineStore} from "pinia";
import {timestampToLocalTimeString} from "@/utils/timeutils";

export const useBilladmStore = defineStore("billbooks", {
    state: () => {
        return {
            billbooks: [],
            currentBook: '',
            bills: [],
            showBillDisplayAside: true,
            showBillForm: false,
            billForm: {
                id: '',
                money: '',
                type: '',
                income: 'false',
                bookId: '',
                description: '',
                tags: [],
                creationTime: new Date(),
            },
            timeRange: [new Date(), new Date()],
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
        refreshBills() {
            window.appObject.getAllBillsByBookID(this.currentBook).then((response) => {
                let bills = [];
                response.forEach(item => {
                    let newItem = {
                        id: item.id,
                        money: String(item.money),
                        type: item.type,
                        income: item.income,
                        bookId: item.book_id,
                        description: item.description,
                        tags: JSON.parse(item.tags),
                        creationTime: new Date(item.creation_time),
                    };
                    bills.push(newItem);
                });
                this.bills = bills;
            });
        },
        addOneBill() {
            let newItem = {
                money: Number(this.billForm.money),
                type: this.billForm.type,
                income: this.billForm.income,
                bookId: this.currentBook,
                description: this.billForm.description,
                tags: JSON.stringify(this.billForm.tags),
                creationTime: this.billForm.creationTime.getTime(),
            };
            window.appObject.addOneBill(newItem);
            this.refreshBills();
        },
        toggleShowBillDisplayAside() {
            this.showBillDisplayAside = !this.showBillDisplayAside;
        },
        toggleShowBillForm() {
            this.showBillForm = !this.showBillForm;
        },
        resetBillForm() {
            this.billForm = {
                id: '',
                money: '',
                type: '',
                income: 'false',
                bookId: '',
                description: '',
                tags: [],
                creationTime: new Date(),
            }
        },
    }
})