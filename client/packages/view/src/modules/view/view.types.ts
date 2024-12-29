import { ReadonlySignal } from '@preact-signals/safe-react'
import { FlexStyle, ShadowStyleIOS, ViewStyle } from 'react-native'

import { DefaultTheme } from '../../theme'
import { StyleDimensions, Theme, ThemeStyle } from '../theme/theme.types'

export const flexStyleDimensions = {
  // Flex
  aspectRatio: 'aspectRatio',
  // - Border
  borderBottomWidth: 'size',
  borderEndWidth: 'size',
  borderLeftWidth: 'size',
  borderRightWidth: 'size',
  borderStartWidth: 'size',
  borderTopWidth: 'size',
  borderWidth: 'size',

  // - Flex/gap
  rowGap: 'space',
  gap: 'space',
  columnGap: 'space',
  flexGrow: 'space',
  flexShrink: 'space',

  // - Position
  top: 'space',
  right: 'space',
  bottom: 'space',
  left: 'space',
  start: 'space',
  end: 'space',
  // -- Inset
  inset: 'space',
  insetBlock: 'space',
  insetBlockEnd: 'space',
  insetBlockStart: 'space',
  insetInline: 'space',
  insetInlineEnd: 'space',
  insetInlineStart: 'space',

  // - Margin
  margin: 'space',
  marginBottom: 'space',
  marginEnd: 'space',
  marginHorizontal: 'space',
  marginLeft: 'space',
  marginRight: 'space',
  marginStart: 'space',
  marginTop: 'space',
  marginVertical: 'space',
  // -- Block
  marginBlock: 'space',
  marginBlockEnd: 'space',
  marginBlockStart: 'space',
  marginInline: 'space',
  marginInlineEnd: 'space',
  marginInlineStart: 'space',

  // - Padding
  padding: 'space',
  paddingBottom: 'space',
  paddingEnd: 'space',
  paddingHorizontal: 'space',
  paddingLeft: 'space',
  paddingRight: 'space',
  paddingStart: 'space',
  paddingTop: 'space',
  paddingVertical: 'space',
  // -- Block
  paddingBlock: 'space',
  paddingBlockEnd: 'space',
  paddingBlockStart: 'space',
  paddingInline: 'space',
  paddingInlineEnd: 'space',
  paddingInlineStart: 'space',

  // - Size
  minHeight: 'size',
  height: 'size',
  maxHeight: 'size',
  minWidth: 'size',
  width: 'size',
  maxWidth: 'size',
} as const satisfies StyleDimensions<FlexStyle>

export const shadowStyleDimensions = {
  // Shadow
  shadowColor: 'color',
  shadowOffset: 'space',
  shadowRadius: 'borderRadius',
} as const satisfies StyleDimensions<ShadowStyleIOS>

export const viewStyleDimensions = {
  ...flexStyleDimensions,
  ...shadowStyleDimensions,

  // View
  backgroundColor: 'color',
  borderBlockColor: 'color',
  borderBlockEndColor: 'color',
  borderBlockStartColor: 'color',
  borderBottomColor: 'color',
  borderBottomEndRadius: 'borderRadius',
  borderBottomLeftRadius: 'borderRadius',
  borderBottomRightRadius: 'borderRadius',
  borderBottomStartRadius: 'borderRadius',
  borderColor: 'color',
  borderEndColor: 'color',
  borderEndEndRadius: 'borderRadius',
  borderEndStartRadius: 'borderRadius',
  borderLeftColor: 'color',
  borderRadius: 'borderRadius',
  borderRightColor: 'color',
  borderStartColor: 'color',
  borderStartEndRadius: 'borderRadius',
  borderStartStartRadius: 'borderRadius',
  borderTopColor: 'color',
  borderTopEndRadius: 'borderRadius',
  borderTopLeftRadius: 'borderRadius',
  borderTopRightRadius: 'borderRadius',
  borderTopStartRadius: 'borderRadius',
} as const satisfies StyleDimensions<ViewStyle>

export type ThemeViewStyle<TTheme extends Theme = DefaultTheme> = ThemeStyle<
  TTheme,
  ViewStyle,
  typeof viewStyleDimensions
>

export interface ViewIsReadyProps {
  /** Signal to indicate whether the view is ready. Useful when the readiness comes from the parent */
  readonly isReady?: ReadonlySignal<unknown>

  /** Signal to indicate whether the view is not ready. Opposite to ready prop, useful for placeholders */
  readonly isNotReady?: ReadonlySignal<boolean>

  /** Animates the view, ready after first first effect */
  readonly delayReady?: boolean

  /** Animates the view, ready after being debounced */
  readonly debounceReady?: true | number
}
