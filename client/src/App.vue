<script setup lang="ts">
import { defineModel, ref, watch } from "vue";
import TagChip from "./components/TagChip.vue";

type SearchResponse = {
  results: Tag[];
};

const debounceMs = 200;
const query = defineModel("query", {default: ""});
const results = ref<Tag[]>();
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
  <input type="text" v-model="query" />
  <ul>
    <li v-for="r in results"><TagChip :tag="r" /></li>
  </ul>
</template>

<style scoped>

</style>
