<script setup lang="ts">
import { COMMIT_SHA, LAST_COMMIT_DATE } from "@/config";
</script>

<template>
    <p>
        Booruview is an
        <a href="https://github.com/Kangaroux/booru-viewer" target="_blank"
            >open source</a
        >
        site for browsing
        <a href="https://gelbooru.com" target="_blank">Gelbooru</a> posts. Check
        the Help tab for more info.
    </p>
    <p>
        Feedback and suggestions are welcome. You can use the Github issue
        tracker or send me an
        <a href="mailto:2302541+Kangaroux@users.noreply.github.com">email</a>.
    </p>
    <p>
        This site does not use tracking or cookies. Searches are cached briefly
        and entirely anonymous.
    </p>
    <p class="version">last commit {{ COMMIT_SHA }} @ {{ LAST_COMMIT_DATE }}</p>
</template>

<style scoped>
.version {
    font-family: "Courier New", Courier, monospace;
    font-size: 12px;
    text-align: right;
}
</style>
