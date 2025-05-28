<script setup lang="ts">
import { ref, useTemplateRef, watch } from "vue";
import SearchSuggestions from "./Suggestions.vue";
import store from "@/store";

type SearchResponse = {
    results: Tag[];
};
const debounceMs = 100;

const emit = defineEmits<{
    onSearch: [];
    onTagSelect: [value: Tag, negated: boolean];
}>();

const { showSpinner = false } = defineProps<{
    showSpinner?: boolean;
}>();

const forceRenderKey = ref(0);
const inputVal = ref("");
const selectedIndex = ref(-1);
const suggestions = ref<Tag[]>([]);
const timer = ref();
const inputRef = useTemplateRef("input");
const showSuggestions = ref(false);

function doTagSearch(query: string) {
    // Encoding the query prevents trailing whitespace from being stripped
    fetch("/api/tagsearch?q=" + encodeURIComponent(query))
        .then((resp) => {
            // Request took too long and results don't match the input, discard
            if (query !== inputVal.value) {
                return;
            }

            resp.json().then((json: SearchResponse) => {
                suggestions.value = json.results.filter(
                    // Don't suggest tags already added to the search
                    (t) =>
                        !store.query.include.has(t.name) &&
                        !store.query.exclude.has(t.name),
                );
            });
        })
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
    let val = suggestions.value[index].name;

    // Preserve the NOT op
    if (inputVal.value.startsWith("-")) {
        val = "-" + val;
    }

    inputVal.value = val;
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
    const changed = inputVal.value !== newVal;

    if (changed) {
        inputVal.value = newVal;
        showSuggestions.value = true;
    } else {
        // query.value didn't change but the DOM element still has the whitespace.
        // Incrementing the key will force Vue to re-render with the cleaned value
        forceRenderKey.value++;
    }
}

// Emits an event with a selected tag or to trigger a post search
function onSubmit() {
    if (inputVal.value.length === 0) {
        emit("onSearch");
        return;
    }

    let tag: Tag;
    const negated = inputVal.value.startsWith("-");

    if (selectedIndex.value !== -1) {
        tag = suggestions.value[selectedIndex.value];

        // Refocus the search input if the user was selecting a suggestion
        inputRef.value?.focus();
    } else {
        const value = negated ? inputVal.value.slice(1) : inputVal.value;
        const match = suggestions.value.find((t) => t.name === value);

        // User is submitting a raw tag that wasn't in the search suggestions
        if (match) {
            tag = match;
        } else {
            tag = {
                count: 0,
                name: value,
                type: "unknown",
            };
        }
    }

    inputVal.value = "";
    selectedIndex.value = -1;
    emit("onTagSelect", tag, negated);
}

// Setup a debounce to fetch search results shortly after the user stops typing
watch(inputVal, (query, _, onCleanup) => {
    onCleanup(() => clearTimeout(timer.value));

    selectedIndex.value = -1;

    if (query.length) {
        timer.value = setTimeout(() => doTagSearch(query), debounceMs);
    } else {
        suggestions.value = [];
    }
});
</script>

<template>
    <div
        class="container"
        @keydown.enter.prevent="onSubmit()"
        @keydown.tab.prevent="autoComplete()"
        @keydown.up.prevent="changeSelection(-1)"
        @keydown.down.prevent="changeSelection(1)"
        @keydown.esc.prevent="showSuggestions = false"
    >
        <!-- forceRenderKey triggers a re-render when changed -->
        <template :key="forceRenderKey" />

        <input
            class="input"
            type="text"
            ref="input"
            placeholder="e.g: 1girl"
            :value="inputVal"
            @input="onInput"
            @focus="selectedIndex = -1"
            autofocus
        />
        <SearchSuggestions
            class="suggestions"
            :tags="suggestions"
            :selected-index="selectedIndex"
            @click="onSuggestionClick"
            v-if="showSuggestions"
        />
    </div>
    <button
        class="submit-btn"
        type="submit"
        @click="onSubmit"
        :disabled="showSpinner"
    >
        <span v-if="!showSpinner">search</span>
        <span v-else class="spinner"></span>
    </button>
</template>

<style lang="scss" scoped>
.input {
    background-color: #252525;
    border: 1px solid #555;
    color: #ddd;
    display: block;
    width: 100%;
    padding-left: 8px;
    height: 40px;
}

.container {
    position: relative;
}

.suggestions {
    position: absolute;
    width: 100%;
    z-index: 1;
    padding-bottom: 10px;
}

.submit-btn {
    $btn-color: #342b3a;
    $border-color: lighten($btn-color, 20%);
    $spinner-size: 20px;

    display: block;
    width: 100%;
    margin: 10px 0;
    border-radius: 4px;
    color: #bb9fce;
    padding: 8px;
    cursor: pointer;
    font-size: 16px;
    line-height: $spinner-size;

    background-color: $btn-color;
    border: 1px solid $border-color;

    &:hover {
        background-color: lighten($btn-color, 2.5%);
        border-color: lighten($border-color, 2.5%);
    }

    &:disabled {
        cursor: default;
        background-color: darken($btn-color, 2.5%);
        border-color: darken($border-color, 2.5%);
    }

    span {
        display: block;
        margin: 0 auto;
    }

    .spinner {
        width: $spinner-size;
        height: $spinner-size;
        border-radius: 50%;
        border: 2px solid #fff;
        border-bottom-color: transparent;
        animation: linear 1s spin-anim infinite;
        display: block;

        @keyframes spin-anim {
            from {
                rotate: 0;
            }
            to {
                rotate: 360deg;
            }
        }
    }
}
</style>
