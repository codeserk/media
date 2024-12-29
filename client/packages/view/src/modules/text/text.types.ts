import { TextStyle as RNTextStyle } from 'react-native'

import { DefaultTheme } from '../../theme'
import { StyleDimensions, StyleDimensionsValue, Theme } from '../theme/theme.types'
import { viewStyleDimensions } from '../view/view.types'

export const textStyleDimensions = {
  // View
  ...viewStyleDimensions,

  // Text
  color: 'color',
  fontFamily: 'fontFamily',
  fontSize: 'fontSize',
  textDecorationColor: 'color',
  textShadowColor: 'color',
} as const satisfies StyleDimensions<RNTextStyle>

export type ThemeTextStyle<TTheme extends Theme = DefaultTheme> = {
  [K in keyof RNTextStyle]?: StyleDimensionsValue<
    TTheme,
    RNTextStyle,
    typeof textStyleDimensions,
    K
  >
}
