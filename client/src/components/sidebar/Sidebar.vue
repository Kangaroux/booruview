<script setup lang="ts">
import store from "@/store";
import TagList from "../TagList.vue";
import SearchForm from "../search/SearchForm.vue";
import TabContainer from "./TabContainer.vue";

defineEmits(["toggle"]);

const { closed } = defineProps<{ closed: boolean }>();

function doPostSearch() {
    if (store.fetchingPosts) {
        return;
    }

    store.currentPage = 1;
    store.posts.clear();

    store.searchPosts().then(() => {
        if (store.settings.closeSidebarOnSearch) {
            store.sidebarClosed = true;
        }
    });
}

function onTagClick(tag: Tag) {
    if (store.query.include.has(tag.name)) {
        store.query.excludeTag(tag);
    } else {
        store.query.removeTag(tag);
    }
}

function onTagSelect(tag: Tag, negated: boolean) {
    if (!negated) {
        store.query.includeTag(tag);
    } else {
        store.query.excludeTag(tag);
    }
}
</script>

<template>
    <header class="sidebar-container">
        <div class="sidebar-header">
            <button class="toggle-btn" @click="$emit('toggle')">
                <i class="bi bi-list"></i>
            </button>
        </div>

        <div class="sidebar-content" v-show="!closed">
            <div class="search">
                <SearchForm
                    @on-search="doPostSearch"
                    @on-tag-select="onTagSelect"
                    :show-spinner="store.fetchingPosts"
                />

                <div class="taglist-container">
                    <TagList
                        :jiggle="true"
                        :excludeTags="[...store.query.exclude.values()]"
                        :includeTags="[...store.query.include.values()]"
                        @click="onTagClick"
                    />
                </div>
            </div>

            <TabContainer />
        </div>
    </header>
</template>

<style lang="scss" scoped>
@import "@/assets/breakpoints";
@import "@/assets/colors";

.sidebar-container {
    display: flex;
    flex-direction: column;
    background-color: $color-sidebar;
    height: 100%;

    @media (max-width: $mobile-width) {
        .sidebar-open & {
            width: 100%;
        }

        .sidebar-closed & {
            height: auto;
        }
    }
}

.sidebar-content {
    width: 450px;
    margin-top: 10px;
    position: relative;
    display: flex;
    flex-direction: column;
    flex: 1;

    @media (max-width: $mobile-width) {
        .sidebar-open & {
            width: 100%;
        }
    }
}

.search {
    padding: 0 10px;
    min-height: 0;
    flex: 1;
    display: flex;
    flex-direction: column;
}

.taglist-container {
    overflow-y: scroll;
    flex: 1;
}

.toggle-btn {
    background: none;
    border: none;
    font-size: 40px;
    cursor: pointer;

    color: $color-primary-lighter;
}
</style>
