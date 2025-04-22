<script setup lang="ts">
import { defineModel, ref, watch } from "vue";
import TagList from "./components/TagList.vue";

type SearchResponse = {
  results: Tag[];
};

const debounceMs = 200;
const query = defineModel("query", {default: ""});
const results = ref<Tag[]>([]);
const timer = ref();

const doSearch = (query: string) => {
  fetch("/api/search?q=" + query)
  .then(resp => resp.json().then((json: SearchResponse) => {
    results.value = json.results;
}))
  .catch(err => console.error(err));
}

watch(query, (query, _, onCleanup) => {
  onCleanup(() => clearTimeout(timer.value));
  if(query.length) {
    timer.value = setTimeout(() => doSearch(query), debounceMs);
  } else {
    results.value = [];
  }
});
</script>

<template>
  <div class="search-container">
    <input class="search" type="text" v-model="query" />
    <TagList :tags="results" />
  </div>
</template>

<style scoped>
.search {
  background-color: #252525;
  border: 1px solid #555;
  color: #DDD;
  display: block;
  width: 100%;
  padding: 8px;
}
.search-container {
  width: 300px;
}
</style>
