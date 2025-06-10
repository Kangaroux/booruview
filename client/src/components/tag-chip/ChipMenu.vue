<script setup lang="ts">
import store from "@/store";
import { computed, useTemplateRef } from "vue";
import { useDismiss } from "@/composable";

const containerRef = useTemplateRef("container");

const isIncluded = computed(() => {
    if (store.tagMenu === null) {
        return false;
    }

    return store.query.include.has(store.tagMenu.tag.name);
});

const isExcluded = computed(() => {
    if (store.tagMenu === null) {
        return false;
    }

    return store.query.exclude.has(store.tagMenu.tag.name);
});

function closeMenu() {
    store.tagMenu = null;
}

useDismiss([containerRef.value, store.tagMenu?.ref || null], closeMenu);

const menuPosition = computed(() => {
    const ref = store.tagMenu?.ref;

    if (ref == null) {
        return;
    }

    const windowHeight = window.innerHeight;
    const allowedMargin = 125;
    const { left, bottom, top } = ref.getBoundingClientRect();

    // Anchor the menu to the bottom of the chip if there is enough space in the viewport
    if (windowHeight - bottom >= allowedMargin) {
        return {
            left: left + "px",
            top: bottom + "px",
        };
    }

    // If space is limited move the anchor to the top of the chip
    const bottomToTop = windowHeight - top;

    return {
        left: left + "px",
        bottom: bottomToTop + "px",
    };
});

function onAdd() {
    if (!store.tagMenu) {
        return;
    }

    store.query.includeTag(store.tagMenu.tag);
    closeMenu();
}

function onExclude() {
    if (!store.tagMenu) {
        return;
    }

    store.query.excludeTag(store.tagMenu.tag);
    closeMenu();
}

function onRemove() {
    if (!store.tagMenu) {
        return;
    }

    store.query.removeTag(store.tagMenu.tag);
    closeMenu();
}
</script>

<template>
    <div class="options" ref="container" :style="menuPosition">
        <button
            class="btn-primary option-btn"
            v-if="isExcluded || !isIncluded"
            @click="onAdd"
        >
            <i class="bi bi-plus-lg"></i> include
        </button>
        <button
            class="btn-primary option-btn"
            v-if="!isExcluded || isIncluded"
            @click="onExclude"
        >
            <i class="bi bi-dash-lg"></i> exclude
        </button>
        <button
            class="btn-primary option-btn"
            v-if="isExcluded || isIncluded"
            @click="onRemove"
        >
            <i class="bi bi-x-lg"></i> remove
        </button>
        <button class="btn-primary option-btn blacklist" disabled>
            <i class="bi bi-ban"></i> blacklist
        </button>
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/buttons";
@import "@/assets/colors";

.options {
    position: absolute;
    display: flex;
    flex-direction: column;
    z-index: 999;
    width: max-content;
    box-shadow: 0 0 10px black;
}

.option-btn {
    text-align: left;

    .bi {
        margin-right: 5px;
    }

    &:first-of-type {
        border-radius: 4px 4px 0 0;
    }

    &:last-of-type {
        border-radius: 0 0 4px 4px;
    }

    &:not(:last-of-type) {
        border-bottom: none;
    }
}

.blacklist {
    color: #ff5d5d;
}
</style>
