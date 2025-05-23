export enum MediaSource {
  Jellyfin = 'jellyfin',
}

export enum MediaType {
  Music = 'music',
}

export interface MusicSong {
  readonly id: string
  readonly artistId: string
  readonly albumId: string
  readonly source: MediaSource
  readonly name: string
  readonly releaseDate: string
  readonly index: number
}

export interface MusicArtist {
  readonly id: string
  readonly name: string
}

export interface MusicAlbum {
  readonly id: string
  readonly artistId: string
  readonly name: string
}
