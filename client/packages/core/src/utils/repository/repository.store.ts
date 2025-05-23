import { createContext, useContext } from 'react'

import { useMapStore } from '../map/map.store'
import { Entity, RepositoryStore } from './repository.types'

export function useRepositoryStore<TEntity extends Entity<string>>(): RepositoryStore<TEntity> {
  // State

  const itemsMap = useMapStore<string, TEntity>()

  // Actions

  // Adds one item
  const addOne = (item: TEntity, clear = false) => {
    itemsMap.addOne(item.id, item, clear)
  }

  // Adds many items
  const addMany = (items: TEntity[] | Record<string, TEntity>, clear = false) => {
    const newItems = Array.isArray(items)
      ? items.reduce(
          (result, item) => {
            result[item.id] = item
            return result
          },
          {} as Record<string, TEntity>,
        )
      : items

    itemsMap.addMany(newItems, clear)
  }

  return {
    items: itemsMap.values,
    itemById: itemsMap.valueByKey,
    itemsByIds: itemsMap.valuesByKeys,

    addOne,
    addMany,
    removeOne: itemsMap.removeOne,
    clear: itemsMap.clear,
  }
}

export const RepositoryStoreContext = createContext<RepositoryStore<any>>({
  isLoading: true,
} as any)

export function getRepositoryStore() {
  return useContext(RepositoryStoreContext)
}
