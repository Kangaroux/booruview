export function isSetEqual<T>(a: Set<T>, b: Set<T>): boolean {
    if (a.size !== b.size) {
        return false;
    }

    for (const k of a.keys()) {
        if (!b.has(k)) {
            return false;
        }
    }

    return true;
}
