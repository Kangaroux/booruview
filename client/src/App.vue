<script setup lang="ts">
import { ref } from "vue";
import store from "@/store";
import PostContainer from "./components/PostContainer.vue";
import TagSearch from "./components/search/TagSearch.vue";
import TagList from "./components/TagList.vue";
import Sidebar from "./components/Sidebar.vue";

const showHelp = ref(localStorage.getItem("hide-help") === null);

function onCloseHelp() {
    showHelp.value = false;
    localStorage.setItem("hide-help", "1");
}
</script>

<template>
    <div class="app">
        <header class="sidebar-container">
            <Sidebar />
        </header>
        <main>
            <div
                class="search-help"
                v-if="showHelp && store.posts.length === 0"
            >
                <div class="help-content">
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
                </div>
            </div>
            <PostContainer v-if="store.posts.length > 0" :posts="store.posts" />
        </main>
    </div>
</template>

<style scoped lang="scss">
@import "assets/colors";

.app {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;
    gap: 10px;
}

header {
    position: relative;
}

main {
    flex: 1;
    overflow-y: scroll;
}

.search-btn {
    $btn-color: #342b3a;
    $border-color: lighten($btn-color, 20%);
    $spinner-size: 20px;

    display: block;
    width: 100%;
    margin-top: 8px;
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

.search-help {
    background-color: #2c3932;
    width: 300px;
    font-size: 14px;
    padding: 12px;

    p,
    ul {
        margin: 0;

        &:not(:last-child) {
            margin-bottom: 1.2em;
        }
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
