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

    store
        .searchPosts()
        .then(() => {
            if (store.settings.closeSidebarOnSearch) {
                store.sidebarClosed = true;
            }
        })
        .catch(() => store.posts.clear());
}

function onTagSelect(tag: Tag, negated: boolean) {
    if (!negated) {
        store.query.includeTag(tag);
    } else {
        store.query.excludeTag(tag);
    }

    // Slight hack: try looking up unknown tags and replace it with the real version
    if (tag.type === "unknown") {
        store.loadTags([tag.name]).then(() => {
            const real = store.cachedTags.get(tag.name);

            if (real === undefined) {
                return;
            }

            store.query.removeTag(tag);

            if (!negated) {
                store.query.includeTag(real);
            } else {
                store.query.excludeTag(real);
            }
        });
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
                        :tags="store.query.all()"
                        :show-checkmark="false"
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
    min-height: 0;

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
