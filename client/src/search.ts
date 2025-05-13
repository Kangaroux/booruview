class SearchQuery {
    include: Map<string, Tag>;
    exclude: Map<string, Tag>;

    constructor() {
        this.include = new Map();
        this.exclude = new Map();
    }

    includeTag(t: Tag) {
        this.include.set(t.name, t);
        this.exclude.delete(t.name);
    }

    excludeTag(t: Tag) {
        this.exclude.set(t.name, t);
        this.include.delete(t.name);
    }

    removeTag(t: Tag) {
        this.include.delete(t.name);
        this.exclude.delete(t.name);
    }

    asList(): string[] {
        const include = Array.from(this.include.values(), (t) => t.name);
        const exclude = Array.from(this.exclude.values(), (t) => "-" + t.name);
        return include.concat(exclude);
    }

    equals(o: SearchQuery): boolean {
        if (
            this.include.size !== o.include.size ||
            this.exclude.size !== o.exclude.size
        ) {
            return false;
        }

        for (const k of this.include.keys()) {
            if (!o.include.has(k)) {
                return false;
            }
        }

        for (const k of this.exclude.keys()) {
            if (!o.exclude.has(k)) {
                return false;
            }
        }

        return true;
    }

    copy(): SearchQuery {
        const clone = new SearchQuery();
        clone.include = new Map(this.include);
        clone.exclude = new Map(this.exclude);
        return clone;
    }
}

export default SearchQuery;
