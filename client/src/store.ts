import { reactive } from "vue";

type Store = {
    currentPage: number;
    pagesFetched: number;
    totalPostCount: number;
    resultsPerPage: number;

    posts: Post[];
    tags: { [key: string]: Tag };

    loadTags(tags: string[]): Promise<void>;
    nextPage(): void;
    postsForCurrentPage(): Post[];
    prevPage(): void;
    maxPage(): number;
};

const store = reactive<Store>({
    currentPage: 0,
    pagesFetched: 0,
    totalPostCount: 0,
    resultsPerPage: 0,

    posts: [],
    tags: {},

    maxPage(): number {
        const lastPage = Math.ceil(this.totalPostCount * this.resultsPerPage);

        // Pages are 0-indexed
        return lastPage - 1;
    },

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

    postsForCurrentPage(): Post[] {
        const start = (this.currentPage - 1) * this.resultsPerPage;
        const end = this.currentPage * this.resultsPerPage;
        return this.posts.slice(start, end);
    },

    nextPage() {
        if (this.currentPage < this.maxPage()) {
            this.currentPage++;
        }
    },

    prevPage() {
        if (this.currentPage > 0) {
            this.currentPage--;
        }
    },
});

export default store;
