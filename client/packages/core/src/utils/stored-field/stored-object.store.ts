import { parseObject } from '../parser'
import { StorageRepository } from './storage-repository.types'
import { useStoredField } from './stored-field.store'

export function useStoredObject<T extends object>(
  key: string,
  defaultValue: T | undefined,
  storage?: StorageRepository,
) {
  return useStoredField<T | undefined>(
    key,
    defaultValue,
    (value) => (value ? JSON.stringify(value) : undefined),
    (value: unknown) => parseObject(value) as T | undefined,
    storage,
  )
}
