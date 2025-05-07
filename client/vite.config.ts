import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
import vueDevTools from "vite-plugin-vue-devtools";

// https://vite.dev/config/
export default defineConfig({
    plugins: [vue(), vueJsx(), vueDevTools()],
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./src", import.meta.url)),
            "~": fileURLToPath(new URL("./node_modules", import.meta.url)),
        },
    },
    css: {
        preprocessorOptions: {
            scss: {
                // I don't want to use your shitty verbose color API
                // give me back my lighten and darken >:(
                silenceDeprecations: [
                    "color-functions",
                    "global-builtin",
                    "import",
                ],
            },
        },
    },
});
