import { createContext, useContext } from 'react'

import { useLoadingStore } from '../../utils/loading.store'
import { useStoredRepository } from '../../utils/stored-field/stored-repository.store'

const MEDIA_KEY = 'media'

export function useMediaStore() {
  const { isLoading, isLoadingKey, operation } = useLoadingStore()
  const repository = useStoredRepository(MEDIA_KEY, {})

  return {
    isLoading,
    isLoadingKey,
    repository,
    mediaItems: repository.items,
    mediaItemById: repository.itemById,
    mediaItemsByIds: repository.itemsByIds,
  }
}

export type MediaStore = ReturnType<typeof useMediaStore>

export const MediaStoreContext = createContext<MediaStore>({
  isLoading: true,
} as any)

export function useMediaStoreContext() {
  return useContext(MediaStoreContext)
}
