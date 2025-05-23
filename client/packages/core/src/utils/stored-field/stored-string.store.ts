import { parseString } from '../parser'
import { StorageRepository } from './storage-repository.types'
import { useStoredField } from './stored-field.store'

export function useStoredString(
  key: string,
  defaultValue: string | undefined,
  storage?: StorageRepository,
) {
  return useStoredField<string | undefined>(
    key,
    defaultValue,
    (value) => value ?? '',
    (value) => {
      const strValue = parseString(value)
      if (strValue === 'undefined') {
        return
      }

      return strValue
    },
    storage,
  )
}
