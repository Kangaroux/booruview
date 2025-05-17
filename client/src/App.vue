<script setup lang="ts">
import { useTemplateRef, watch } from "vue";
import store from "@/store";
import PostContainer from "./components/PostContainer.vue";
import Sidebar from "./components/sidebar/Sidebar.vue";
import FullscreenView from "./components/fullscreen-view/FullscreenView.vue";

const mainContentContainer = useTemplateRef("main-content");

watch(
    () => [mainContentContainer, store.posts],
    () => {
        if (!mainContentContainer.value) {
            return;
        }
        mainContentContainer.value.scrollTop = 0;
    },
    {
        flush: "post",
        deep: true,
    },
);
</script>

<template>
    <div
        class="app"
        :class="{
            'sidebar-closed': store.sidebarClosed,
            'sidebar-open': !store.sidebarClosed,
        }"
    >
        <Sidebar
            :closed="store.sidebarClosed"
            @toggle="store.sidebarClosed = !store.sidebarClosed"
        />
        <main>
            <FullscreenView v-if="store.fullscreenPost !== null" />
            <div
                class="main-content"
                ref="main-content"
                v-if="store.hasResults()"
            >
                <PostContainer :posts="store.postsForCurrentPage() || []" />
                <footer>
                    <p>
                        page {{ store.currentPage }} of
                        {{ store.maxPage() }} ({{ store.totalPostCount }}
                        results)
                    </p>
                    <button
                        @click="store.prevPage()"
                        v-if="store.currentPage > 1"
                    >
                        &lt;&lt; prev page
                    </button>
                    <button
                        @click="store.nextPage()"
                        v-if="store.currentPage < store.maxPage()"
                    >
                        next page &gt;&gt;
                    </button>
                </footer>
            </div>
        </main>
    </div>
</template>

<style scoped lang="scss">
@import "assets/breakpoints";
@import "assets/colors";

.app {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;
}

main {
    flex: 1;

    @media (max-width: $mobile-width) {
        .sidebar-open & {
            display: none;
        }
    }
}

.main-content {
    height: 100%;
    overflow-y: scroll;
}
</style>
