<script setup lang="ts">
const emit = defineEmits(["on-close"]);
</script>

<template>
    <div class="help-container">
        <div class="help-content">
            <p>
                <button class="close-btn" @click="emit('on-close')">
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
</template>

<style lang="scss" scoped>
.help-container {
    background-color: #2c3932;
    width: 300px;
    font-size: 14px;
    padding: 12px;
    margin-top: 10px;

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
