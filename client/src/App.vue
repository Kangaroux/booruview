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
    <div class="app">
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
        <main class="post-container">
            <Post :post="p" v-for="p in posts" :key="p.id" />
        </main>
    </div>
</template>

<style scoped lang="scss">
.app {
    display: flex;
    flex-direction: row;
}

.nav {
    width: 350px;
}

.post-container {
    flex: 1;

    padding: 0 10px;

    column-width: 600px;
    column-gap: 10px;
}

.post {
    break-inside: avoid-column;
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
