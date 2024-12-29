export type ObjectPath<TObject, TKey extends keyof TObject = keyof TObject> = TKey extends
  | string
  | number
  ? TObject[TKey] extends infer R
    ? `${TKey}` | (R extends Record<string, unknown> ? `${TKey}.${ObjectPath<R>}` : never)
    : never
  : never

export type ObjectValue<
  TObject,
  TPath extends ObjectPath<TObject>,
> = TPath extends `${infer K}.${infer Rest}`
  ? TObject[(K extends `${infer R extends number}` ? R : K) & keyof TObject] extends infer S
    ? S extends never
      ? never
      : Rest extends ObjectPath<S>
        ? ObjectValue<S, Rest>
        : never
    : never
  : TObject[(TPath extends `${infer R extends number}` ? R : TPath) & keyof TObject]

export function deepCopy<T = any>(object: T): T {
  return JSON.parse(JSON.stringify(object))
}

export function getObjectValueFromPath<
  TObject extends object,
  TObjectPath extends ObjectPath<TObject> = ObjectPath<TObject>,
>(object: TObject, path: TObjectPath): ObjectValue<TObject, TObjectPath> | undefined {
  const pathParts = path.split('.')
  let index = 0
  let current: any = object

  do {
    current = current?.[pathParts[index]]
  } while (++index < pathParts.length)

  return current
}

export function combineObjectPath<
  TObject extends object,
  TObjectPath extends ObjectPath<TObject> = ObjectPath<TObject>,
>(path: TObjectPath | undefined, key: string): ObjectPath<TObject> {
  if (!path) {
    return key as ObjectPath<TObject>
  }

  return `${path}.${key}` as ObjectPath<TObject>
}

/**
 * Returns a map of the objects by one of the field
 * @param objects
 * @param key
 * @returns map
 */
export function objectsByKey<T extends Record<string, any>>(
  objects: T[] | undefined,
  key: keyof T | ((object: T) => unknown) = 'id',
): Record<string, T> {
  if (!objects?.length) {
    return {}
  }

  return objects.reduce(
    (result, object) => {
      const objectKey = typeof key === 'function' ? key(object) : object[key]

      if (typeof objectKey === 'string') {
        result[objectKey as string] = object
      }

      return result
    },
    {} as Record<string, T>,
  )
}

export function flattenStringObject(
  obj: Record<string, unknown>,
  parentKey?: string,
): Record<string, string> {
  let result: Record<string, string> = {}

  for (const key in obj) {
    const composedKey = parentKey ? `${parentKey}.${key}` : key
    const value = obj[key]
    if (typeof value === 'object') {
      result = { ...result, ...flattenStringObject(value as Record<string, string>, composedKey) }
    } else if (typeof value === 'string') {
      result[composedKey] = value
    }
  }

  return result
}
