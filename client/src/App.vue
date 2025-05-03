<script setup lang="ts">
import { ref } from "vue";
import Post from "./components/Post.vue";
import TagSearch from "./components/TagSearch.vue";
import TagChip from "./components/TagChip.vue";

const showHelp = ref(localStorage.getItem("hide-help") === null);
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

function onCloseHelp() {
    showHelp.value = false;
    localStorage.setItem("hide-help", "1");
}
</script>

<template>
    <div class="app">
        <header>
            <nav class="nav">
                <aside
                    class="search-help"
                    v-if="showHelp && posts.length === 0"
                >
                    <p>
                        <button class="close-btn" @click="onCloseHelp">
                            Close
                        </button>

                        Enter one tag at a time.
                    </p>
                    <ul>
                        <li>
                            <kbd>Tab</kbd>
                            will auto-complete.
                        </li>
                        <li>
                            <kbd>Up/Down</kbd>
                            selects a tag.
                        </li>
                        <li>
                            <kbd>Enter</kbd>
                            adds the tag to your search.
                        </li>
                    </ul>
                    <p>
                        <a
                            href="https://gelbooru.com/index.php?page=wiki&s=&s=view&id=26263"
                            target="_blank"
                        >
                            Gelbooru Search Help
                        </a>
                    </p>
                </aside>

                <TagSearch
                    @on-search="doSearch"
                    @on-submit="(t) => (tags = tags.concat(t))"
                    :exclude-tags="tags"
                />
                <button class="search-btn" type="submit" @click="doSearch">
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
    position: relative;
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

.search-btn {
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

.search-help {
    position: absolute;
    background-color: #2c3932;
    left: calc(100% + 10px);
    width: 300px;
    font-size: 14px;

    p,
    ul {
        margin: 12px;
    }

    ul {
        padding: 0;
        padding-left: 24px;
    }

    kbd {
        font-family: "Courier New", Courier, monospace;
    }

    a {
        color: #62b588;
    }
}

.close-btn {
    padding: 0;
    background: none;
    border: none;
    color: desaturate(#62b588, 30%);
    cursor: pointer;
    font-size: inherit;
    float: right;
    margin-left: 10px;

    &:hover {
        text-decoration: underline;
    }
}
</style>
