<script setup lang="ts">
import { COMMIT_SHA, LAST_COMMIT_DATE } from "@/config";
</script>

<template>
    <p>
        Booruview makes it easy to search and browse
        <a href="https://gelbooru.com" target="_blank">Gelbooru</a> posts. Check
        the help tab for more info.
    </p>
    <p>
        This project is open source and development is ongoing. Feedback and
        suggestions are welcome. You can use the
        <a
            href="https://github.com/jessienyan/booruview/issues"
            target="_blank"
            >Github issue tracker</a
        >
        or send me an
        <a href="mailto:216619670+jessienyan@users.noreply.github.com">email</a>.
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
