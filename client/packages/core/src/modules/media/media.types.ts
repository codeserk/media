export enum MediaSource {
  Jellyfin = 'jellyfin',
}

export enum MediaType {
  Music = 'music',
}

export interface Media<TType extends MediaType, TContent extends object> {
  readonly id: string
  readonly source: MediaSource
  readonly type: TType
  readonly content: TContent
}

export interface MusicMediaContent {

}

