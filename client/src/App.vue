<script setup lang="ts">
import { computed, useTemplateRef, watch } from "vue";
import store from "@/store";
import PostContainer from "./components/PostContainer.vue";
import Sidebar from "./components/sidebar/Sidebar.vue";
import FullscreenView from "./components/fullscreen-view/FullscreenView.vue";
import Router from "./Router.vue";
import ContentWarning from "./components/ContentWarning.vue";

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
        <ContentWarning v-if="!hasConsented" />
        <Sidebar
            :closed="store.sidebarClosed"
            @toggle="store.sidebarClosed = !store.sidebarClosed"
        />
        <main ref="main">
            <FullscreenView v-if="store.fullscreenPost !== null" />
            <div v-if="store.hasResults()">
                <PostContainer :posts="store.postsForCurrentPage() || []" />
                <footer class="page-nav">
                    <button
                        class="btn-primary btn-rounded"
                        @click="store.prevPage()"
                        v-if="store.currentPage > 1"
                    >
                        <i class="bi bi-arrow-left"></i> prev
                    </button>
                    <button
                        class="btn-primary btn-rounded"
                        @click="store.nextPage()"
                        v-if="store.currentPage < store.maxPage()"
                    >
                        next <i class="bi bi-arrow-right"></i>
                    </button>
                    <p>
                        page {{ store.currentPage }} of
                        {{ store.maxPage() }} ({{ store.totalPostCount }}
                        results)
                    </p>
                </footer>
            </div>
        </main>
    </div>
</template>

<style scoped lang="scss">
@import "assets/breakpoints";
@import "assets/buttons";
@import "assets/colors";

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

.page-nav {
    margin-top: 40px;
    text-align: center;

    .btn-primary:nth-of-type(2) {
        margin-left: 10px;
    }
}
</style>
