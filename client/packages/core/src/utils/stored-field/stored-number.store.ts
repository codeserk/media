import { parseNumber } from '../parser'
import { StorageRepository } from './storage-repository.types'
import { useStoredField } from './stored-field.store'

export function useStoredNumber(
  key: string,
  defaultValue: number | undefined,
  storage?: StorageRepository,
) {
  return useStoredField<number | undefined>(
    key,
    defaultValue,
    (value) => value?.toString(),
    parseNumber,
    storage,
  )
}
