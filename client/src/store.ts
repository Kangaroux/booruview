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
                        json.results.forEach((t) => (this.tags[t.name] = t));
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
