import { reactive } from "vue";

type Store = {
    posts: Post[];
    tags: { [key: string]: Tag };
    fetchingTags: boolean;
    loadTags(tags: string[]): Promise<void>;
};

const store = reactive<Store>({
    posts: [],
    tags: {},
    fetchingTags: false,
    loadTags(tags: string[]): Promise<void> {
        this.fetchingTags = true;

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
                        const newTags: {[key: string]: Tag} = {};
                        json.results.forEach((t) => (newTags[t.name] = t));
                        this.tags = {...this.tags, ...newTags};
                        this.fetchingTags = false;
                        resolve();
                    });
                })
                .catch((err) => {
                    console.error(err);
                    this.fetchingTags = false;
                    reject();
                });
        });
    },
});

export default store;
