<script setup lang="ts">
import { ref } from "vue";
import Post from "./components/Post.vue";
import TagSearch from "./components/TagSearch.vue";
import TagChip from "./components/TagChip.vue";

const posts = ref<Post[]>([]);
const tags = ref<Tag[]>([]);

function doSearch() {
    fetch(
        "/api/posts?q=" +
            encodeURIComponent(tags.value.map((t) => t.name).join(" ")),
    )
        .then((resp) => {
            resp.json().then((json) => {
                posts.value = json.results;
                console.log(json);
            });
        })
        .catch((err) => console.error(err));
}
</script>

<template>
    <header>
        <nav class="nav">
            <TagSearch @submit="(t) => (tags = tags.concat(t))" :exclude-tags="tags" />
            <button class="searchButton" type="submit" @click="doSearch">search</button>

            <TagChip v-for="t in tags" :tag="t" />
        </nav>
    </header>
    <main>
        <div v-for="p in posts" :key="p.id">
            <Post :post="p" />
        </div>
    </main>
</template>

<style scoped lang="scss">
@use "sass:color";

.nav {
    max-width: 350px;
}

$btnColor: #342b3a;

.searchButton {
    display: block;
    width: 100%;
    margin-top: 8px;
    background-color: $btnColor;
    border-radius: 4px;
    border: 1px solid color.adjust($btnColor, $lightness: 20%);
    color: #fff;
    padding: 8px;
}
</style>
