import { reactive } from "vue";
import { isSetEqual } from "./set";

type Store = {
    currentPage: number;
    totalPostCount: number;
    resultsPerPage: number;

    fullscreenPost: Post | null;

    searchQuery: Set<Tag>;
    lastSearchQuery: Set<Tag>;

    /** mapping of page number to posts */
    posts: Map<number, Post[]>;
    cachedTags: Map<string, Tag>;

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
    setQueryParams(): void;
};

const store = reactive<Store>({
    currentPage: 1,
    totalPostCount: 0,
    resultsPerPage: 0,

    fullscreenPost: null,

    searchQuery: new Set(),
    lastSearchQuery: new Set(),
    posts: new Map(),
    cachedTags: new Map(),

    tagsForPost(post: Post): Promise<Tag[]> {
        return new Promise<Tag[]>((resolve, reject) => {
            store
                .loadTags(post.tags)
                .then(() =>
                    resolve(
                        post.tags
                            .map((t) => store.cachedTags.get(t))
                            .filter((t) => t != null),
                    ),
                )
                .catch(reject);
        });
    },

    searchTags(): Tag[] {
        return Array.from(this.searchQuery.keys());
    },

    addSearchTag(tag: Tag) {
        this.searchQuery.add(tag);
    },

    removeSearchTag(tag: Tag) {
        this.searchQuery.delete(tag);
    },

    hasResults(): boolean {
        return this.totalPostCount > 0;
    },

    setQueryParams() {
        const url = new URL(window.location.href);
        url.searchParams.set("page", this.currentPage.toString());
        url.searchParams.set("tags", this.searchTags().map(t => t.name).join(","));
        window.history.pushState(null, "", url.toString());
    },

    searchPosts(): Promise<void> {
        type PostListResponse = {
            count_per_page: number;
            total_count: number;
            results: Post[];
        };

        return new Promise((resolve, reject) => {
            const searchChanged = !isSetEqual(
                this.lastSearchQuery,
                this.searchQuery,
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

                        this.posts.set(this.currentPage, json.results);
                        this.resultsPerPage = json.count_per_page;
                        this.totalPostCount = json.total_count;
                        this.lastSearchQuery = new Set(this.searchQuery);

                        this.setQueryParams();
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
            const missing = tags.filter((t) => !this.cachedTags.has(t));

            if (missing.length === 0) {
                resolve();
                return;
            }

            fetch("/api/tags?q=" + encodeURIComponent(missing.join(" ")))
                .then((resp) => {
                    resp.json().then((json: TagResponse) => {
                        json.results.forEach((t) =>
                            this.cachedTags.set(t.name, t),
                        );
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
