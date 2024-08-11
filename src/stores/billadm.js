/*存储账本信息*/
import {defineStore} from "pinia";
import {BUILT_IN_BILLBOOK} from "@/utils/constants";
import {dateObjectToUTCTimeString, isValidDate, utcTimeStringToDateObject} from "@/utils/timeutils";
import {toRaw} from "vue";
import {notify} from "@/utils/notify";


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
            showHelpMenu: false,
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
                notify('error', '获取账本失败');
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
                notify('error', '获取消费记录失败');
            }
        },
        async addOneBill() {
            if (this.currentBook === '') {
                notify('warning', '未选中任何账本');
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
                notify('success', '新增消费记录成功');
            } catch (err) {
                notify('error', '新增消费记录失败');
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
                notify('success', '修改消费记录成功');
            } catch (err) {
                notify('error', '修改消费记录失败');
            }
        },
        async deleteBills(idList) {
            if (idList.length === 0) {
                notify('warning', '未选中任何记录');
                return;
            }
            try {
                await window.appObject.deleteBills(idList);
                notify('success', '删除消费记录成功');
            } catch (err) {
                notify('error', '删除消费记录失败');
            }
        },
        async addOneBillbook() {
            if (toRaw(this.billbookForm.name) === '') {
                notify('error', '账本名称不能为空');
                return
            }
            try {
                let newItem = {
                    name: this.billbookForm.name,
                    description: this.billForm.description,
                };
                await window.appObject.addOneBillbook(newItem);
                notify('success', '新增账本成功');
            } catch (err) {
                notify('error', '新增账本失败');
            }
        },
        async deleteOneBillbook() {
            if (this.currentBook === BUILT_IN_BILLBOOK.id) {
                notify('warning', '默认账本无法删除');
                return;
            }
            try {
                await window.appObject.deleteOneBillbook(this.currentBook);
                this.currentBook = '';
                notify('success', '删除账本成功');
            } catch (err) {
                notify('error', '删除账本失败');
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
        toggleShowHelpMenu() {
            this.showHelpMenu = !this.showHelpMenu;
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
                notify('error', '获取消费类型失败');
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
                notify('warning', '工作空间已打开');
                return true;
            }
            const res = await window.appObject.initWorkspace(workspaceDir);
            if (res) {
                notify('success', '工作空间初始化成功');
            } else {
                notify('error', '工作空间初始化失败');
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
        async removeWorkspace(workspaceDir) {
            try {
                const ret = await window.appObject.removeWorkspace(workspaceDir);
                if (ret.status === 'opened' && ret.newWorkspaceName.length > 0) {
                    notify('success', `工作空间已移除，自动切换到另一个工作空间：${ret.newWorkspaceName}`);
                } else {
                    notify('success', '工作空间已移除');
                }
                return true;
            } catch (e) {
                notify('error', '工作空间移除失败');
                return false;
            }
        },
        async isFirstOpen() {
            return await window.appObject.isFirstOpen();
        },
        async refreshWorkspaceState() {
            this.workspaceState = await window.appObject.workspaceState();
        },
    }
})