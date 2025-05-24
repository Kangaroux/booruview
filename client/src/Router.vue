<script setup lang="ts">
import { onMounted } from "vue";
import store from "./store";

function parsePage(raw: string | null): number {
    if (raw === null) {
        return 1;
    }

    const val = parseInt(raw);
    if (!Number.isSafeInteger(val) || val < 1) {
        return 1;
    }

    return val;
}

function parseQuery(raw: string | null) {
    if (raw === null || raw.length === 0) {
        return;
    }

    const tagNames = raw.split(",").filter((v) => v.length > 0);

    store.loadTags(tagNames).then(() => {
        for (let name of tagNames) {
            const negate = name[0] === "-";
            if (negate) {
                name = name.slice(1);
            }

            let tag = store.cachedTags.get(name);
            if (tag === undefined) {
                tag = {
                    count: 0,
                    name: name,
                    type: "unknown",
                };
            }

            if (negate) {
                store.query.excludeTag(tag);
            } else {
                store.query.includeTag(tag);
            }
        }
    });
}

function loadQueryParams(): boolean {
    const params = new URLSearchParams(window.location.hash.replace(/^#/, ""));
    const page = params.get("page");
    const query = params.get("q");
    store.currentPage = parsePage(page);
    parseQuery(query);

    return page !== null && query !== null;
}

function onRouteChange() {
    const lastPage = store.currentPage;
    const lastSearch = store.query.copy();

    const paramsLoaded = loadQueryParams();
    if (!paramsLoaded) {
        return;
    }

    const pageChanged = store.currentPage !== lastPage;
    const searchChanged = !lastSearch.equals(store.query);

    if (searchChanged) {
        store.posts.clear();
    }

    if (pageChanged || searchChanged) {
        store.searchPosts();
    }
}

function onPageLoad() {
    const paramsLoaded = loadQueryParams();

    if (paramsLoaded) {
        store.searchPosts();
    }
}

window.addEventListener("hashchange", onRouteChange);
onMounted(onPageLoad);
</script>

<template></template>
