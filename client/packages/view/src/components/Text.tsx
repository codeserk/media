import { signalValue } from 'core/modules/signals/signals'
import { SignalOrValue } from 'core/modules/signals/signals.types'
import { FC, PropsWithChildren } from 'react'
import { Text as RNText } from 'react-native'

interface Props extends PropsWithChildren {
  readonly $text?: SignalOrValue<string>
}

export const Text: FC<Props> = ({ $text, children }) => {
  const text = signalValue($text)

  return <RNText>{children ?? text}</RNText>
}
