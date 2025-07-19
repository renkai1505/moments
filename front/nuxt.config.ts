// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: '2024-04-03',
    devtools: {enabled: false},
    modules: ["@nuxt/ui", '@nuxt/icon', '@nuxtjs/color-mode', '@vueuse/nuxt', 'dayjs-nuxt'],
    ssr: false,
    dayjs: {
        locales: ['zh'],
        defaultLocale: 'zh'
    },
    icon: {
        clientBundle: {
            scan: {
                globInclude: ['**/*.{vue,jsx,tsx}', 'node_modules/@nuxt/ui/**/*.js'],
                globExclude: ['.*', 'coverage', 'test', 'tests', 'dist', 'build'],
            },
        },
    },
    tailwindcss: {
        safelist: [
            'grid-cols-1',
            'grid-cols-3',
        ]
    },
    vue: {
        compilerOptions: {
            isCustomElement: (tag:string) => ['meting-js'].includes(tag),
        },
    },
    app: {
        head: {
            title: '儿童成长相册 - Kids Album',
            meta: [
                { name: "viewport", content: "width=device-width, initial-scale=1, user-scalable=no" },
                { charset: "utf-8" },
                { name: "description", content: "专为记录儿童成长而设计的相册应用，让每个珍贵瞬间都值得珍藏！" },
                { name: "keywords", content: "儿童相册,成长记录,时间轴,照片管理" },
            ],
            link: [
                {href: `/css/APlayer.min.css`, rel: 'stylesheet'},
            ],
            script: [
                {src: `/js/APlayer.min.js`, type: 'text/javascript', async: true, defer: true},
                {src: `/js/Meting.min.js`, type: 'text/javascript', async: true, defer: true},
                {src: `/js/main.js`, type: 'text/javascript', async: true, defer: true},
            ]
        }
    },
    vite: {
        server: {
            proxy: {
                "/api": {
                    target: "http://localhost:37892",
                    // changeOrigin: true,
                },
                "/upload": {
                    target: "http://localhost:37892",
                    // changeOrigin: true,
                },
                "/rss": {
                    target: "http://localhost:37892",
                    // changeOrigin: true,
                },
                "/swagger": {
                    target: "http://localhost:37892",
                    // changeOrigin: true,
                },
            },
        },
        build: {
            rollupOptions: {
                output: {
                    hashCharacters: 'base36'
                }
            }
        }
    }
})