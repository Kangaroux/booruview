<script setup lang="ts">
import store from "@/store";
import { computed, onMounted, onUnmounted, ref } from "vue";
import ContentTab from "./ContentTab.vue";
import InfoTab from "./InfoTab.vue";
import ScreenCover from "../ScreenCover.vue";
import { useIsVideo, useStationaryClick } from "@/composable";

type Tab = "content" | "info";
type PostNavInfo = { page: number; index: number } | null;

const { post } = defineProps<{ post: Post }>();
const currentTab = ref<Tab>("content");

function close() {
    store.fullscreenPost = null;
}

function onKeyDown(e: KeyboardEvent) {
    if (e.key === "Esc" || e.key === "Escape") {
        e.preventDefault();
        close();
    }
}

const isVideo = useIsVideo(() => post);

const tabClasses = computed(() => {
    // Center unless we're displaying an image and viewing the content tab.
    // Image viewing uses panzoom which breaks when the container is flexbox
    const centered = !(currentTab.value === "content" && !isVideo.value);
    return {
        "tab-centered": centered,
    };
});

const tabHandler = useStationaryClick(close);

const currentPostIndex = computed(() =>
    store.posts.get(store.currentPage)?.findIndex((p) => p.id === post.id),
);
const nextPost = computed<PostNavInfo>(() => {
    if (currentPostIndex.value == null) {
        return null;
    }

    const isLastPage = store.currentPage === store.maxPage();
    const isLastResult = currentPostIndex.value === store.resultsPerPage - 1;

    if (isLastPage && isLastResult) {
        return null;
    }

    if (isLastResult) {
        return {
            page: store.currentPage + 1,
            index: 0,
        };
    }

    return {
        page: store.currentPage,
        index: currentPostIndex.value + 1,
    };
});

const prevPost = computed<PostNavInfo>(() => {
    if (currentPostIndex.value == null) {
        return null;
    }

    const isFirstPage = store.currentPage === 1;
    const isFirstResult = currentPostIndex.value === 0;

    if (isFirstPage && isFirstResult) {
        return null;
    }

    if (isFirstResult) {
        return {
            page: store.currentPage - 1,
            index: store.resultsPerPage - 1,
        };
    }

    return {
        page: store.currentPage,
        index: currentPostIndex.value - 1,
    };
});

function showNextPost() {
    const nav = nextPost.value;

    if (nav === null) {
        return;
    }

    if (store.currentPage === nav.page) {
        store.fullscreenPost = store.posts.get(nav.page)![nav.index] || null;
    } else {
        store
            .nextPage()
            ?.then(
                () =>
                    (store.fullscreenPost =
                        store.posts.get(nav.page)![nav.index] || null),
            );
    }
}

function showPrevPost() {
    const nav = prevPost.value;

    if (nav === null) {
        return;
    }

    if (store.currentPage === nav.page) {
        store.fullscreenPost = store.posts.get(nav.page)![nav.index] || null;
    } else {
        store
            .prevPage()
            ?.then(
                () =>
                    (store.fullscreenPost =
                        store.posts.get(nav.page)![nav.index] || null),
            );
    }
}

onMounted(() => {
    document.addEventListener("keydown", onKeyDown);
});

onUnmounted(() => {
    document.removeEventListener("keydown", onKeyDown);
});
</script>

<template>
    <div class="fullscreen-viewer">
        <ScreenCover />
        <div class="viewer-container">
            <div
                class="tab"
                :class="tabClasses"
                @mousedown.self="tabHandler.mouseDown"
                @mouseup.self="tabHandler.mouseUp"
            >
                <KeepAlive>
                    <ContentTab v-if="currentTab == 'content'" :post="post" />
                    <InfoTab v-else-if="currentTab == 'info'" :post="post" />
                </KeepAlive>
            </div>
            <footer class="tab-menu">
                <button
                    class="menu-btn"
                    :class="{ active: currentTab == 'content' }"
                    @click="currentTab = 'content'"
                    title="view content"
                >
                    <i class="bi bi-image"></i>
                </button>
                <button
                    class="menu-btn"
                    :class="{ active: currentTab == 'info' }"
                    @click="currentTab = 'info'"
                    title="view tags"
                >
                    <i class="bi bi-info-circle"></i>
                </button>
                <button
                    class="menu-btn"
                    title="previous image"
                    @click="showPrevPost"
                    :disabled="prevPost === null"
                >
                    <i class="bi bi-arrow-left"></i>
                </button>
                <button
                    class="menu-btn"
                    title="next image"
                    @click="showNextPost"
                    :disabled="nextPost === null"
                >
                    <i class="bi bi-arrow-right"></i>
                </button>
                <button
                    class="menu-btn close-btn"
                    @click="close()"
                    title="close"
                >
                    <i class="bi bi-x-lg"></i>
                </button>
            </footer>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/breakpoints";

.fullscreen-viewer {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    overflow: hidden;
}

.viewer-container {
    height: 100%;
    position: relative;
    z-index: 2;
    width: 100%;
}

.tab {
    min-height: 0;
    height: 100%;
}

.tab-centered {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 10px;
    padding-bottom: 100px;
}

.tab-menu {
    display: flex;
    z-index: 2;
    background-color: rgba(0, 0, 0, 0.8);
    border-radius: 500px;
    padding: 0 10px;
    margin-bottom: 10px;
    box-shadow: 0 0 10px black;

    position: absolute;
    bottom: 10px;
    left: 50%;
    transform: translateX(-50%);
}

.menu-btn {
    background: none;
    border: none;
    color: white;
    font-size: 30px;
    text-shadow: 0 0 5px white;
    opacity: 0.5;
    transition: opacity 150ms;
    cursor: pointer;
    padding: 10px 15px;

    &:not(:disabled) {
        &:hover,
        &.active {
            opacity: 1;
        }
    }

    &:disabled {
        opacity: 0.2;
        cursor: default;
    }
}

.close-btn {
    color: #bb9fce;
    text-shadow: 0 0 5px #bb9fce;
    padding-right: 10px;
}
</style>
