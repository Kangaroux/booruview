<script setup lang="ts">
import { onMounted, ref } from "vue";
import ScreenCover from "./ScreenCover.vue";
import store from "@/store";

const timeRemaining = ref(3);

onMounted(() => {
    const timer = setInterval(() => {
        timeRemaining.value--;
        if (timeRemaining.value === 0) {
            clearInterval(timer);
        }
    }, 1000);
});

function consent() {
    store.settings.consented = true;
    store.settings.save();
}
</script>

<template>
    <ScreenCover blackout />
    <div class="cw-container">
        <div class="content-warning">
            <h3>CONTENT WARNING</h3>
            <p>
                This site is an image browser for Gelbooru, a
                <span class="nsfw-warning">NSFW image board.</span>
            </p>
            <p>
                Do NOT continue unless you consent to viewing potentially adult
                content.
            </p>
            <p>
                <button
                    class="btn-primary btn-rounded consent"
                    :disabled="timeRemaining > 0"
                    @click="consent"
                >
                    continue<span v-if="timeRemaining > 0">
                        ({{ timeRemaining }})</span
                    >
                </button>
            </p>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@import "@/assets/buttons";
@import "@/assets/colors";

.cw-container {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1;
}

.content-warning {
    padding: 10px;
}

h3 {
    text-align: center;
}

.nsfw-warning {
    font-weight: bold;
    color: #a00;
}

.consent {
    display: block;
    margin: 40px auto 0;
}
</style>
