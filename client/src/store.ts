import { reactive } from "vue";
import { isSetEqual } from "./set";

type Store = {
    currentPage: number;
    totalPostCount: number;
    resultsPerPage: number;

    lastSearch: Set<Tag>;
    posts: Map<number, Post[]>;
    cachedTags: { [key: string]: Tag };

    hasResults(): boolean;
    loadTags(tags: string[]): Promise<void>;
    maxPage(): number;
    nextPage(): void;
    postsForCurrentPage(): Post[] | undefined;
    prevPage(): void;
    searchPosts(tags: Tag[]): Promise<void>;
};

const store = reactive<Store>({
    currentPage: 1,
    totalPostCount: 0,
    resultsPerPage: 0,

    lastSearch: new Set(),
    posts: new Map(),
    cachedTags: {},

    hasResults(): boolean {
        return this.totalPostCount > 0;
    },

    searchPosts(tags: Tag[]): Promise<void> {
        type PostListResponse = {
            count_per_page: number;
            total_count: number;
            results: Post[];
        };

        return new Promise((resolve, reject) => {
            const searchTags = new Set(tags);
            const searchChanged = !isSetEqual(this.lastSearch, searchTags);
            const hasPage = this.posts.has(this.currentPage);

            if (!searchChanged && hasPage) {
                resolve();
                return;
            }

            const query = `q=${encodeURIComponent(tags.map((t) => t.name).join(" "))}&page=${this.currentPage}`;
            fetch("/api/posts?" + query)
                .then((resp) => {
                    resp.json().then((json: PostListResponse) => {
                        if (searchChanged) {
                            this.posts.clear();
                        }

                        store.posts.set(this.currentPage, json.results);
                        store.resultsPerPage = json.count_per_page;
                        store.totalPostCount = json.total_count;
                        store.lastSearch = searchTags;
                        resolve();
                    });
                })
                .catch((err) => {
                    console.error(err);
                    reject(err);
                });
        });
    },

    maxPage(): number {
        return Math.ceil(this.totalPostCount / this.resultsPerPage);
    },

    loadTags(tags: string[]): Promise<void> {
        type TagResponse = {
            results: Tag[];
        };

        return new Promise((resolve, reject) => {
            const missing = tags.filter((t) => !(t in this.cachedTags));

            if (missing.length === 0) {
                resolve();
                return;
            }

            fetch("/api/tags?q=" + encodeURIComponent(missing.join(" ")))
                .then((resp) => {
                    resp.json().then((json: TagResponse) => {
                        const newTags: { [key: string]: Tag } = {};
                        json.results.forEach((t) => (newTags[t.name] = t));
                        this.cachedTags = { ...this.cachedTags, ...newTags };
                        resolve();
                    });
                })
                .catch((err) => {
                    console.error(err);
                    reject();
                });
        });
    },

    postsForCurrentPage(): Post[] | undefined {
        return this.posts.get(this.currentPage);
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
