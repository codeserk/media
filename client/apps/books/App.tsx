import { MediaStoreContext, useMediaStore } from 'core/modules/media/media.store'
import { MusicStoreContext } from 'core/modules/music/music.store'
import { Compose } from 'core/utils/compose'
import { StatusBar } from 'expo-status-bar'
import { HStack } from 'view/components/HStack'
import { Text } from 'view/components/Text'
import { View } from 'view/components/View'
import { VStack } from 'view/components/VStack'
import { ThemeStoreContext, useThemeStore } from 'view/modules/theme/theme.store'

export default function App() {
  const theme = useThemeStore({
    theme: {
      color: { red: '#994123' },
      space: { 2: 24 },
    },
  })
  const media = useMediaStore()

  return (
    <Compose
      components={[
        [ThemeStoreContext.Provider, { value: theme }],
        [MediaStoreContext.Provider, { value: media }],
        [MusicStoreContext.Provider, { value: media }],
      ]}>
      <View flex={1} alignItems="center" justifyContent="center">
        <Text>Open up App.tsx to start working on your app!</Text>
        <Text debounceReady={500} color="green" skeleton={<Text text="loading..." />}>
          OK
        </Text>
        <StatusBar style="auto" />

        <VStack space={2}>
          <Text>1</Text>
          <Text>2</Text>
          <Text>3</Text>
        </VStack>

        <HStack space={1}>
          <Text>1</Text>
          <Text>2</Text>
          <Text>3</Text>
        </HStack>
      </View>
    </Compose>
  )
}
