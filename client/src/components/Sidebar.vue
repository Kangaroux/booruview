<script setup lang="ts">
import store from "@/store";
import { ref } from "vue";
import TagList from "./TagList.vue";
import TagSearch from "./search/TagSearch.vue";

const tags = ref<Tag[]>([]);
const sidebarClosed = ref(false);
const fetching = ref(false);

function doSearch() {
    if (fetching.value) {
        return;
    }

    fetching.value = true;

    fetch(
        "/api/posts?q=" +
            encodeURIComponent(tags.value.map((t) => t.name).join(" ")),
    )
        .then((resp) => {
            resp.json().then((json) => {
                store.posts = json.results;
                console.log(json);
            });
        })
        .catch((err) => console.error(err))
        .finally(() => (fetching.value = false));
}
</script>

<template>
    <button class="toggle-btn" @click="sidebarClosed = !sidebarClosed">
        <i v-if="sidebarClosed" class="bi bi-chevron-right"></i>
        <i v-else class="bi bi-chevron-left"></i>
    </button>
    <nav class="sidebar" v-if="!sidebarClosed">
        <TagSearch
            @on-search="doSearch"
            @on-submit="(t) => (tags = tags.concat(t))"
            :exclude-tags="tags"
        />
        <button
            class="search-btn"
            type="submit"
            @click="doSearch"
            :disabled="fetching"
        >
            <span v-if="!fetching">search</span>
            <span v-else class="spinner"></span>
        </button>

        <TagList :tags="tags" />
    </nav>
</template>

<style lang="scss" scoped>
@import "../assets/colors";

.sidebar {
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
}
</style>
