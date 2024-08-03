/*存储账本信息*/
import {defineStore} from "pinia";
import {ElNotification} from "element-plus";

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
        async refreshBillbooks() {
            try {
                const res = await window.appObject.getAllBillbooks()
                this.billbooks = [];
                res.forEach(billbook => {
                    this.billbooks.push(billbook);
                });
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '获取账本失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
        },
        async refreshBills() {
            try {
                const res = await window.appObject.getAllBillsByBookID(this.currentBook);
                let bills = [];
                res.forEach(item => {
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
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '获取消费记录失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
        },
        async addOneBill() {
            let newItem = {
                money: Number(this.billForm.money),
                type: this.billForm.type,
                income: this.billForm.income,
                bookId: this.currentBook,
                description: this.billForm.description,
                tags: JSON.stringify(this.billForm.tags),
                creationTime: this.billForm.creationTime.getTime(),
            };
            try {
                await window.appObject.addOneBill(newItem);
                ElNotification({
                    type: 'success',
                    message: '新增消费记录成功',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '新增消费记录失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
        },
        async deleteBills(idList) {
            try {
                await window.appObject.deleteBills(idList);
                ElNotification({
                    type: 'success',
                    message: '删除消费记录成功',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '删除消费记录失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
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