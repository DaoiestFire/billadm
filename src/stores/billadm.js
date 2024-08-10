/*存储账本信息*/
import {defineStore} from "pinia";
import {ElNotification} from "element-plus";
import {BUILT_IN_BILLBOOK} from "@/utils/constants";
import {dateObjectToUTCTimeString, isValidDate, utcTimeStringToDateObject} from "@/utils/timeutils";
import {toRaw} from "vue";


export const useBilladmStore = defineStore("billbooks", {
    state: () => {
        return {
            billbooks: [],
            bills: [],
            billTypes: new Map(),
            currentBook: '',
            showBillDisplayAside: true,
            showBillbookForm: false,
            billbookForm: {
                name: '',
                description: '',
            },
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
            showInitWorkspaceForm: false,
            showInitWorkspaceFormCloseButton: false,
            showAdvancedMenu: false,
            showMenu: true,
            workspaceState: {
                current: '',
                workspaces: new Map(),
            },
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
        async refreshBills(filters) {
            try {
                let res;
                if (filters) {
                    res = await window.appObject.getAllBillsByBookIDWithFilters(this.currentBook, filters);
                } else {
                    res = await window.appObject.getAllBillsByBookID(this.currentBook);
                }
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
                        creationTime: utcTimeStringToDateObject(item.creation_time),
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
            if (this.currentBook === '') {
                ElNotification({
                    type: 'warning',
                    message: '未选中任何账本',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
                return;
            }
            this.billForm.creationTime.setHours(12, 12, 12);
            let newItem = {
                money: Number(this.billForm.money),
                type: this.billForm.type,
                income: this.billForm.income,
                bookId: this.currentBook,
                description: this.billForm.description,
                tags: JSON.stringify(this.billForm.tags),
                creationTime: dateObjectToUTCTimeString(this.billForm.creationTime),
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
        async editOneBill() {
            this.billForm.creationTime.setHours(12, 12, 12);
            let newItem = {
                id: this.billForm.id,
                money: Number(this.billForm.money),
                type: this.billForm.type,
                income: this.billForm.income,
                description: this.billForm.description,
                tags: JSON.stringify(this.billForm.tags),
                creationTime: dateObjectToUTCTimeString(this.billForm.creationTime),
            };
            try {
                await window.appObject.editOneBill(newItem);
                ElNotification({
                    type: 'success',
                    message: '修改消费记录成功',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '  修改消费记录失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
        },
        async deleteBills(idList) {
            if (idList.length === 0) {
                ElNotification({
                    type: 'warning',
                    message: '未选中任何记录',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
                return;
            }
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
        async addOneBillbook() {
            if (toRaw(this.billbookForm.name) === '') {
                ElNotification({
                    type: 'error',
                    message: '账本名称不能为空',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
                return
            }
            try {
                let newItem = {
                    name: this.billbookForm.name,
                    description: this.billForm.description,
                };
                await window.appObject.addOneBillbook(newItem);
                ElNotification({
                    type: 'success',
                    message: '新增账本成功',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '新增账本失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
        },
        async deleteOneBillbook() {
            if (this.currentBook === BUILT_IN_BILLBOOK.id) {
                ElNotification({
                    type: 'warning',
                    message: '默认账本无法删除',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
                return;
            }
            try {
                await window.appObject.deleteOneBillbook(this.currentBook);
                this.currentBook = '';
                ElNotification({
                    type: 'success',
                    message: '删除账本成功',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '删除账本失败',
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
        toggleShowBillbookForm() {
            this.showBillbookForm = !this.showBillbookForm;
        },
        toggleShowAdvancedMenu() {
            this.showAdvancedMenu = !this.showAdvancedMenu;
        },
        toggleShowMenu() {
            this.showMenu = !this.showMenu;
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
            };
        },
        resetBillbookForm() {
            this.billbookForm = {
                name: '',
                description: '',
            };
        },
        async refreshBillTypes() {
            try {
                const res = await window.appObject.getAllBillTypes();
                res.forEach(item => {
                    this.billTypes.set(item.id, item.name);
                });
            } catch (err) {
                ElNotification({
                    type: 'error',
                    message: '获取消费类型失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
        },
        async handleTimeRangeChange() {
            let timeRange = toRaw(this.timeRange);
            if (Array.isArray(timeRange) && isValidDate(timeRange[0]) && isValidDate(timeRange[1])) {
                timeRange[0].setHours(0, 0, 0);
                timeRange[1].setHours(23, 59, 59);
                let filters = {
                    start_time: dateObjectToUTCTimeString(timeRange[0]),
                    end_time: dateObjectToUTCTimeString(timeRange[1]),
                };
                await this.refreshBills(filters);
            } else {
                await this.refreshBills();
            }
        },
        async chooseWorkspaceDirectory() {
            return await window.appObject.chooseWorkspaceDirectory();
        },
        async initWorkspace(workspaceDir) {
            if (toRaw(this.workspaceState.current) !== '' && workspaceDir.endsWith(toRaw(this.workspaceState.current))) {
                ElNotification({
                    type: 'warning',
                    message: '工作空间已打开',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
                return true;
            }
            const res = await window.appObject.initWorkspace(workspaceDir);
            if (res) {
                ElNotification({
                    type: 'success',
                    message: '工作空间初始化成功',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            } else {
                ElNotification({
                    type: 'error',
                    message: '工作空间初始化失败',
                    position: 'bottom-right',
                    duration: 2000,
                    offset: 40,
                });
            }
            return res;
        },
        // 新建或更换工作空间时从新工作空间空间重新初始化内容
        async refreshWorkspace() {
            await this.refreshBillbooks();
            await this.refreshBillTypes();
            this.setCurrentBook(BUILT_IN_BILLBOOK.id);
            await this.refreshBills();
        },
        async isFirstOpen() {
            return await window.appObject.isFirstOpen();
        },
        async refreshWorkspaceState() {
            this.workspaceState = await window.appObject.workspaceState();
        },
    }
})