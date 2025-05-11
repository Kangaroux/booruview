<script setup lang="ts">
import store from "@/store";
import { computed, onMounted, onUnmounted, ref } from "vue";

const fitHeight = ref(true);

const url = computed(() => {
    if (store.postFocus === null) {
        return "";
    }

    return store.postFocus.image_url;
});

function close() {
    store.postFocus = null;
}

function onKeyDown(e: KeyboardEvent) {
    if (e.key === "Esc" || e.key === "Escape") {
        e.preventDefault();
        close();
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
    <div class="post-focus">
        <div class="screen-cover" @click="close()"></div>
        <div class="focus-container">
            <div class="content-container">
                <img
                    class="content"
                    :class="{
                        'fit-height': fitHeight,
                        'fit-width': !fitHeight,
                    }"
                    :src="url"
                    @click="fitHeight = !fitHeight"
                />
            </div>
            <div class="info-drawer">
                <div class="drawer-btn">
                    <i class="bi bi-info-circle"></i>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.post-focus {
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

.focus-container {
    margin: 0 100px;
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100%;
}

.content-container {
    z-index: 2;
    height: 100%;
    overflow-y: scroll;
    line-height: 0;

    scrollbar-width: none;

    &::-webkit-scrollbar {
        display: none;
    }
}

.content {
    &.fit-height {
        max-height: 100%;
        width: auto;

        cursor: zoom-in;
    }

    &.fit-width {
        max-width: 100%;
        height: auto;

        cursor: zoom-out;
    }
}

.info-drawer {
    z-index: 3;

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
