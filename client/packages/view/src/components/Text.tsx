import { useComputed } from '@preact-signals/safe-react'
import { signalValue } from 'core/modules/signals/signals'
import { SignalOrValue } from 'core/modules/signals/signals.types'
import { FC, PropsWithChildren, ReactNode } from 'react'
import { Text as RNText, TextStyle } from 'react-native'

import { ThemeTextStyle } from '../modules/text/text.types'
import { useThemeStoreContext } from '../modules/theme/theme.store'
import { ViewIsReadyProps } from '../modules/view/view.types'
import { useViewIsReady } from '../modules/view/view-is-ready.hook'

interface Props extends ThemeTextStyle, ViewIsReadyProps, PropsWithChildren {
  readonly text?: SignalOrValue<string>
  readonly style?: TextStyle
  readonly skeleton?: ReactNode
}

export const Text: FC<Props> = (props) => {
  const isReady = useViewIsReady(props)
  const { compileTextStyle } = useThemeStoreContext()

  const style = useComputed(() => ({ ...compileTextStyle(props), ...props.style }))

  const text = signalValue(props.text)

  if (!isReady.value) {
    return props.skeleton ?? null
  }

  return (
    <RNText data-component="Text" style={style.value}>
      {props.children ?? text}
    </RNText>
  )
}
