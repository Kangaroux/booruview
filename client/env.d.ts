/// <reference types="vite/client" />

type TagType = 	"tag" | "artist" | "copyright" | "character" | "metadata" | "deprecated" | "unknown";
type Tag = {
    name: string;
    type: TagType;
    count: number;
};
