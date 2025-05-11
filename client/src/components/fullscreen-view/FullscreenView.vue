<script setup lang="ts">
import store from "@/store";
import { onMounted, onUnmounted, ref, watchEffect } from "vue";
import TagList from "../TagList.vue";
import ImageTab from "./ImageTab.vue";

const drawerOpen = ref(false);
const tags = ref<Tag[]>([]);

function close() {
    store.fullscreenPost = null;
}

function onKeyDown(e: KeyboardEvent) {
    if (e.key === "Esc" || e.key === "Escape") {
        e.preventDefault();
        close();
    }
}

watchEffect(() => {
    if (store.fullscreenPost === null) {
        return;
    }

    store.tagsForPost(store.fullscreenPost).then((val) => (tags.value = val));
});

onMounted(() => {
    document.addEventListener("keydown", onKeyDown);
});

onUnmounted(() => {
    document.removeEventListener("keydown", onKeyDown);
});
</script>

<template>
    <div class="fullscreen-viewer">
        <div class="screen-cover" @click="close()"></div>
        <div class="outer-container">
            <KeepAlive>
                <ImageTab />
            </KeepAlive>
            <div class="info-drawer" :class="{ 'drawer-open': drawerOpen }">
                <div class="drawer-btn" @click="drawerOpen = !drawerOpen">
                    <i class="bi bi-info-circle"></i>
                </div>
                <TagList v-if="drawerOpen" :jiggle="false" :tags="tags" />
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.fullscreen-viewer {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
}

.screen-cover {
    position: absolute;
    z-index: 1;
    background-color: rgba(0, 0, 0, 0.95);
    width: 100%;
    height: 100%;
}

.outer-container {
    height: 100%;
    margin: 0 100px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 10px;
}

.info-drawer {
    z-index: 3;
    position: relative;

    &.drawer-open {
        bottom: 0;
    }

    & .drawer-btn {
        font-size: 30px;
        padding: 10px 30px;
        text-shadow: 0 0 5px white;
        opacity: 0.6;
        transition: opacity 200ms;
        cursor: pointer;

        &:hover {
            opacity: 1;
        }
    }
}
</style>
