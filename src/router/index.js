import {createRouter, createWebHashHistory} from 'vue-router'
import BillDisplay from '@/components/BillDisplay.vue'
import AnalysisDisplay from '@/components/AnalysisDisplay.vue'
import PropertyDisplay from '@/components/PropertyDisplay.vue'
import SettingDisplay from '@/components/SettingDisplay.vue'
import EmptyDisplay from '@/components/EmptyDisplay.vue'

const routes = [
    {
        path: '/bill',
        name: 'bill',
        component: BillDisplay,
    },
    {
        path: '/analysis',
        name: 'analysis',
        component: EmptyDisplay,
    },
    {
        path: '/property',
        name: 'property',
        component: EmptyDisplay,
    },
    {
        path: '/setting',
        name: 'setting',
        component: EmptyDisplay,
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router