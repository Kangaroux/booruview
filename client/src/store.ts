import { reactive } from "vue";
import SearchQuery from "./search";

type SettingsKey = keyof Omit<Store["settings"], "save" | "write">;

function loadValue<K extends SettingsKey, V = Store["settings"][K]>(
    key: K,
    defaultValue: V,
    transform: (val: string) => V,
): V {
    const val = localStorage?.getItem(key);
    if (val === null) {
        return defaultValue;
    }
    return transform(val);
}

type Store = {
    currentPage: number;
    totalPostCount: number;
    resultsPerPage: number;

    fullscreenPost: Post | null;

    sidebarClosed: boolean;

    settings: {
        columnSizing: "fixed" | "dynamic";
        columnCount: number;
        columnWidth: number;
        helpClosed: boolean;

        save(): void;
        write<K extends SettingsKey, V = Store["settings"][K]>(
            key: K,
            transform: (val: V) => string,
        ): void;
    };

    search: {
        query: SearchQuery;
        previousQuery: SearchQuery;
    };

    /** mapping of page number to posts */
    posts: Map<number, Post[]>;
    cachedTags: Map<string, Tag>;

    tagsForPost(post: Post): Promise<Tag[]>;

    hasResults(): boolean;
    loadTags(tags: string[]): Promise<void>;
    maxPage(): number;
    nextPage(): void;
    postsForCurrentPage(): Post[] | undefined;
    prevPage(): void;
    searchPosts({ closeSidebar }: { closeSidebar: boolean }): Promise<void>;
    setQueryParams(): void;
};

const store = reactive<Store>({
    currentPage: 1,
    totalPostCount: 0,
    resultsPerPage: 0,

    fullscreenPost: null,

    sidebarClosed: false,

    settings: {
        columnSizing: loadValue("columnSizing", "dynamic", (v) => v as any),
        columnCount: loadValue("columnCount", 3, parseInt),
        columnWidth: loadValue("columnWidth", 400, parseInt),
        helpClosed: loadValue("helpClosed", false, JSON.parse),

        save() {
            this.write("columnWidth", (v) => v.toString());
            this.write("helpClosed", JSON.stringify);
        },

        write<K extends keyof Store["settings"], V = Store["settings"][K]>(
            key: K,
            transform: (val: V) => string,
        ) {
            localStorage?.setItem(key, transform(this[key] as V));
        },
    },

    search: {
        query: new SearchQuery(),
        previousQuery: new SearchQuery(),
    },

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

    hasResults(): boolean {
        return this.totalPostCount > 0;
    },

    setQueryParams() {
        const url = new URL(window.location.href);
        url.searchParams.set("page", this.currentPage.toString());
        url.searchParams.set("q", this.search.query.asList().join(","));
        window.history.pushState(null, "", url.toString());
    },

    searchPosts({ closeSidebar }: { closeSidebar: boolean }): Promise<void> {
        type PostListResponse = {
            count_per_page: number;
            total_count: number;
            results: Post[];
        };

        return new Promise((resolve, reject) => {
            const searchHasChanged = !this.search.query.equals(
                this.search.previousQuery,
            );
            const query =
                `q=${encodeURIComponent(this.search.query.asList().join(" "))}` +
                `&page=${this.currentPage}`;

            fetch("/api/posts?" + query)
                .then((resp) => {
                    resp.json().then((json: PostListResponse) => {
                        if (searchHasChanged) {
                            this.posts.clear();
                        }

                        this.posts.set(this.currentPage, json.results);
                        this.resultsPerPage = json.count_per_page;
                        this.totalPostCount = json.total_count;
                        this.search.previousQuery = this.search.query.copy();

                        if (closeSidebar) {
                            this.sidebarClosed = true;
                        }

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
            this.searchPosts({ closeSidebar: false });
        }
    },

    prevPage() {
        if (this.currentPage > 0) {
            this.currentPage--;
            this.searchPosts({ closeSidebar: false });
        }
    },
});

export default store;
