import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {createSvgIconsPlugin} from "vite-plugin-svg-icons";
import path from "path";


export default defineConfig({
    plugins: [
        vue(),
        createSvgIconsPlugin({
            iconDirs: [path.resolve(__dirname, 'src/icons')],
            symbolId: 'icon-[dir]-[name]'
        })
    ],
    resolve: {
        alias: [
            {
                find: '@',
                replacement: path.resolve(__dirname, './src')
            }
        ]
    }
})