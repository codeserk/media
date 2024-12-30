import { useComputed } from '@preact-signals/safe-react'
import { SignalOrValue } from 'core/modules/signals/signals.types'
import { FC, PropsWithChildren, ReactNode } from 'react'
import { View as RNView, ViewStyle } from 'react-native'
import Animated, { EntryExitAnimationFunction } from 'react-native-reanimated'

import { useThemeStoreContext } from '../modules/theme/theme.store'
import { ThemeViewStyle, ViewIsReadyProps } from '../modules/view/view.types'
import { useViewIsReady } from '../modules/view/view-is-ready.hook'

export interface ViewProps extends ThemeViewStyle, ViewIsReadyProps, PropsWithChildren {
  readonly style?: ViewStyle
  readonly space?: SignalOrValue<number>
  readonly skeleton?: ReactNode
  readonly animated?: {
    readonly entering?: EntryExitAnimationFunction
    readonly exiting?: EntryExitAnimationFunction
  }
}

export const View: FC<ViewProps> = (props) => {
  const isReady = useViewIsReady(props)
  const { compileViewStyle } = useThemeStoreContext()

  const style = useComputed(() => ({
    ...compileViewStyle({ gap: props.space, ...props }),
    ...props.style,
  }))

  if (!isReady.value) {
    return props.skeleton ?? null
  }

  if (!props.animated) {
    return <RNView style={style.value}>{props.children}</RNView>
  }

  return (
    <Animated.View
      style={style.value}
      entering={props.animated.entering}
      exiting={props.animated.exiting}>
      {props.children}
    </Animated.View>
  )
}
