import { useComputed, useSignal } from '@preact-signals/safe-react'
import { useJellyfinSourceStore } from 'core/modules/sources/jellyfin/jellyfin.store'
import { Compose } from 'core/utils/compose'
import { Audio } from 'expo-av'
import { StatusBar } from 'expo-status-bar'
import { Image, Pressable, StyleSheet } from 'react-native'
import { Text } from 'view/components/Text'
import { View } from 'view/components/View'
import { ThemeStoreContext, useThemeStore } from 'view/modules/theme/theme.store'

export default function App() {
  const theme = useThemeStore({
    theme: {
      color: { red: '#994123' },
    },
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

  return (
    <Compose components={[[ThemeStoreContext.Provider, { value: theme }]]}>
      <View flex={1} alignItems="center" justifyContent="center">
        <Text>Open up App.tsx to start working on your app!</Text>
        <Text>OK</Text>
        <StatusBar style="auto" />

        <View
          debounceReady={500}
          height={30}
          width={width}
          backgroundColor="red"
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
