import { batch, useSignal } from '@preact-signals/safe-react'
import { useEffect, useMemo } from 'react'

import { useLoadingStore } from '../loading.store'
import { StorageRepository } from './storage-repository.types'

export function useStoredField<TValue>(
  key: string,
  defaultValue: TValue,
  serialize: (value: TValue) => string | undefined,
  deserialize: (raw: unknown) => TValue,
  storage?: StorageRepository,
) {
  const { isLoading, operation } = useLoadingStore()

  const loaded = useSignal(false)
  const current = useSignal<TValue>(defaultValue)

  const load = operation(async () => {
    batch(async () => {
      if (storage) {
        const newValue = deserialize(await storage.get(key))
        if (newValue) {
          current.value = newValue ?? current.value ?? defaultValue
        }
      }

      loaded.value = true
    })
  })

  const save = operation(async (newValue: TValue) => {
    current.value = newValue
    await storage?.set(key, serialize(newValue))
  })

  const remove = operation(async () => {
    current.value = defaultValue

    await storage?.set(key, null)
  })

  useEffect(() => {
    load()
  }, [])

  return useMemo(
    () => ({
      isLoading,
      current,
      loaded,

      load,
      save,
      remove,
    }),
    [],
  )
}
