import { useComputed, useSignal } from '@preact-signals/safe-react'
import { MediaStoreContext, useMediaStore } from 'core/modules/media/media.store'
import { MusicStoreContext, useMusicStore } from 'core/modules/music/music.store'
import { useJellyfinSourceStore } from 'core/modules/sources/jellyfin/jellyfin.store'
import { Compose } from 'core/utils/compose'
import { Audio } from 'expo-av'
import { useFonts } from 'expo-font'
import { StatusBar } from 'expo-status-bar'
import { useEffect } from 'react'
import { Image, Pressable, StyleSheet } from 'react-native'
import { FadeInUp } from 'react-native-reanimated'
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
  const music = useMusicStore()

  const [loaded, error] = useFonts({
    'Poppins-Bold': require('./assets/fonts/Poppins-Bold.ttf'),
  })

  const width = useSignal(150)
  const { servers, collections, items, getServers, getImageUrl, getAudioStream } =
    useJellyfinSourceStore()
  const test = useSignal(0)
  const testValue = useComputed(() => test.toString())

  const playingId = useSignal<string | undefined>('')
  const audioSrc = useSignal<string | undefined>(undefined)

  const play = async (id: string) => {
    await Audio.setAudioModeAsync({
      playsInSilentModeIOS: true,
      staysActiveInBackground: true,
    })

    playingId.value = id

    // const file = await getAudioStream(id)
    // if (!file) {
    //   return
    // }

    const playbackObject = new Audio.Sound()
    await playbackObject.loadAsync({
      uri: `http://192.168.2.250:8096/Audio/${id}/universal?UserId=c61421596e5c45adac6cdd366a7a9471&DeviceId=TW96aWxsYS81LjAgKE1hY2ludG9zaDsgSW50ZWwgTWFjIE9TIFggMTBfMTVfNykgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEzMC4wLjAuMCBTYWZhcmkvNTM3LjM2fDE3MzEyNTgxOTM0MDk1&MaxStreamingBitrate=140000000&Container=opus%2Cwebm%7Copus%2Cmp3%2Caac%2Cm4a%7Caac%2Cm4b%7Caac%2Cflac%2Cwebma%2Cwebm%7Cwebma%2Cwav%2Cogg&TranscodingContainer=ts&TranscodingProtocol=hls&AudioCodec=aac&api_key=05df751c7ed84d4081aa21812ee30843&PlaySessionId=1735235783954&StartTimeTicks=0&EnableRedirection=true&EnableRemoteMedia=false`,
    })
    // playbackObject
    await playbackObject.playAsync()
    // playbackObject.
    //
    // console.log('file is', file)
    // audioSrc.value = `http://192.168.2.250:8096/Audio/${id}/universal?UserId=c61421596e5c45adac6cdd366a7a9471&DeviceId=TW96aWxsYS81LjAgKE1hY2ludG9zaDsgSW50ZWwgTWFjIE9TIFggMTBfMTVfNykgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzEzMC4wLjAuMCBTYWZhcmkvNTM3LjM2fDE3MzEyNTgxOTM0MDk1&MaxStreamingBitrate=140000000&Container=opus%2Cwebm%7Copus%2Cmp3%2Caac%2Cm4a%7Caac%2Cm4b%7Caac%2Cflac%2Cwebma%2Cwebm%7Cwebma%2Cwav%2Cogg&TranscodingContainer=ts&TranscodingProtocol=hls&AudioCodec=aac&api_key=05df751c7ed84d4081aa21812ee30843&PlaySessionId=1735235783954&StartTimeTicks=0&EnableRedirection=true&EnableRemoteMedia=false`
  }

  useEffect(() => {
    console.log('items are', items.value)

    for (const item of items.value) {
      if (item.Type === 'MusicArtist') {
        console.log(item)
      }
    }
  }, [items.value])

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

        <View
          debounceReady={500}
          height={40}
          width={width}
          backgroundColor="red"
          animated={{ entering: FadeInUp }}
          filter="grayscale: 10">
          <Text>Hello?</Text>
          {console.log('render inside')}
        </View>
        {console.log('render')}

        <Pressable
          style={({ pressed }) => [
            { opacity: pressed ? 0.5 : 1.0, padding: 20, backgroundColor: 'blue' },
          ]}
          onPress={() => (width.value += 10)}>
          <Text>Click</Text>
        </Pressable>

        {collections.value.map((it) => (
          <View key={it.Id ?? '-'}>
            <Text>Name? {it.Name}</Text>
            <Image
              source={{ uri: getImageUrl(it) }}
              width={200}
              height={200}
              style={{ width: 200, height: 200 }}
            />
          </View>
        ))}

        {/* <ScrollView>
          {items.value.map((it) => (
            <View key={it.Id ?? '-'}>
              <Text>{it.Name}</Text>

              <TouchableOpacity onPress={() => play(it.Id)}>
                <Image
                  source={{ uri: getImageUrl(it) }}
                  width={200}
                  height={200}
                  style={{ width: 200, height: 200 }}
                />
              </TouchableOpacity>
            </View>
          ))}
        </ScrollView> */}
      </View>
    </Compose>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
})
