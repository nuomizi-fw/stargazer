export interface Bangumi {
    id: number;
    title: string;
    englishTitle: string;
    japaneseTitle: string;
    status: string;
    score: number;
    genres: string[];
    synopsis: string;
    coverImage: string;
    trailerURL: string;
    seasons: Season[];
}

export interface Season {
    id: number;
    bangumiId: number;
    seasonNumber: number;
    startDate: string;
    endDate: string;
    episodes: number;
    cast: CastMember[];
}

export interface CastMember {
    id: number;
    seasonId: number;
    name: string;
    role: string;
    imageURL: string;
}
