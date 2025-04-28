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
    rating: Rating;
    source_url: string;
    uploader: string;
    tags: string;
    thumbnail_url: string;
    thumbnail_width: number;
    thumbnail_height: number;
    lowres_url: string;
    low_res_width: number;
    low_res_height: number;
    image_url: string;
    width: number;
    height: number;
};
