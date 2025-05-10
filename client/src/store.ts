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
    searchPosts(tags: Tag[]): Promise<void>;
};

const store = reactive<Store>({
    currentPage: 1,
    pagesFetched: 0,
    totalPostCount: 0,
    resultsPerPage: 0,

    posts: [],
    tags: {},

    searchPosts(tags: Tag[]): Promise<void> {
        type PostListResponse = {
            count_per_page: number;
            total_count: number;
            results: Post[];
        };

        return new Promise((resolve, reject) => {
            const query = `q=${encodeURIComponent(tags.map((t) => t.name).join(" "))}&page=${this.currentPage}`;
            fetch("/api/posts?" + query)
                .then((resp) => {
                    resp.json().then((json: PostListResponse) => {
                        store.posts = json.results;
                        store.resultsPerPage = json.count_per_page;
                        store.totalPostCount = json.total_count;
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
            const missing = tags.filter((t) => !(t in this.tags));

            if (missing.length === 0) {
                resolve();
                return;
            }

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
