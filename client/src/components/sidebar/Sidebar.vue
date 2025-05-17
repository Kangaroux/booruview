<script setup lang="ts">
import store from "@/store";
import { ref } from "vue";
import TagList from "../TagList.vue";
import SearchForm from "../search/SearchForm.vue";
import SearchHelp from "./SearchHelp.vue";
import TabContainer from "./TabContainer.vue";

defineEmits(["toggle"]);

const { closed } = defineProps<{ closed: boolean }>();
const fetching = ref(false);

function onCloseHelp() {
    store.settings.helpClosed = true;
    store.settings.save();
}

function doPostSearch() {
    if (fetching.value) {
        return;
    }

    fetching.value = true;
    store
        .searchPosts({ closeSidebar: true })
        .finally(() => (fetching.value = false));
}

function onTagClick(tag: Tag) {
    if (store.search.query.include.has(tag.name)) {
        store.search.query.excludeTag(tag);
    } else {
        store.search.query.removeTag(tag);
    }
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
        <div class="sidebar-content" v-show="!closed">
            <div class="search">
                <SearchHelp
                    v-if="!store.settings.helpClosed"
                    @on-close="onCloseHelp"
                />
                <SearchForm
                    @on-search="doPostSearch"
                    @on-tag-select="onTagSelect"
                    :show-spinner="fetching"
                />

                <TagList
                    :jiggle="true"
                    :excludeTags="[...store.search.query.exclude.values()]"
                    :includeTags="[...store.search.query.include.values()]"
                    @click="onTagClick"
                />
            </div>

            <TabContainer />
        </div>
    </header>
</template>

<style lang="scss" scoped>
@import "@/assets/breakpoints";
@import "@/assets/colors";

.sidebar-container {
    position: relative;

    @media (max-width: $mobile-width) {
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
    display: flex;
    flex-direction: column;

    @media (max-width: $mobile-width) {
        .sidebar-open & {
            width: 100%;
        }
    }
}

.search {
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
    z-index: 1;

    cursor: pointer;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.6);

    .sidebar-closed & {
        opacity: 0.5;

        // Don't change opacity on mobile (fixes button getting stuck in hover state)
        @media (pointer: fine) {
            &:hover {
                opacity: 1;
            }
        }
    }

    // Move the toggle btn to the bottom left side of the screen on mobile
    @media (max-width: $mobile-width) {
        top: auto;
        bottom: 50px;

        .sidebar-open & {
            left: 0;
        }
    }
}
</style>
