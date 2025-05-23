import { computed, useComputed, useSignal } from '@preact-signals/safe-react'
import { createContext, useCallback, useContext } from 'react'

import { MapStore } from './map.types'

export function useMapStore<TKey extends string | number, TValue>(): MapStore<TKey, TValue> {
  // State

  const map = useSignal<Record<TKey, TValue>>({} as Record<TKey, TValue>)

  // Getters

  const entries = useComputed(() => Object.entries(map.value) as [TKey, TValue][])
  const values = useComputed(() => Object.values(map.value) as TValue[])
  const valueByKey = useCallback(
    (key?: TKey) => computed(() => (key ? map.value[key] : undefined)),
    [],
  )
  const valuesByKeys = useCallback(
    (keys: TKey[]) => computed(() => keys.map((key) => map.value[key]).filter(Boolean)),
    [],
  )

  // Actions

  // Adds one value
  const addOne = (key: TKey, value: TValue, clear = false) => {
    const newValue = (clear ? {} : Object.assign({}, map.value)) as Record<TKey, TValue>
    newValue[key] = value

    map.value = newValue
  }

  // Adds many values
  const addMany = (values: Record<TKey, TValue>, clear = false) => {
    const newValue = (clear ? {} : Object.assign({}, map.value)) as Record<TKey, TValue>
    for (const key in values) {
      newValue[key] = values[key]
    }

    map.value = newValue
  }

  // Removes one value
  const removeOne = (key: TKey) => {
    delete map.value[key]
  }

  // Clears the map
  const clear = () => (map.value = {} as Record<TKey, TValue>)

  return {
    map,
    entries,
    values,
    valueByKey,
    valuesByKeys,

    addOne,
    addMany,
    removeOne,
    clear,
  }
}

export const MapStoreContext = createContext<MapStore<any, any>>({
  isLoading: true,
} as any)

export function getMapStore() {
  return useContext(MapStoreContext)
}
