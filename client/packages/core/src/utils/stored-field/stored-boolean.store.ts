import { parseBoolean } from '../parser'
import { StorageRepository } from './storage-repository.types'
import { useStoredField } from './stored-field.store'

export function useStoredBoolean(key: string, defaultValue: boolean, storage?: StorageRepository) {
  return useStoredField<boolean>(
    key,
    defaultValue,
    (value) => value?.toString(),
    (value) => parseBoolean(value) ?? false,
    storage,
  )
}
