import { computed, ReadonlySignal, useComputed } from '@preact-signals/safe-react'
import { useMemo } from 'react'

import { StorageRepository } from './storage-repository.types'
import { useStoredObject } from './stored-object.store'

export function useStoredMap<TKey extends string | number, TValue>(
  key: string,
  defaultValue: Record<TKey, TValue> | undefined,
  storage?: StorageRepository,
) {
  const field = useStoredObject<Record<TKey, TValue>>(key, defaultValue, storage)

  // Getters

  const entries = useComputed(() => Object.entries(field.current.value ?? {}) as [TKey, TValue][])
  const values = useComputed(() => Object.values(field.current.value ?? {}) as TValue[])
  const valueByKey = (key?: TKey) => computed(() => (key ? field.current.value?.[key] : undefined))
  const valuesByKeys = (keys: TKey[]) =>
    computed(() => keys.map((key) => field.current.value?.[key]).filter(Boolean) as TValue[])

  // Actions

  // Adds one value
  const addOne = async (key: TKey, value: TValue, clear = false) => {
    const newValue = (clear ? {} : Object.assign({}, field.current.value)) as Record<TKey, TValue>
    newValue[key] = value

    await field.save(newValue)
  }

  // Adds many values
  const addMany = async (values: Record<TKey, TValue>, clear = false) => {
    const newValue = (clear ? {} : Object.assign({}, field.current.value)) as Record<TKey, TValue>
    for (const key in values) {
      newValue[key] = values[key]
    }

    await field.save(newValue)
  }

  // Removes one value
  const removeOne = async (key: TKey) => {
    const newValue = Object.assign({}, field.current.value)
    delete newValue[key]

    await field.save(newValue)
  }

  // Clears the map
  const clear = async () => await field.save({} as Record<TKey, TValue>)

  return useMemo(
    () => ({
      isLoading: field.isLoading,
      load: field.load,
      loaded: field.loaded,
      map: field.current as ReadonlySignal<Record<TKey, TValue>>,
      entries,
      values,
      valueByKey,
      valuesByKeys,

      addOne,
      addMany,
      removeOne,
      clear,
    }),
    [],
  )
}
