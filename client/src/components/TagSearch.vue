<script setup lang="ts">
import { defineEmits, ref, watch } from "vue";
import SearchSuggestions from "./SearchSuggestions.vue";

type SearchResponse = {
    results: Tag[];
};

const emit = defineEmits<{submit: [value: Tag]}>();

const debounceMs = 200;
const query = ref("");
const selectedIndex = ref(-1);
const suggestions = ref<Tag[]>([]);
const timer = ref();
const forceRenderKey = ref(0);

function doSearch(query: string) {
    // Encoding the query prevents trailing whitespace from being stripped
    fetch("/api/search?q=" + encodeURIComponent(query))
        .then((resp) =>
            resp.json().then((json: SearchResponse) => {
                suggestions.value = json.results;
            }),
        )
        .catch((err) => console.error(err));
}

function changeSelection(direction: number) {
    direction += selectedIndex.value;

    if (direction < 0) {
        direction = suggestions.value.length - 1;
    } else if (direction >= suggestions.value.length) {
        direction = 0;
    }

    selectedIndex.value = direction;
}

function chooseSuggestion() {
    if(suggestions.value.length === 0) {
        return;
    }

    const index = selectedIndex.value === -1 ? 0 : selectedIndex.value;
    query.value = suggestions.value[index].name;
}

function onInput(e: Event) {
    if(e.target === null) {
        return;
    }

    // Prevent the user from entering any leading whitespace
    const newVal = (e.target as HTMLInputElement).value.trimStart();
    const changed = query.value !== newVal;

    if(changed) {
        query.value = newVal;
    } else {
        // query.value didn't change but the DOM element still has the whitespace.
        // Incrementing the key will force Vue to re-render with the cleaned value
        forceRenderKey.value++;
    }
}

function submit() {
    if(query.value.length === 0) {
        return;
    }

    let tag: Tag;

    if(selectedIndex.value !== -1) {
        tag = suggestions.value[selectedIndex.value];
    } else {
        const match = suggestions.value.find(t => t.name === query.value);

        // User is submitting a raw tag that wasn't in the search suggestions
        if(match) {
            tag = match;
        } else {
            tag = {
                count: 0,
                name: query.value,
                type: "unknown",
            }
        }
    }

    query.value = "";
    selectedIndex.value = -1;
    emit("submit", tag);
}

watch(query, (query, _, onCleanup) => {
    onCleanup(() => clearTimeout(timer.value));

    selectedIndex.value = -1;

    if (query.length) {
        timer.value = setTimeout(() => doSearch(query), debounceMs);
    } else {
        suggestions.value = [];
    }
});
</script>

<template>
    <div
        class="search-container"
        @keydown.enter.prevent="submit()"
        @keydown.tab.prevent="chooseSuggestion()"
        @keydown.up.prevent="changeSelection(-1)"
        @keydown.down.prevent="changeSelection(1)"
    >
        <!-- forceRenderKey triggers a re-render when changed -->
        <template :key="forceRenderKey" />

        <input
            class="search"
            type="text"
            :value="query"
            @input="onInput"
            @focus="selectedIndex = -1"
        />
        <SearchSuggestions :tags="suggestions" :selected-index="selectedIndex" />
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
