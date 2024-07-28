import { createRouter, createWebHashHistory } from 'vue-router'
import BillDisplay from '../components/BillDisplay.vue'
import AnalysisDisplay from '../components/AnalysisDisplay.vue'
import PropertyDisplay from '../components/PropertyDisplay.vue'
import SettingDisplay from '../components/SettingDisplay.vue'

const routes = [
    {
        path: '/bill',
        name: 'bill',
        component: BillDisplay,
    },
    {
        path: '/analysis',
        name: 'analysis',
        component: AnalysisDisplay,
    },
    {
        path: '/property',
        name: 'property',
        component: PropertyDisplay,
    },
    {
        path: '/setting',
        name: 'setting',
        component: SettingDisplay,
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router