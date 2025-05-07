<script setup lang="ts">
import { computed, ref, useTemplateRef, watch } from "vue";
import SearchSuggestions from "./SearchSuggestions.vue";

type SearchResponse = {
    results: Tag[];
};
const debounceMs = 150;

const emit = defineEmits<{ onSearch: []; onSubmit: [value: Tag] }>();

const { excludeTags = [] } = defineProps<{ excludeTags?: Tag[] }>();

const forceRenderKey = ref(0);
const query = ref("");
const selectedIndex = ref(-1);
const suggestions = ref<Tag[]>([]);
const timer = ref();
const inputRef = useTemplateRef("input");

// Convert excluded tags to a set for fast lookups
const excludeSet = computed(() => new Set(excludeTags.map((t) => t.name)));

function doSearch(query: string) {
    // Encoding the query prevents trailing whitespace from being stripped
    fetch("/api/tagsearch?q=" + encodeURIComponent(query))
        .then((resp) =>
            resp.json().then((json: SearchResponse) => {
                // Remove excluded tags from suggestions
                suggestions.value = json.results.filter(
                    (x) => !excludeSet.value.has(x.name),
                );
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

function autoComplete() {
    if (suggestions.value.length === 0) {
        return;
    }

    const index = selectedIndex.value === -1 ? 0 : selectedIndex.value;
    query.value = suggestions.value[index].name;
}

function onSuggestionClick(index: number) {
    selectedIndex.value = index;
    autoComplete();
    onSubmit();
}

function onInput(e: Event) {
    if (e.target === null) {
        return;
    }

    // Prevent the user from entering any leading whitespace
    const newVal = (e.target as HTMLInputElement).value.trimStart();
    const changed = query.value !== newVal;

    if (changed) {
        query.value = newVal;
    } else {
        // query.value didn't change but the DOM element still has the whitespace.
        // Incrementing the key will force Vue to re-render with the cleaned value
        forceRenderKey.value++;
    }
}

// Emits an event to the parent with the tag info
function onSubmit() {
    if (query.value.length === 0) {
        emit("onSearch");
        return;
    }

    let tag: Tag;

    if (selectedIndex.value !== -1) {
        tag = suggestions.value[selectedIndex.value];

        // Refocus the search input if the user was selecting a suggestion
        inputRef.value?.focus();
    } else {
        const match = suggestions.value.find((t) => t.name === query.value);

        // User is submitting a raw tag that wasn't in the search suggestions
        if (match) {
            tag = match;
        } else {
            tag = {
                count: 0,
                name: query.value,
                type: "unknown",
            };
        }
    }

    query.value = "";
    selectedIndex.value = -1;
    emit("onSubmit", tag);
}

// Setup a debounce to fetch search results shortly after the user stops typing
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
        @keydown.enter.prevent="onSubmit()"
        @keydown.tab.prevent="autoComplete()"
        @keydown.up.prevent="changeSelection(-1)"
        @keydown.down.prevent="changeSelection(1)"
    >
        <!-- forceRenderKey triggers a re-render when changed -->
        <template :key="forceRenderKey" />

        <input
            class="search"
            type="text"
            ref="input"
            placeholder="e.g: 1girl"
            :value="query"
            @input="onInput"
            @focus="selectedIndex = -1"
        />
        <SearchSuggestions
            class="suggestions"
            :tags="suggestions"
            :selected-index="selectedIndex"
            @on-click="onSuggestionClick"
        />
    </div>
</template>

<style scoped>
.search {
    background-color: #252525;
    border: 1px solid #555;
    color: #ddd;
    display: block;
    width: 100%;
    padding-left: 8px;
    height: 40px;
}

.search-container {
    position: relative;
}

.suggestions {
    position: absolute;
    width: 100%;
}
</style>
