import { createContext, useContext } from 'react'
import { TextStyle, ViewStyle } from 'react-native'

import { compileThemeTextStyle } from '../text/text.service'
import { compileThemeViewStyle } from '../view/view.service'
import { StyleDimensions, Theme, ThemeStyle } from './theme.types'

interface Props<TTheme extends Theme> {
  readonly theme: TTheme
}

export function useThemeStore<TTheme extends Theme>(props: Props<TTheme>) {
  // State

  const compileViewStyle = (
    style: ThemeStyle<TTheme, ViewStyle, StyleDimensions<ViewStyle>>,
  ): ViewStyle => {
    return compileThemeViewStyle(props.theme, style)
  }

  const compileTextStyle = (
    style: ThemeStyle<TTheme, TextStyle, StyleDimensions<TextStyle>>,
  ): TextStyle => {
    return compileThemeTextStyle(props.theme, style)
  }

  return {
    theme: props.theme,

    compileViewStyle,
    compileTextStyle,
  }
}

export type ThemeStore = ReturnType<typeof useThemeStore>

export const ThemeStoreContext = createContext<ThemeStore>({
  theme: {},
  compileViewStyle: (style) => style as ViewStyle,
  compileTextStyle: (style) => style as TextStyle,
})

export function useThemeStoreContext() {
  return useContext(ThemeStoreContext)
}
