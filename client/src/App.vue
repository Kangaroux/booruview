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

            <template v-for="t in tags">
                <TagChip :tag="t" />
            </template>

            <button type="submit" @click="doSearch">search</button>
        </nav>
    </header>
    <main>
        <div v-for="p in posts" :key="p.id">
            <Post :post="p" />
        </div>
    </main>
</template>
