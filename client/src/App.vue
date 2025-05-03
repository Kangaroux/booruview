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
            <TagSearch
                @on-search="doSearch"
                @on-submit="(t) => (tags = tags.concat(t))"
                :exclude-tags="tags"
            />
            <button class="searchButton" type="submit" @click="doSearch">
                search
            </button>

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
.nav {
    max-width: 350px;
}

.searchButton {
    $btn-color: #342b3a;
    $border-color: lighten($btn-color, 20%);

    display: block;
    width: 100%;
    margin-top: 8px;
    border-radius: 4px;
    color: #fff;
    padding: 8px;
    cursor: pointer;

    background-color: $btn-color;
    border: 1px solid $border-color;

    &:hover {
        background-color: lighten($btn-color, 2.5%);
        border-color: lighten($border-color, 2.5%);
    }
}
</style>
