import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import { pluginExposeRenderer } from './vite.base.config.mjs';
import path from "path";

// https://vitejs.dev/config
export default defineConfig((env) => {
  /** @type {import('vite').ConfigEnv<'renderer'>} */
  const forgeEnv = env;
  const { root, mode, forgeConfigSelf } = forgeEnv;
  const name = forgeConfigSelf.name ?? '';

  /** @type {import('vite').UserConfig} */
  return {
    root,
    mode,
    base: './',
    build: {
      outDir: `.vite/renderer/${name}`,
    },
    plugins: [
      vue(),
      pluginExposeRenderer(name),
      createSvgIconsPlugin({
        iconDirs: [path.resolve(__dirname, 'src/icons')],
        symbolId: 'icon-[dir]-[name]'
      })
    ],
    resolve: {
      preserveSymlinks: true,
    },
    clearScreen: false,
  };
});
