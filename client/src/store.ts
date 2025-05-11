import { reactive } from "vue";
import { isSetEqual } from "./set";

type Store = {
    currentPage: number;
    totalPostCount: number;
    resultsPerPage: number;

    fullscreenPost: Post | null;

    searchTagsSet: Set<Tag>;
    lastSearchTags: Set<Tag>;
    posts: Map<number, Post[]>;
    cachedTags: { [key: string]: Tag };

    addSearchTag(tag: Tag): void;
    removeSearchTag(tag: Tag): void;
    searchTags(): Tag[];
    tagsForPost(post: Post): Promise<Tag[]>;

    hasResults(): boolean;
    loadTags(tags: string[]): Promise<void>;
    maxPage(): number;
    nextPage(): void;
    postsForCurrentPage(): Post[] | undefined;
    prevPage(): void;
    searchPosts(): Promise<void>;
};

const store = reactive<Store>({
    currentPage: 1,
    totalPostCount: 0,
    resultsPerPage: 0,

    fullscreenPost: null,

    searchTagsSet: new Set(),
    lastSearchTags: new Set(),
    posts: new Map(),
    cachedTags: {},

    tagsForPost(post: Post): Promise<Tag[]> {
        return new Promise<Tag[]>((resolve, reject) => {
            store
                .loadTags(post.tags)
                .then(() => resolve(post.tags.map((t) => store.cachedTags[t])))
                .catch(reject);
        });
    },

    searchTags(): Tag[] {
        return Array.from(this.searchTagsSet.keys());
    },

    addSearchTag(tag: Tag) {
        this.searchTagsSet.add(tag);
    },

    removeSearchTag(tag: Tag) {
        this.searchTagsSet.delete(tag);
    },

    hasResults(): boolean {
        return this.totalPostCount > 0;
    },

    searchPosts(): Promise<void> {
        type PostListResponse = {
            count_per_page: number;
            total_count: number;
            results: Post[];
        };

        return new Promise((resolve, reject) => {
            const searchChanged = !isSetEqual(
                this.lastSearchTags,
                this.searchTagsSet,
            );
            const hasPage = this.posts.has(this.currentPage);

            if (!searchChanged && hasPage) {
                resolve();
                return;
            }

            const tags = this.searchTags().map((t) => t.name);
            const query = `q=${encodeURIComponent(tags.join(" "))}&page=${this.currentPage}`;

            fetch("/api/posts?" + query)
                .then((resp) => {
                    resp.json().then((json: PostListResponse) => {
                        if (searchChanged) {
                            this.posts.clear();
                        }

                        store.posts.set(this.currentPage, json.results);
                        store.resultsPerPage = json.count_per_page;
                        store.totalPostCount = json.total_count;
                        store.lastSearchTags = new Set(this.searchTagsSet);

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

        const maxTagsPerRequest = 100;

        // Fetch tags in parallel if there are too many for one request
        if (tags.length > maxTagsPerRequest) {
            let requests: Promise<void>[] = [];

            for (let i = 0; i < tags.length; i += maxTagsPerRequest) {
                const start = i;
                const end = i + maxTagsPerRequest;
                requests = requests.concat(
                    this.loadTags(tags.slice(start, end)),
                );
            }

            return new Promise((resolve, reject) =>
                Promise.all(requests)
                    .then(() => resolve())
                    .catch(() => reject()),
            );
        }

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
            this.searchPosts();
        }
    },

    prevPage() {
        if (this.currentPage > 0) {
            this.currentPage--;
            this.searchPosts();
        }
    },
});

export default store;
