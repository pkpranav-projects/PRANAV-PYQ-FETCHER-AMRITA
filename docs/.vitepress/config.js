import { defineConfig } from "vitepress";

export default defineConfig({
    base: "/PRANAV-PYQ-FETCHER-AMRITA/",
    title: "PRANAV PYQ FETCHER AMRITA",
    titleTemplate: ":title | PRANAV PYQ FETCHER AMRITA",
    description: "An interactive Terminal CLI tool for accessing Amrita University PYQs.",
    cleanUrls: true,

    themeConfig: {
        nav: [
            { text: "Home", link: "/" },
            { text: "Docs", link: "/getting-started" },
        ],
        search: {
            provider: "local",
            placeholder: "Search",
        },
        sidebar: [
            {
                text: "Docs",
                items: [
                    { text: "Getting Started", link: "/getting-started" },
                    { text: "Testing", link: "/testing" },
                ],
            },
        ],
        socialLinks: [
            {
                icon: "github",
                link: "https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA",
            },
        ],
    },
});
