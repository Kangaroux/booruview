/// <reference types="vite/client" />

type TagType =
    | "tag"
    | "artist"
    | "copyright"
    | "character"
    | "metadata"
    | "deprecated"
    | "unknown";

type Tag = {
    name: string;
    type: TagType;
    count: number;
};

type Rating = "general" | "sensitive" | "questionable" | "explicit";

type Post = {
    id: number;
    created_at: number;
    score: number;
    width: number;
    height: number;
    rating: Rating;
    source_url: string;
    uploader: string;
    tags: string;
    thumbnail_url: string;
    lowres_url: string;
    image_url: string;
};
