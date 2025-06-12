<script setup lang="ts">
import { computed, useTemplateRef, watch } from "vue";
import store from "@/store";
import PostContainer from "./components/PostContainer.vue";
import Sidebar from "./components/sidebar/Sidebar.vue";
import FullscreenView from "./components/fullscreen-view/FullscreenView.vue";
import Router from "./Router.vue";
import ContentWarning from "./components/ContentWarning.vue";
import Footer from "./components/Footer.vue";
import NoResults from "./components/NoResults.vue";
import ChipMenu from "./components/tag-chip/ChipMenu.vue";

const mainContainer = useTemplateRef("main");

watch(
    () => [mainContainer, store.posts],
    () => {
        // Scroll needs to be deferred in order to work on mobile
        setTimeout(() => {
            if (mainContainer.value === null) {
                return;
            }
            mainContainer.value.scrollTop = 0;
        }, 0);
    },
    {
        flush: "post",
        deep: true,
    },
);

const hasConsented = computed(() => {
    if (store.settings.consented) {
        return true;
    }

    // Don't show consent modal for search engine crawlers
    const crawlers = /Googlebot|Bingbot|DuckDuckbot/;
    const isCrawler = crawlers.exec(navigator.userAgent) !== null;
    return isCrawler;
});
</script>

<template>
    <Router />
    <div
        class="app"
        :class="{
            'sidebar-closed': store.sidebarClosed,
            'sidebar-open': !store.sidebarClosed,
        }"
    >
        <ChipMenu v-if="store.tagMenu !== null" />
        <ContentWarning v-if="!hasConsented" />

        <FullscreenView
            v-if="store.fullscreenPost !== null"
            :post="store.fullscreenPost"
        />

        <Sidebar
            :closed="store.sidebarClosed"
            @toggle="store.sidebarClosed = !store.sidebarClosed"
        />
        <main ref="main">
            <template v-if="store.hasSearched">
                <NoResults v-if="store.totalPostCount === 0" />
                <template v-else>
                    <PostContainer :posts="store.postsForCurrentPage() || []" />
                    <Footer v-if="!store.fetchingPosts" />
                </template>
            </template>
        </main>
    </div>
</template>

<style scoped lang="scss">
@import "@/assets/breakpoints";
@import "@/assets/colors";

.app {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;

    @media (max-width: $mobile-width) {
        &.sidebar-closed {
            flex-direction: column;
        }
    }
}

main {
    flex: 1;
    min-height: 0;
    overflow-y: scroll;

    @media (max-width: $mobile-width) {
        .sidebar-open & {
            display: none;
        }
    }
}
</style>
