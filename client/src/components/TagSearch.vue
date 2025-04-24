<script setup lang="ts">
import { defineModel, ref, watch } from "vue";
import SearchSuggestions from "./SearchSuggestions.vue";

type SearchResponse = {
    results: Tag[];
};

const debounceMs = 200;
const query = defineModel("query", { default: "" });
const results = ref<Tag[]>([]);
const timer = ref();
const selectedIndex = ref(-1);

function doSearch(query: string) {
    // Encoding the query prevents trailing whitespace from being stripped
    fetch("/api/search?q=" + encodeURIComponent(query))
        .then((resp) =>
            resp.json().then((json: SearchResponse) => {
                results.value = json.results;
            }),
        )
        .catch((err) => console.error(err));
}

function changeSelection(direction: number) {
    direction += selectedIndex.value;

    if (direction < 0) {
        direction = results.value.length - 1;
    } else if (direction >= results.value.length) {
        direction = 0;
    }

    selectedIndex.value = direction;
}

watch(query, (query, _, onCleanup) => {
    onCleanup(() => clearTimeout(timer.value));

    selectedIndex.value = -1;

    if (query.length) {
        timer.value = setTimeout(() => doSearch(query), debounceMs);
    } else {
        results.value = [];
    }
});
</script>

<template>
    <div
        class="search-container"
        @keydown.up.prevent="changeSelection(-1)"
        @keydown.down.prevent="changeSelection(1)"
    >
        <input
            class="search"
            type="text"
            v-model="query"
            @focus="selectedIndex = -1"
        />
        <SearchSuggestions :tags="results" :selected-index="selectedIndex" />
    </div>
</template>

<style scoped>
.search {
    background-color: #252525;
    border: 1px solid #555;
    color: #ddd;
    display: block;
    width: 100%;
    padding: 8px;
}
.search-container {
    width: 300px;
}
</style>
