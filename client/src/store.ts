import { reactive } from "vue";

type Store = {
    posts: Post[];
    tags: { [key: string]: Tag };
    loadTags(tags: string[]): Promise<void>;
};

const store = reactive<Store>({
    posts: [],
    tags: {},
    loadTags(tags: string[]): Promise<void> {
        return new Promise((resolve, reject) => {
            const missing = tags.filter((t) => !(t in this.tags));

            if (missing.length === 0) {
                resolve();
                return;
            }

            type TagResponse = {
                results: Tag[];
            };

            fetch("/api/tags?q=" + encodeURIComponent(missing.join(" ")))
                .then((resp) => {
                    resp.json().then((json: TagResponse) => {
                        const newTags: { [key: string]: Tag } = {};
                        json.results.forEach((t) => (newTags[t.name] = t));
                        this.tags = { ...this.tags, ...newTags };
                        resolve();
                    });
                })
                .catch((err) => {
                    console.error(err);
                    reject();
                });
        });
    },
});

export default store;
