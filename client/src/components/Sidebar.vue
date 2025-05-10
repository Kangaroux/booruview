<script setup lang="ts">
import store from "@/store";
import { ref } from "vue";
import TagList from "./TagList.vue";
import SearchForm from "./search/SearchForm.vue";

const sidebarClosed = ref(false);
const fetching = ref(false);

function doPostSearch() {
    if (fetching.value) {
        return;
    }

    fetching.value = true;
    store.searchPosts().finally(() => (fetching.value = false));
}
</script>

<template>
    <header
        class="sidebar-container"
        :class="{ 'sidebar-closed': sidebarClosed }"
    >
        <button class="toggle-btn" @click="sidebarClosed = !sidebarClosed">
            <i v-if="sidebarClosed" class="bi bi-chevron-right"></i>
            <i v-else class="bi bi-chevron-left"></i>
        </button>
        <nav class="sidebar-content" v-show="!sidebarClosed">
            <SearchForm
                @on-search="doPostSearch"
                @on-tag-select="(t) => store.addSearchTag(t)"
                :show-spinner="fetching"
            />

            <TagList :tags="store.searchTags()" />
        </nav>
    </header>
</template>

<style lang="scss" scoped>
@import "../assets/colors";

.sidebar-container {
    position: relative;
}

.sidebar-content {
    background-color: darken($bg-color, 2.5%);
    width: 350px;
    height: 100%;
    position: relative;
    padding: 10px;
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

    cursor: pointer;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.6);

    .sidebar-closed & {
        opacity: 0.5;

        &:hover {
            opacity: 1;
        }
    }
}
</style>
