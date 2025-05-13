<script setup lang="ts">
const emit = defineEmits(["on-close"]);
</script>

<template>
    <div class="help-container">
        <div class="help-content">
            <button class="close-btn" @click="emit('on-close')">
                Close
            </button>

            <h3>Quickstart</h3>
            <p>
                Booruview lets you search and browse <a href="https://gelbooru.com" target="_blank">Gelbooru</a> posts.
            </p>
            <p>
                Add tags to your search one at a time. To exclude a tag, use <code>-tag</code>. Results must match every tag in your search (AND).
            </p>
            <p>
                SFW Example:
            </p>
            <p>Type <code>rating:general</code> and confirm the tag. Type <code>blue_sky</code> and confirm. Click search.</p>
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
    font-size: 14px;
    padding: 12px;
    margin-bottom: 10px;

    h3 {
        margin-top: 0;
    }

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

    code {
        font-family: "Courier New", Courier, monospace;
        background-color: rgba(0,0,0,0.2);
        padding: 1px 3px;
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
