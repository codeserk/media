import { useComputed } from '@preact-signals/safe-react'
import { FC, PropsWithChildren, ReactNode } from 'react'
import { View as RNView, ViewStyle } from 'react-native'

import { useThemeStoreContext } from '../modules/theme/theme.store'
import { ThemeViewStyle, ViewIsReadyProps } from '../modules/view/view.types'
import { useViewIsReady } from '../modules/view/view-is-ready.hook'

export interface ViewProps extends ThemeViewStyle, ViewIsReadyProps, PropsWithChildren {
  readonly style?: ViewStyle
  readonly skeleton?: ReactNode
}

export const View: FC<ViewProps> = (props) => {
  const isReady = useViewIsReady(props)
  const { compileViewStyle } = useThemeStoreContext()

  const style = useComputed(() => ({ ...compileViewStyle(props), ...props.style }))

  if (!isReady.value) {
    return props.skeleton ?? null
  }

  return <RNView style={style.value}>{props.children}</RNView>
}
