import { useMemo } from 'react'

import { Entity } from '../repository/repository.types'
import { StorageRepository } from './storage-repository.types'
import { useStoredMap } from './stored-map.store'

export function useStoredRepository<TEntity extends Entity<string>>(
  key: string,
  defaultValue: Record<string, TEntity> | undefined,
  storage?: StorageRepository,
) {
  const field = useStoredMap<string, TEntity>(key, defaultValue, storage)

  // Actions

  // Adds one item
  const addOne = async (item: TEntity, clear = false) => {
    await field.addOne(item.id, item, clear)
  }

  // Adds many items
  const addMany = async (items: TEntity[] | Record<string, TEntity>, clear = false) => {
    const newItems = Array.isArray(items)
      ? items.reduce(
          (result, item) => {
            result[item.id] = item
            return result
          },
          {} as Record<string, TEntity>,
        )
      : items

    await field.addMany(newItems, clear)
  }

  return useMemo(
    () => ({
      isLoading: field.isLoading,
      load: field.load,
      loaded: field.loaded,
      items: field.values,
      itemById: field.valueByKey,
      itemsByIds: field.valuesByKeys,

      addOne,
      addMany,
      removeOne: field.removeOne,
      clear: field.clear,
    }),
    [],
  )
}
