<script setup lang="ts">
import { onMounted } from "vue";
import store from "./store";

function parsePage(raw: string): number {
    const val = parseInt(raw);
    if (!Number.isSafeInteger(val) || val < 1) {
        return 1;
    }

    return val;
}

// Promise resolves with the number of tags in the query
function parseQuery(raw: string): Promise<number> {
    return new Promise<number>((resolve, reject) => {
        if (raw.length === 0) {
            resolve(0);
            return;
        }

        const tagNames = raw.split(",").filter((v) => v.length > 0);

        store
            .loadTags(tagNames)
            .then(() => {
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

                resolve(tagNames.length);
            })
            .catch(reject);
    });
}

// Loads the query params into the store. Resolves if the query is non-empty
function loadQueryParams(): Promise<void> {
    return new Promise((resolve, reject) => {
        const params = new URLSearchParams(
            window.location.hash.replace(/^#/, ""),
        );
        const page = params.get("page") || "1";
        const query = params.get("q") || "";
        store.currentPage = parsePage(page);
        parseQuery(query)
            .then((tagCount) => (tagCount > 0 ? resolve() : reject()))
            .catch(reject);
    });
}

function onRouteChange() {
    const lastPage = store.currentPage;
    const lastSearch = store.query.copy();

    loadQueryParams()
        .then(() => {
            const pageChanged = store.currentPage !== lastPage;
            const searchChanged = !lastSearch.equals(store.query);

            if (searchChanged) {
                store.posts.clear();
            }

            if (pageChanged || searchChanged) {
                store.searchPosts();
            }
        })
        .catch(() => {});
}

function onPageLoad() {
    loadQueryParams()
        .then(() => {
            if (store.settings.searchOnLoad) {
                store.searchPosts();
            }
        })
        .catch(() => {
            store.setQueryParams();
        });
}

window.addEventListener("hashchange", onRouteChange);
onMounted(onPageLoad);
</script>

<template></template>
