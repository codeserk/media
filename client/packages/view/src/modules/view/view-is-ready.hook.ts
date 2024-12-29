import { ReadonlySignal, useComputed, useSignal } from '@preact-signals/safe-react'
import { useEffect } from 'react'

import { ViewIsReadyProps } from './view.types'

/**
 * Hook to determine whether the view is ready to be displayed.
 * NOTE: The params must be static, since they will decide what other hooks are going to be triggered.
 * For example, it's not possible to change from `debounceReady` method to `delayReady`. That will cause react bad stuff
 * @param props
 * @returns signal telling whether the view is ready
 */
export function useViewIsReady(props: ViewIsReadyProps): ReadonlySignal<boolean> {
  // Debounce
  if (props.debounceReady !== undefined) {
    const isDebounced = useSignal(false)
    const isReady = useComputed(() => {
      if (!isDebounced.value) {
        return false
      }
      if (props.isReady !== undefined) {
        return !!props.isReady.value
      }
      if (props.isNotReady !== undefined) {
        return !props.isNotReady.value
      }

      return true
    })

    useEffect(() => {
      const debounceTimeout = setTimeout(
        () => (isDebounced.value = true),
        props.debounceReady === true ? 500 : props.debounceReady,
      )

      return () => clearTimeout(debounceTimeout)
    }, [])

    return isReady
  }

  // Simple animation
  if (props.delayReady) {
    const isDelayed = useSignal(false)
    const isReady = useComputed(() => {
      if (!isDelayed.value) {
        return false
      }
      if (props.isReady !== undefined) {
        return !!props.isReady.value
      }
      if (props.isNotReady !== undefined) {
        return !props.isNotReady.value
      }

      return true
    })

    useEffect(() => {
      isDelayed.value = true
    }, [])

    return isReady
  }

  // From parent
  if (props.isReady !== undefined) {
    return useComputed(() => !!props.isReady?.value)
  }

  // From parent, opposite value
  if (props.isNotReady !== undefined) {
    return useComputed(() => !props.isNotReady?.value)
  }

  // Not configured, always ready
  return useComputed(() => true)
}
