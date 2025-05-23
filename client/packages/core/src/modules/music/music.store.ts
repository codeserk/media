import { createContext, useContext } from 'react'

import { useLoadingStore } from '../../utils/loading.store'
import { useRepositoryStore } from '../../utils/repository/repository.store'

export function useMusicStore() {
  const { isLoading } = useLoadingStore()
  const artists = useRepositoryStore()
  const albums = useRepositoryStore()
  const songs = useRepositoryStore()

  return {
    isLoading,
    artists,
    albums,
    songs,
  }
}

export type MusicStore = ReturnType<typeof useMusicStore>

export const MusicStoreContext = createContext<MusicStore>({
  isLoading: true,
} as any)

export function useMusicStoreContext() {
  return useContext(MusicStoreContext)
}
