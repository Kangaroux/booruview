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

function parseQuery(raw: string | null): Promise<void> {
    return new Promise((resolve, reject) => {
        if (raw === null || raw.length === 0) {
            resolve();
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

                resolve();
            })
            .catch(reject);
    });
}

function loadQueryParams(): Promise<void> {
    return new Promise((resolve, reject) => {
        const params = new URLSearchParams(
            window.location.hash.replace(/^#/, ""),
        );
        const page = params.get("page") || "1";
        const query = params.get("q");

        if (query === null) {
            reject();
            return;
        }

        store.currentPage = parsePage(page);
        parseQuery(query).then(resolve).catch(reject);
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
        .then(() => store.searchPosts())
        .catch(() => {
            // TODO: default search?
        });
}

window.addEventListener("hashchange", onRouteChange);
onMounted(onPageLoad);
</script>

<template></template>
