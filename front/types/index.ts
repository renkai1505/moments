export type ResultVO<T> = {
    code: number
    message?: string
    data: T
}

export type LoginResp = {
    id: number
    token: string
    username: string
}

export type CommentVO = {
    id: number
    content: string
    username: string
    website?: string
    replyTo: string
    createdAt: string
    updatedAt: string
    memoId: number
    author: number
}


export type  MemoVO = {
    id: number
    content: string
    location: string
    imgs: string
    favCount: number
    userId: number
    createdAt: string
    updatedAt: string
    externalFavicon: string
    pinned: string
    ext: string
    externalTitle: string
    externalUrl: string
    showType: number
    user: UserVO,
    comments: Array<CommentVO>
    tags: string
}

export type UserVO = {
    id: number
    username: string
    nickname: string
    avatarUrl: string
    slogan: string
    coverUrl: string
    email: string
}
export type SysConfigVO = {
    version: string,
    commitId: string,
    adminUserName: string,
    title: string,
    favicon: string,
    beiAnNo: string,
    css: string,
    js: string,
    rss: string,
    enableAutoLoadNextPage: boolean
    enableS3: boolean
    enableRegister: boolean
    enableGoogleRecaptcha: boolean,
    googleSiteKey: string,
    enableComment: boolean,
    maxCommentLength: number,
    memoMaxHeight: number,
    commentOrder: 'desc' | 'asc',
    timeFormat: 'timeAgo' | 'time',
    s3:{
        thumbnailSuffix:string
    }
    enableEmail: boolean
    smtpHost: string
    smtpPort: string
    smtpUsername: string
    smtpPassword: string
}


export type MetingJSDTO = {
    id: string | undefined
    api: string | undefined
    server: "netease" | "tencent" | "kugou" | "xiami" | "baidu" | undefined,
    type: "song" | "playlist" | "album" | "search" | "artist" | undefined
}

export type MetingMusicServer = Exclude<MetingJSDTO['server'], undefined>
export type MetingMusicType = Exclude<MetingJSDTO['type'], undefined>


export type ExtDTO = {
    music: MusicDTO,
    doubanBook: DoubanBook,
    doubanMovie: DoubanMovie,
    video: Video,
}

export type MusicDTO = {
    id?: string,
    server?: MetingMusicServer,
    type?: MetingMusicType,
    api?: string
}

export type DoubanBook = {
    url?: string
    id?: string
    title?: string
    desc?: string
    image?: string
    isbn?: string
    author?: string
    rating?: string
    pubDate?: string
    keywords?: string
}

export type DoubanMovie = {
    url?: string
    id?: string
    title?: string
    desc?: string
    image?: string
    director?: string
    releaseDate?: string
    rating?: string
    actors?: string
    runtime?: string
}

export type Video = {
    type: 'youtube' | 'bilibili' | 'online'
    value: string
}

export type VideoType = Video["type"]

export type Friend = {
    id: number;
    name: string;
    icon: string;
    url: string;
    desc: string;
}

// 儿童相关类型
export type ChildVO = {
    id: number;
    name: string;
    nickname: string;
    birthDate: string;
    age: number;
    ageInDays: number;
    gender: 'M' | 'F';
    avatarUrl: string;
    coverUrl: string;
    height: number;
    weight: number;
    bloodType: string;
    hobbies: string;
    description: string;
    parentId: number;
    createdAt: string;
    updatedAt: string;
    parent?: UserVO;
}

export type SaveChildReq = {
    id?: number;
    name: string;
    nickname?: string;
    birthDate?: string;
    gender?: 'M' | 'F';
    avatarUrl?: string;
    coverUrl?: string;
    height?: number;
    weight?: number;
    bloodType?: string;
    hobbies?: string;
    description?: string;
}

export type GrowthRecordVO = {
    id: number;
    childId: number;
    title: string;
    content: string;
    recordType: 'growth' | 'health' | 'study' | 'play' | 'milestone';
    height?: number;
    weight?: number;
    imgs: string;
    location: string;
    recordDate: string;
    milestone: string;
    mood: string;
    weather: string;
    tags: string;
    parentId: number;
    pinned: boolean;
    showType: number;
    createdAt: string;
    updatedAt: string;
    child?: ChildVO;
    parent?: UserVO;
    imgConfigs?: any[];
}

export type SaveGrowthRecordReq = {
    id?: number;
    childId: number;
    title?: string;
    content: string;
    recordType: 'growth' | 'health' | 'study' | 'play' | 'milestone';
    height?: number;
    weight?: number;
    imgs?: string[];
    location?: string;
    recordDate?: string;
    milestone?: string;
    mood?: string;
    weather?: string;
    tags?: string[];
    pinned?: boolean;
    showType?: number;
}

export type TimelineItemVO = {
    date: string;
    records: GrowthRecordVO[];
}

export type GrowthStatsVO = {
    childId: number;
    totalRecords: number;
    recordTypes: { [key: string]: number };
    heightGrowth: GrowthDataPoint[];
    weightGrowth: GrowthDataPoint[];
    milestones: MilestoneVO[];
}

export type GrowthDataPoint = {
    date: string;
    value: number;
}

export type MilestoneVO = {
    id: number;
    title: string;
    description: string;
    date: string;
    ageAtTime: string;
}
