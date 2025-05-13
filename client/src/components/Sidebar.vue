<script setup lang="ts">
import store from "@/store";
import { ref } from "vue";
import TagList from "./TagList.vue";
import SearchForm from "./search/SearchForm.vue";

defineEmits(["toggle"]);
const { closed } = defineProps<{ closed: boolean }>();
const fetching = ref(false);

function doPostSearch() {
    if (fetching.value) {
        return;
    }

    fetching.value = true;
    store.searchPosts().finally(() => (fetching.value = false));
}

function onTagSelect(tag: Tag, negated: boolean) {
    if (!negated) {
        store.search.query.includeTag(tag);
    } else {
        store.search.query.excludeTag(tag);
    }
}
</script>

<template>
    <header class="sidebar-container">
        <button class="toggle-btn" @click="$emit('toggle')">
            <i v-if="closed" class="bi bi-chevron-right"></i>
            <i v-else class="bi bi-chevron-left"></i>
        </button>
        <nav class="sidebar-content" v-show="!closed">
            <SearchForm
                @on-search="doPostSearch"
                @on-tag-select="onTagSelect"
                :show-spinner="fetching"
            />

            <TagList
                :jiggle="true"
                :excludeTags="[...store.search.query.exclude.values()]"
                :includeTags="[...store.search.query.include.values()]"
            />
        </nav>
    </header>
</template>

<style lang="scss" scoped>
@import "../assets/colors";

.sidebar-container {
    position: relative;

    @media (max-width: 600px) {
        .sidebar-open & {
            width: 100%;
        }
    }
}

.sidebar-content {
    background-color: darken($bg-color, 2.5%);
    width: 350px;
    height: 100%;
    position: relative;
    padding: 10px;

    @media (max-width: 600px) {
        .sidebar-open & {
            width: 100%;
        }
    }
}

.toggle-btn {
    $btn-color: #342b3a;
    $border-color: lighten($btn-color, 20%);
    $btn-height: 70px;

    background-color: $btn-color;
    border: 1px solid $border-color;
    color: #bb9fce;

    font-size: 24px;
    height: $btn-height;
    padding: 0 8px;
    border-left: none;
    border-radius: 0 4px 4px 0;

    position: absolute;
    left: 100%;
    top: calc(50% - $btn-height / 2);
    z-index: 1;

    cursor: pointer;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.6);

    .sidebar-closed & {
        opacity: 0.5;

        // Fixes the toggle btn not dimming when pressed on mobile
        @media (any-hover: hover) {
            &:hover {
                opacity: 1;
            }
        }
    }

    // Move the toggle btn to the left side of the screen on mobile
    @media (max-width: 600px) {
        .sidebar-open & {
            left: 0;
        }
    }
}
</style>
